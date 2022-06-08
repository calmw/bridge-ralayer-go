package tests

import (
	"bridge-relayer/binding/bridge"
	"bridge-relayer/binding/manager"
	"bridge-relayer/config"
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"path/filepath"
)

type Manager struct {
	TransactOpts *bind.TransactOpts
	Contract     *manager.Manager
}
type Bridge struct {
	TransactOpts *bind.TransactOpts
	Contract     *bridge.Bridge
}

var ResourceId = "3565375400000000000000000000000000000000000000000000000000000000"
var DeployTransactOpts *bind.TransactOpts
var ValidatorRole = "a95257aebefccffaada4758f028bce81ea992693be70592f620c4c9a0d9e715a"
var ValidatorAddress = []string{"0x8f03D4Ce81C3c2dB006C1C725d2e70C3ecC69916", "0x3bd1a4c59b575eC77dDBd9c9c0a46633E5D5Bec7", "0x5384ee6148a201Dd6Ba962EA2d7673c493B01B5e"}
var ReLayer = []map[string]string{
	{
		"privateKey": "8699040b13da6c1994f97bef8d2fe458bf5c23e6ca5a97d45bd4663eaf90b856",
		"address":    "0x8f03D4Ce81C3c2dB006C1C725d2e70C3ecC69916",
	}, {
		"privateKey": "c813f8c65d2c26c019bdf65b64fd55128d27180d9f080f5d1d3e4729a1d4b5d3",
		"address":    "0x3bd1a4c59b575eC77dDBd9c9c0a46633E5D5Bec7",
	}}
var DeployPrivateKey = "7bdc80fb5cb54cb2ebcbbc697b2f58d93ab21c35570aa4f4614651a3781dfb37"
var ManagerContract Manager
var BridgeContracts []Bridge

func init() {

	currentAbPath := config.GetCurrentAbsPathByCaller()
	tomlFile, err := filepath.Abs(currentAbPath + "/config.toml")
	if err != nil {
		panic("read toml file err: " + err.Error())
		return
	}

	if _, err := toml.DecodeFile(tomlFile, &config.Config); err != nil {
		panic("read toml file err: " + err.Error())
		return
	}

	config.ChainCfg = map[int]config.Chain{}
	for i := 0; i < len(config.Config.Chains); i++ {
		Cli, err := ethclient.Dial(config.Config.Chains[i].Endpoint)
		if nil != err {
			fmt.Println("err: ", err)
			return
		}

		b := Bridge{}
		b.Contract, err = bridge.NewBridge(common.HexToAddress(config.Config.Chains[i].Bridge), Cli)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}

		deployPrivateKeyEcdsa, err := crypto.HexToECDSA(DeployPrivateKey)
		if err != nil {
			fmt.Println("deploy privateKey err ", err)
			return
		}
		deployAuth, err := bind.NewKeyedTransactorWithChainID(deployPrivateKeyEcdsa, big.NewInt(config.Config.Chains[i].ChainId))
		if err != nil {
			log.Panicln(err)
		}
		transactOpts := bind.TransactOpts{
			From:      deployAuth.From,
			Nonce:     nil,
			Signer:    deployAuth.Signer, // Method to use for signing the transaction (mandatory)
			Value:     big.NewInt(0),
			GasPrice:  nil,
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   context.Background(),
			NoSend:    false, // Do all transact steps but do not send the transaction
		}
		b.TransactOpts = &transactOpts
		BridgeContracts = append(BridgeContracts, b)
	}

	Cli, err := ethclient.Dial(config.EngineCfg.Endpoint)
	if nil != err {
		fmt.Println("err: ", err)
		return
	}
	ManagerContract.Contract, err = manager.NewManager(common.HexToAddress(config.Config.Engine.ManagerAddress), Cli)
	if nil != err {
		fmt.Println("err: ", err)
		return
	}
}
