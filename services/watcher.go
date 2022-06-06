package services

import (
	"bridge-relayer/binding/bridge"
	"bridge-relayer/config"
	"bridge-relayer/internal"
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
	"sync"
	"time"
)

const VoteInactive = 0
const VoteActive = 1
const VotePassed = 2
const VoteCancelled = 3

var BlockRetryInterval = time.Second * 2

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
	ethCli, err := ethclient.Dial(config.ChainCfg[id].Endpoint)
	if nil != err {
		log.Logger.Error(err.Error())
		return nil, err
	}

	bridgeContract, err := bridge.NewBridge(config.ChainCfg[id].Bridge, ethCli)
	if nil != err {
		log.Logger.Error(err.Error())
		return nil, err
	}

	return &Watcher{
		Id:             id,
		Cfg:            config.ChainCfg[id],
		EthCli:         ethCli,
		Log:            log15.Root().New("chain", config.ChainCfg[id].Name),
		BridgeContract: bridgeContract,
	}, nil
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
		watcher.Log.Info("Starting listener... ")
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
			w.Log.Error("GetBlockNum ", "err", err)
			continue
		}
		// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
		blockDelay := big.NewInt(config.DefaultBlockConfirmations)
		if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(blockDelay) == -1 {
			w.Log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
			continue
		}

		w.Log.Info("Polling Blocks...", "block", currentBlock.String(), "latest", latestBlock)

		// Parse out ConfirmedRequest events
		err, msg := w.getBridgeEventLogsFromBlock(event.ConfirmedRequestEvent.EventSignature, currentBlock)
		if err != nil {
			w.Log.Warn("get ConfirmedRequest logs", "err", err, "block", currentBlock.String())
		} else {
			internal.MessageAll.Save(msg.MessageId, msg.Data)
			err := w.Vote(msg, currentBlock)
			if err != nil {
				w.Log.Error("ConfirmedRequest Vote err", "err", err, "messageId", hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)))
			} else {
				log.Logger.Info("ConfirmedRequest Vote success")
			}
		}

		// Parse out CallRequest events
		err, msg = w.getBridgeEventLogsFromBlock(event.CallRequestEvent.EventSignature, currentBlock)
		if err != nil {
			w.Log.Warn("get CallRequest logs", "err", err, "block", currentBlock.String())
		} else {
			internal.MessageAll.Save(msg.MessageId, msg.Data)
			err := w.Vote(msg, currentBlock)
			if err != nil {
				w.Log.Error("CallRequest Vote err", "err", err, "messageId", "0x"+hex.EncodeToString(utils.Byte32ToByteSlice(msg.MessageId)))
			} else {
				log.Logger.Info("CallRequest Vote success")
			}
		}

		w.SetBlockStore(currentBlock)

		currentBlock.Add(currentBlock, big.NewInt(1))
	}

}

func (w *Watcher) SetBlockStore(blockNumber *big.Int) {
	err := relayer.ThisReLayer.SetBlockStore(w.Cfg.Name, blockNumber)
	if err != nil {
		w.Log.Error("SetBlockStore", "err", err, "block", blockNumber.String())
	}
}

func (w *Watcher) getBridgeEventLogsFromBlock(sig event.Sig, latestBlock *big.Int) (error, DataMsg) {

	query := internal.BuildQuery(w.Cfg.Bridge, sig, latestBlock, latestBlock)

	// querying for logs
	logs, err := w.EthCli.FilterLogs(context.Background(), query)
	if err != nil {
		return errors.New("unable to Filter Logs"), DataMsg{}
	}

	if len(logs) <= 0 {
		return errors.New("no event Logs"), DataMsg{}
	}

	contractAbi, err := internal.GetContractAbi("bridge")
	var eventData []interface{}
	if err != nil {
		log.Logger.Error(err.Error())
		return err, DataMsg{}
	}

	var msg []DataMsg

	// read through the log events and handle their deposit event if handler is recognized
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
	hasVote, err := w.HasVote(msg.MessageId)
	if err != nil {
		return err
	}

	voteStatus, err := w.VoteStatus(msg.MessageId)
	if err != nil {
		w.Log.Error("voteStatus err", "err", err)
		log.Logger.Error(err.Error())
		return err
	}
	w.Log.Info("voteStatus ", "voteStatus", voteStatus)
	if voteStatus < VotePassed {
		if hasVote {
			return errors.New("already voted")
		}

		engine, err := NewEngine()
		if err != nil {
			w.Log.Error("NewEngine err ", "err", err)
			return err
		}

		MessageId := utils.Byte32ToByteSlice(msg.MessageId)
		signature, err := utils.Sign(MessageId, relayer.ThisReLayer.PrivateKey)
		if err != nil {
			w.Log.Error("Sign ", "Sign err", err)
			return err
		}

		_, err = engine.Vote(relayer.ThisReLayer.TransactOpts, msg.ResourceID, msg.MessageId, msg.SourceChainId, msg.SourceNonce, msg.TargetChainId,
			msg.Target, msg.DataHash, signature)
		if err != nil {
			w.Log.Error("Vote ", "Vote err", err)
			return err
		} else {
			w.Log.Info("Vote success", "messageId", hex.EncodeToString(MessageId))
			//w.WatchVote(currentBlock)

		}
	} else if voteStatus == VotePassed {
		log.Logger.Error("The vote already passed")
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
