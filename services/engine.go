package services

import (
	"bridge-relayer/binding/manager"
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
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"time"
)

type Engine struct {
	Cfg             config.Engine
	EthCli          *ethclient.Client
	Log             log15.Logger
	ManagerContract *manager.Manager // instance of bound manager contract
}

func NewEngine() (*Engine, error) {
	ethCli, err := ethclient.Dial(config.EngineCfg.Endpoint)
	if nil != err {
		log.Logger.Error(err.Error())
		return nil, err
	}

	managerContract, err := manager.NewManager(config.EngineCfg.ManagerAddress, ethCli)
	if nil != err {
		log.Logger.Error(err.Error())
		return nil, err
	}

	return &Engine{
		config.EngineCfg,
		ethCli,
		log15.Root().New("chain", config.EngineCfg.Name),
		managerContract,
	}, nil
}

func (e *Engine) Vote(transactOpts *bind.TransactOpts, _resourceID [32]byte, messageId [32]byte,
	sourceChainId uint32, sourceNonce *big.Int, targetChainId uint32, target common.Address,
	dataHash [32]byte, signature []byte) (*types.Transaction, error) {

	return e.ManagerContract.Vote(transactOpts, _resourceID, messageId, sourceChainId,
		sourceNonce, targetChainId, target, dataHash, signature)
}

func (e *Engine) GetBlockNum() (*big.Int, error) {
	header, err := e.EthCli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

func (e *Engine) IsAutoCallExecute(voteMsg VoteMsg) (bool, error) {
	chainConfig, err := e.ManagerContract.GetChainConfig(&bind.CallOpts{
		From:    relayer.ThisReLayer.Address,
		Context: context.Background(),
	}, voteMsg.ResourceID)
	if err != nil {
		e.Log.Error("GetChainConfig error", "error", err)
		log.Logger.Error(err.Error())
		return false, err
	}
	if chainConfig.RemoteCallType == 0 {
		e.Log.Trace("remote call", "remote call", "Manual", "messageId",
			hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)), "status", "end of process")
		log.Logger.Sugar().Info("remote call Manual, end of process", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)))
		return false, nil
	}
	return true, nil
}

func (e *Engine) PollBlocks() {
	var currentBlock = e.Cfg.StartBlock
	for true {
		time.Sleep(BlockRetryInterval)

		// Get the latest block
		latestBlock, err := e.GetBlockNum()
		if err != nil {
			e.Log.Error("GetBlockNum ", "err", err)
			continue
		}

		// BlockDelay; (latest - current) < BlockDelay
		blockDelay := big.NewInt(config.DefaultBlockConfirmations)
		if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(blockDelay) == -1 {
			e.Log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
			continue
		}

		// Next block
		currentBlock.Add(currentBlock, big.NewInt(1))

		// Get event data
		err, voteMsg := e.GetSignatureCollectedEventLogsFromBlock(e.Cfg.ManagerAddress, event.SignatureCollectedEvent.EventSignature, currentBlock)
		if err != nil {
			if errors.Is(err, NoEventErr) {
				e.Log.Crit("No SignatureCollected Event", "block", currentBlock)
			} else {
				e.Log.Error("Get SignatureCollected Event error", "error", err)
				e.Log.Error(err.Error())
			}
			continue
		}

		// Check whether the resource ID is called automatically
		isAutoCallExecute, err := e.IsAutoCallExecute(voteMsg)
		if err != nil || !isAutoCallExecute {
			continue
		}

		// Call Bridge
		go e.CallBridge(voteMsg)
	}
}

