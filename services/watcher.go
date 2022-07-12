package services

import (
	"bridge-relayer/binding/bridge"
	"bridge-relayer/config"
	eventInternal "bridge-relayer/internal/event"
	"bridge-relayer/internal/message"
	"bridge-relayer/internal/relayer"
	"bridge-relayer/log"
	"bridge-relayer/services/event"
	"bridge-relayer/utils"
	"context"
	"encoding/hex"
	"errors"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

const VoteInactive = 0
const VoteActive = 1
const VotePassed = 2
const VoteCancelled = 3
const ChainNameLength = 10

var VoteRandTime = time.Duration((1 + rand.Intn(30)) * 1000000000)
var ExecuteRandTime = time.Duration((1 + rand.Intn(60)) * 1000000000)

var BlockRetryInterval = time.Second * 2

var NoEventErr = errors.New("no event logs")

type DataMsg struct {
	ResourceID    [32]byte
	MessageId     [32]byte
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	Target        common.Address
	Data          []byte
	DataHash      [32]byte
}

type VoteMsg struct {
	ResourceID    [32]byte
	VoteStatus    uint8
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	Target        common.Address
	MessageId     [32]byte
	DataHash      [32]byte
	Data          []byte
	Signatures    [][]byte
}

type Watcher struct {
	Id             int
	Cfg            config.Chain
	Log            log15.Logger
	EthCli         *ethclient.Client
	BridgeContract *bridge.Bridge // instance of bound bridge contract
}

func NewWatcher(id int) (*Watcher, error) {
	ok, chain := CheckChainId(id)
	if !ok {
		return nil, errors.New("chain id error")
	}

	ethCli, err := ethclient.Dial(chain.Endpoint)
	if nil != err {
		log.Logger.Sugar().Error(err.Error(), id)
		return nil, err
	}

	bridgeContract, err := bridge.NewBridge(chain.Bridge, ethCli)
	if nil != err {
		log.Logger.Error(err.Error())
		return nil, err
	}

	return &Watcher{
		Id:             id,
		Cfg:            config.ChainCfg[id],
		EthCli:         ethCli,
		Log:            log15.Root().New("chain", chain.Name),
		BridgeContract: bridgeContract,
	}, nil
}

func CheckChainId(id int) (bool, config.Chain) {
	var chain config.Chain
	var ok bool
	chain, ok = config.ChainCfg[id]
	return ok, chain
}

func StartWatcher() {
	wg := sync.WaitGroup{}
	wg.Add(config.ReLayerNum)
	for i := 0; i < config.ReLayerNum; i++ {
		watcher, err := NewWatcher(config.Config.Chains[i].Id)
		if err != nil {
			log.Logger.Sugar().Error("NewWatcher failed ", err)
			panic("NewWatcher failed")
		}
		watcher.Log.Info("Starting watcher ... ")
		go func() {
			watcher.PollBlocks()
			wg.Done()
		}()
	}
	wg.Wait()
}

func (w *Watcher) PollBlocks() {
	var currentBlock = w.Cfg.StartBlock
	for true {
		time.Sleep(BlockRetryInterval)

		latestBlock, err := w.GetBlockNum()
		if err != nil {
			w.Log.Error("GetBlockNum error", "error", err)
			continue
		}
		// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
		blockDelay := big.NewInt(config.DefaultBlockConfirmations)
		if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(blockDelay) == -1 {
			w.Log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
			continue
		}

		// Parse out ConfirmedRequest events
		err, msg := w.getBridgeEventLogsFromBlock(event.ConfirmedRequestEvent.EventSignature, currentBlock)
		if err != nil {
			if errors.Is(err, NoEventErr) {
				w.Log.Crit("No ConfirmedRequest Event", "block", currentBlock, "latestBlock", latestBlock)
			} else {
				w.Log.Error("Get ConfirmedRequest Event error", "error", err)
			}
		} else {
			w.Log.Info("get ConfirmedRequest success ", "messageId", hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)), "block", currentBlock.String())
			log.Logger.Sugar().Info("get ConfirmedRequest success ", "messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)), "block", currentBlock.String())

			// 确认操作放至业务层，暂不处理
			//message.AllMessage.Save(msg.MessageId, msg.Data, false)
			//err := w.Vote(msg, currentBlock)
			//if err != nil {
			//	w.Log.Error("ConfirmedRequest Vote err", "err", err, "messageId", hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)))
			//	log.Logger.Sugar().Error("ConfirmedRequest Vote err ", err, " messageId=", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)),
			//		" reLayerAddress=", relayer.ThisReLayer.Address.String())
			//} else {
			//	log.Logger.Sugar().Info("ConfirmedRequest Vote success ", "messageId=", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)),
			//		" reLayerAddress=", relayer.ThisReLayer.Address.String())
			//}
		}

		// Parse out CallRequest events
		err, msg = w.getBridgeEventLogsFromBlock(event.CallRequestEvent.EventSignature, currentBlock)
		if err != nil {
			if errors.Is(err, NoEventErr) {
				w.Log.Crit("No CallRequest Event", "block", currentBlock)
			} else {
				w.Log.Error("Get CallRequest Event error", "error", err)
			}
		} else {
			w.Log.Info("get CallRequest logs success", "block", currentBlock.String(), "messageId", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)))
			message.AllMessage.Save(msg.MessageId, msg.Data, false)
			go func() {
				err := w.Vote(msg, currentBlock)
				if err != nil {
					w.Log.Error("CallRequest Vote error", "error", err, "messageId", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)))
					log.Logger.Sugar().Error("CallRequest Vote error ", err, " messageId=", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)),
						" reLayerAddress=", relayer.ThisReLayer.Address.String())
				} else {
					log.Logger.Sugar().Info("CallRequest Vote success ", "messageId=", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)),
						" reLayerAddress=", relayer.ThisReLayer.Address.String())
				}
			}()
		}

		w.SetBlockStore(currentBlock)

		currentBlock.Add(currentBlock, big.NewInt(1))
	}

}

