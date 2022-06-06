package services

import (
	"bridge-relayer/binding/manager"
	"bridge-relayer/config"
	"bridge-relayer/internal"
	"bridge-relayer/internal/relayer"
	"bridge-relayer/log"
	"bridge-relayer/services/event"
	"bridge-relayer/utils"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type Engine struct {
	Cfg             config.Engine
	EthCli          *ethclient.Client
	Log             log15.Logger
	ManagerContract *manager.Manager // instance of bound bridge contract
}

type VoteRecord struct {
	ResourceID    [32]byte
	VoteStatus    uint8
	StartBlock    *big.Int
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	DataHash      [32]byte
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

func (e *Engine) GetVoteRecords(messageId [32]byte) (VoteRecord, error) {
	return e.ManagerContract.VoteRecords(nil, messageId)
}

func (e *Engine) GetBlockNum() (*big.Int, error) {
	header, err := e.EthCli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

func (e *Engine) GetSignatureCollectedEvent() {
	var currentBlock = e.Cfg.StartBlock
	for true {
		time.Sleep(BlockRetryInterval)
		latestBlock, err := e.GetBlockNum()
		if err != nil {
			e.Log.Error("GetBlockNum ", "err", err)
			continue
		}
		e.Log.Debug("GetSignatureCollectedEvent", "block", currentBlock, "latestBlock", latestBlock)

		// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
		blockDelay := big.NewInt(config.DefaultBlockConfirmations)
		if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(blockDelay) == -1 {
			e.Log.Debug("engine not ready, will retry", "target", currentBlock, "latest", latestBlock)
			continue
		}

		err, voteMsg := e.GetSignatureCollectedEventLogsFromBlock(e.Cfg.ManagerAddress, event.SignatureCollectedEvent.EventSignature, currentBlock)
		if err != nil {
			e.Log.Error("GetSignatureCollectedEvent err", "err", err, "block", currentBlock)
		} else {
			e.Log.Info("GetSignatureCollectedEvent, success", "voteMsg", voteMsg)
			chainConfig, err := e.ManagerContract.GetChainConfig(&bind.CallOpts{
				From:    relayer.ThisReLayer.Address,
				Context: context.Background(),
			}, voteMsg.ResourceID)
			if err != nil {
				e.Log.Error("GetChainConfig err", "err", err)
				return
			}
			if chainConfig.RemoteCallType == 0 {
				e.Log.Error("remote callType", "remote callType", "Manual", "messageId",
					hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)))
				log.Logger.Sugar().Info("remote callType Manual", hex.EncodeToString(utils.Byte32ToByteSlice(voteMsg.MessageId)))
			} else {
				executed, err := e.IsExecuted(voteMsg)
				log.Logger.Sugar().Info("is execute", executed)
				if err != nil {
					log.Logger.Error(err.Error())
					e.Log.Error("IsExecuted err", "err", err)
				} else {
					if !executed {
						_, err := e.Execute(voteMsg)
						if err != nil {
							log.Logger.Error(err.Error())
							e.Log.Error("Execute err", "err", err)
						}
						log.Logger.Info("Execute")
					} else {
						_, err := e.Confirm(voteMsg)
						if err != nil {
							log.Logger.Error(err.Error())
							e.Log.Error("Confirm err", "err", err)
						}
						log.Logger.Info("Confirm")
					}
					e.Log.Error("Execute or Confirm success")
				}
			}
		}

		currentBlock.Add(currentBlock, big.NewInt(1))
	}
}

func (e *Engine) GetSignatureCollectedEventLogsFromBlock(address common.Address, sig event.Sig, latestBlock *big.Int) (error, VoteMsg) {
	query := internal.BuildQuery(address, sig, latestBlock, latestBlock)

	// querying for logs
	logs, err := e.EthCli.FilterLogs(context.Background(), query)
	if err != nil {
		return errors.New("unable to Filter Logs"), VoteMsg{}
	}

	if len(logs) <= 0 {
		return errors.New("no event Logs"), VoteMsg{}
	}

	contractAbi, err := internal.GetContractAbi("manager")
	var eventData []interface{}
	if err != nil {
		log.Logger.Error(err.Error())
		return err, VoteMsg{}
	}

	var msg []VoteMsg

	// read through the log events and handle their deposit event if handler is recognized
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
		ok, data := internal.MessageAll.Get(messageId)
		if !ok {
			return errors.New("data empty " + messageId.String()), VoteMsg{}
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
			data,
			signatures,
		})
	}
	return nil, msg[0]
}

func (e *Engine) VoteStatus(messageId [32]byte) (uint8, error) {
	voteRecord, err := e.ManagerContract.VoteRecords(nil, messageId)
	if err != nil {
		log.Logger.Error(err.Error())
		return 0, err
	}

	return voteRecord.VoteStatus, nil
}

func (e *Engine) IsExecuted(msg VoteMsg) (bool, error) {
	watcher, err := NewWatcher(int(msg.TargetChainId))
	if err != nil {
		fmt.Println(err)
	}
	return watcher.BridgeContract.ExecutionRecord(nil, msg.MessageId)
}

func (e *Engine) Execute(msg VoteMsg) (*types.Transaction, error) {
	watcher, err := NewWatcher(int(msg.TargetChainId))
	if err != nil {
		fmt.Println(err)
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

func (e *Engine) Confirm(msg VoteMsg) (*types.Transaction, error) {
	watcher, _ := NewWatcher(int(msg.SourceChainId))
	transaction, err := watcher.BridgeContract.Execute(
		relayer.ThisReLayer.TransactOpts, msg.ResourceID, msg.SourceChainId, msg.SourceNonce,
		msg.MessageId, msg.TargetChainId, msg.Target, msg.Data, msg.Signatures)

	return transaction, err
}

func (e *Engine) NewTransactOpts(chainId int) (*bind.TransactOpts, error) {

	auth, err := bind.NewKeyedTransactorWithChainID(relayer.ThisReLayer.PrivateKey, big.NewInt(config.Config.Chains[chainId].ChainId))
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
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false, // Do all transact steps but do not send the transaction
	}
	return &transactOpts, err
}