func (e *Engine) GetSignatureCollectedEventLogsFromBlock(address common.Address, sig event.Sig, block *big.Int) (error, VoteMsg) {

	query := eventInternal.BuildQuery(address, sig, block, block)

	// querying for logs
	logs, err := e.EthCli.FilterLogs(context.Background(), query)
	if err != nil {
		return errors.New("unable to Filter Logs"), VoteMsg{}
	}

	if len(logs) <= 0 {
		return NoEventErr, VoteMsg{}
	}

	contractAbi, err := manager.ManagerMetaData.GetAbi()
	var eventData []interface{}
	if err != nil {
		log.Logger.Error(err.Error())
		return err, VoteMsg{}
	}

	var msg []VoteMsg

	// Read through the log events and handle their SignatureCollected event if handler is recognized
	for _, logE := range logs {
		eventData, _ = contractAbi.Unpack(event.SignatureCollectedEvent.EventName, logE.Data)

		messageId := logE.Topics[1]
		resourceId := eventData[0]
		voteStatus := eventData[1].(uint8)
		//caller := eventData[1]
		sourceChainId := eventData[2].(uint32)
		sourceNonce := eventData[3].(*big.Int)
		targetChainId := eventData[4].(uint32)
		target := eventData[5].(common.Address)
		signatures := eventData[6].([][]byte)
		dataHash := logE.Topics[2]

		ok, _ := CheckChainId(int(targetChainId))
		if !ok {
			return errors.New("event error: target chainId id error"), VoteMsg{}
		}
		ok, _ = CheckChainId(int(sourceChainId))
		if !ok {
			return errors.New("event error: source chainId id error"), VoteMsg{}
		}

		ok, MsgInfo := message.AllMessage.Get(messageId)
		if !ok {
			return errors.New("log data empty " + messageId.String()), VoteMsg{}
		}
		msg = append(msg, VoteMsg{
			resourceId.([32]byte),
			voteStatus,
			sourceChainId,
			sourceNonce,
			targetChainId,
			target,
			messageId,
			dataHash,
			MsgInfo.Data,
			signatures,
		})
	}

	e.Log.Info("Get SignatureCollected Event Success", "messageId",
		hex.EncodeToString(utils.Byte32ToByteSlice(msg[0].MessageId)), "block", block)
	log.Logger.Sugar().Info("GetSignatureCollectedEvent, success ", "messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(msg[0].MessageId)))

	return nil, msg[0]
}

func (e *Engine) VoteStatus(messageId [32]byte) (uint8, error) {
	voteRecord, err := e.ManagerContract.VoteRecords(nil, messageId)
	if err != nil {
		return 0, err
	}

	return voteRecord.VoteStatus, nil
}

func (e *Engine) IsExecuted(msg VoteMsg) (bool, error) {
	watcher, err := NewWatcher(int(msg.TargetChainId))
	if err != nil {
		return false, err
	}
	return watcher.BridgeContract.ExecutionRecord(nil, msg.MessageId)
}

func (e *Engine) Execute(msg VoteMsg) (*types.Transaction, error) {
	watcher, err := NewWatcher(int(msg.TargetChainId))
	if err != nil {
		return nil, err
	}

	opts, err := e.NewTransactOpts(int(msg.TargetChainId))
	if err != nil {
		return nil, err
	}
	transaction, err := watcher.BridgeContract.Execute(
		opts, msg.ResourceID, msg.SourceChainId, msg.SourceNonce,
		msg.MessageId, msg.TargetChainId, msg.Target, msg.Data, msg.Signatures)
	return transaction, err
}

func (e *Engine) CallBridge(voteMsg VoteMsg) {
	e.Log.Trace("remote call", "remote call", "auto", "messageId",
		hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)))

	rand.Seed(time.Now().UnixNano())
	time.Sleep(ExecuteRandTime)

	isExecuted, err := e.IsExecuted(voteMsg)
	log.Logger.Sugar().Info("is execute ", isExecuted)
	if err != nil {
		e.Log.Error("IsExecuted error", "error", err)
		log.Logger.Error(err.Error())
		return
	}

	if !isExecuted {
		tx, err := e.Execute(voteMsg)
		if err != nil {
			log.Logger.Error(err.Error())
			e.Log.Error("Execute err", "err", err)
		} else {
			e.Log.Info("Execute success ", "messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)))
			log.Logger.Sugar().Info("Execute success messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)),
				" reLayerAddress=", relayer.ThisReLayer.Address.String(),
				" txHash=", tx.Hash(),
			)
			message.AllMessage.Save(voteMsg.MessageId, voteMsg.Data, true)
		}
	} else {
		e.Log.Info("Already Executed,end process ", "messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)))
		log.Logger.Sugar().Info("Already Executed,end process messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)),
			" reLayerAddress=", relayer.ThisReLayer.Address.String(),
		)
		// 确认操作放至业务层，暂不处理
		//_, err := e.Confirm(voteMsg)
		//if err != nil {
		//	log.Logger.Error(err.Error())
		//	e.Log.Error("Confirm err", "err", err)
		//} else {
		//	log.Logger.Sugar().Info("Confirm success messageId=", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)), " reLayerAddress=", relayer.ThisReLayer.Address.String())
		//	message.AllMessage.Save(voteMsg.MessageId, voteMsg.Data, true)
		//}
	}
}

func (e *Engine) Confirm(msg VoteMsg) (*types.Transaction, error) {
	watcher, _ := NewWatcher(int(msg.SourceChainId))
	transaction, err := watcher.BridgeContract.Execute(
		relayer.ThisReLayer.TransactOpts, msg.ResourceID, msg.SourceChainId, msg.SourceNonce,
		msg.MessageId, msg.TargetChainId, msg.Target, msg.Data, msg.Signatures)

	return transaction, err
}

func (e *Engine) NewTransactOpts(Id int) (*bind.TransactOpts, error) {

	auth, err := bind.NewKeyedTransactorWithChainID(relayer.ThisReLayer.PrivateKey, big.NewInt(int64(config.ChainCfg[Id].ChainId)))
	if err != nil {
		return nil, err
	}

	transactOpts := bind.TransactOpts{
		From:      auth.From,
		Nonce:     nil,
		Signer:    auth.Signer, // Method to use for signing the transaction (mandatory)
		Value:     big.NewInt(0),
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  config.DefaultGasLimit,
		Context:   context.Background(),
		NoSend:    false, // Do all transact steps but do not send the transaction
	}
	return &transactOpts, err
}