func (w *Watcher) SetBlockStore(blockNumber *big.Int) {
	err := relayer.ThisReLayer.SetBlockStore(w.Cfg.Name, blockNumber)
	if err != nil {
		w.Log.Error("SetBlockStore", "error", err, "block", blockNumber.String())
	}
}

func (w *Watcher) getBridgeEventLogsFromBlock(sig event.Sig, latestBlock *big.Int) (error, DataMsg) {

	query := eventInternal.BuildQuery(w.Cfg.Bridge, sig, latestBlock, latestBlock)

	// querying for logs
	logs, err := w.EthCli.FilterLogs(context.Background(), query)
	if err != nil {
		return NoEventErr, DataMsg{}
	}

	if len(logs) <= 0 {
		return NoEventErr, DataMsg{}
	}

	contractAbi, err := bridge.BridgeMetaData.GetAbi()
	var eventData []interface{}
	if err != nil {
		log.Logger.Error(err.Error())
		return err, DataMsg{}
	}

	var msg []DataMsg

	// read through the log events and handle their callRequest or confirmedRequest event if handler is recognized
	for _, logE := range logs {
		if sig == event.CallRequestEvent.EventSignature {
			eventData, err = contractAbi.Unpack(event.CallRequestEvent.EventName, logE.Data)
		} else if sig == event.ConfirmedRequestEvent.EventSignature {
			eventData, err = contractAbi.Unpack(event.ConfirmedRequestEvent.EventName, logE.Data)
		}
		if err != nil {
			return err, DataMsg{}
		}

		targetChainId := logE.Topics[2].Big().Int64()
		target := common.HexToAddress(logE.Topics[3].String())
		resourceId := eventData[0]
		//caller := eventData[1]
		sourceChainId := eventData[2]
		sourceNonce := eventData[3]
		data := eventData[4]
		dataHash := crypto.Keccak256Hash(data.([]byte))
		ok, _ := CheckChainId(int(targetChainId))
		if !ok {
			return errors.New("event error: target chainId id error"), DataMsg{}
		}
		ok, _ = CheckChainId(int(sourceChainId.(uint32)))
		if !ok {
			return errors.New("event error: source chainId id error"), DataMsg{}
		}
		msg = append(msg, DataMsg{
			resourceId.([32]byte),
			logE.Topics[1],
			sourceChainId.(uint32),
			sourceNonce.(*big.Int),
			uint32(targetChainId),
			target,
			eventData[4].([]byte),
			dataHash,
		})
	}

	return nil, msg[0]
}

