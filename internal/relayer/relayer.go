package relayer

import (
	"bridge-relayer/config"
	"bridge-relayer/keyStore"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"os"
)

type ReLayer struct {
	Address      ethCommon.Address
	PrivateKey   *ecdsa.PrivateKey
	LatestBlock  *big.Int
	TransactOpts *bind.TransactOpts
}

var ThisReLayer ReLayer

func InitReLayer() {
	auth, err := bind.NewKeyedTransactorWithChainID(ThisReLayer.PrivateKey, big.NewInt(config.EngineCfg.ChainId))
	if err != nil {
		log.Panicln(err)
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
	ThisReLayer.TransactOpts = &transactOpts
}

func (r *ReLayer) SetBlockStore(chainName string, blockNum *big.Int) error {
	ksFile := keyStore.GetCurrentAbsPathByCaller() + "/" + chainName + "-" + ThisReLayer.Address.String()
	ks, err := os.OpenFile(ksFile, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	_, err = ks.WriteString(blockNum.String())
	if err != nil {
		return err
	}

	return nil
}
