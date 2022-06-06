package internal

import (
	"bridge-relayer/binding"
	"bridge-relayer/internal/message"
	"bridge-relayer/services/event"
	"fmt"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
	"strings"
)

var MessageAll *message.Message

func init() {
	MessageAll = &message.Message{
		MessageId: map[[32]byte][]byte{},
	}
}

func BuildQuery(contract ethCommon.Address, sig event.Sig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethCommon.Address{contract},
		Topics: [][]ethCommon.Hash{
			{sig.GetTopic()},
		},
	}

	return query
}

func GetContractAbi(abiName string) (abi.ABI, error) {
	pathAbi := fmt.Sprintf(binding.GetCurrentAbsPathByCaller()+"/%s/%s.json", abiName, abiName)
	contractBytes, err := ioutil.ReadFile(pathAbi)
	if err != nil {
		return abi.ABI{}, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(contractBytes)))
	if err != nil {
		return abi.ABI{}, err
	}

	return contractAbi, err
}