func (w *Watcher) HasVote(messageId [32]byte) (bool, error) {

	engine, err := NewEngine()
	if err != nil {
		log.Logger.Error(err.Error())
		return false, err
	}
	has, err := engine.ManagerContract.HasVotedOnMessage(&bind.CallOpts{
		From:    relayer.ThisReLayer.Address,
		Context: context.Background(),
	}, messageId)
	if err != nil {
		return false, err
	}

	return has, nil
}

func (w *Watcher) VoteStatus(messageId [32]byte) (uint8, error) {

	engine, err := NewEngine()
	if err != nil {
		log.Logger.Error(err.Error())
		return 0, err
	}

	voteRecord, err := engine.ManagerContract.VoteRecords(nil, messageId)
	if err != nil {
		log.Logger.Error(err.Error())
		return 0, err
	}

	return voteRecord.VoteStatus, nil
}

func (w *Watcher) Vote(msg DataMsg, currentBlock *big.Int) error {

	rand.Seed(time.Now().UnixNano())
	time.Sleep(VoteRandTime)

	hasVote, err := w.HasVote(msg.MessageId)
	if err != nil {
		return err
	}

	voteStatus, err := w.VoteStatus(msg.MessageId)
	if err != nil {
		w.Log.Error("Get voteStatus error", "error", err)
		log.Logger.Error(err.Error())
		return err
	}
	w.Log.Info("voteStatus", "voteStatus", voteStatus)
	if voteStatus < VotePassed {
		if hasVote {
			return errors.New("already voted")
		}

		engine, err := NewEngine()
		if err != nil {
			w.Log.Error("NewEngine error", "error", err)
			return err
		}

		MessageId := utils.Byte32ToByteSlice(msg.MessageId)
		signature, err := utils.Sign(MessageId, relayer.ThisReLayer.PrivateKey)
		if err != nil {
			w.Log.Error("Sign ", "Sign error", err)
			return err
		}

		_, err = engine.Vote(relayer.ThisReLayer.TransactOpts, msg.ResourceID, msg.MessageId, msg.SourceChainId, msg.SourceNonce, msg.TargetChainId,
			msg.Target, msg.DataHash, signature)
		if err != nil {
			w.Log.Error("Vote ", "Vote error", err)
			return err
		} else {
			w.Log.Info("Vote success", "messageId", hex.EncodeToString(MessageId))
			//w.WatchVote(currentBlock)

		}
	} else if voteStatus == VotePassed {
		return errors.New("The vote already passed. ")
	} else {
		log.Logger.Error("The vote already cancelled")
		return errors.New("The vote already cancelled. ")
	}
	return nil
}

func (w *Watcher) Execute(msg VoteMsg) (*types.Transaction, error) {
	transaction, err := w.BridgeContract.Execute(
		relayer.ThisReLayer.TransactOpts, msg.ResourceID, msg.SourceChainId, msg.SourceNonce,
		msg.MessageId, msg.TargetChainId, msg.Target, msg.Data, msg.Signatures)

	return transaction, err
}

func (w *Watcher) GetBlockNum() (*big.Int, error) {
	header, err := w.EthCli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}
