package lib

import (
	"bridge-relayer/tron/webThree"
	"bridge-relayer/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

type BridgeContract struct {
	OwnerAddress    string
	ContractAddress string
}

func NewBridgeContract(contractAddress, ownerAddress string) *BridgeContract {
	if strings.HasPrefix(ownerAddress, "T") {
		address, _ := utils.Base58ToHex(&ownerAddress)
		ownerAddress = *address
	}
	if strings.HasPrefix(contractAddress, "T") {
		address, _ := utils.Base58ToHex(&contractAddress)
		contractAddress = *address
	}

	return &BridgeContract{
		OwnerAddress:    ownerAddress,
		ContractAddress: contractAddress,
	}
}

func (b *BridgeContract) AdminSetChainId(id int64) error {

	arguments := abi.Arguments{
		{
			Type: webThree.Uint256Ty,
		},
	}
	bytes, _ := arguments.Pack(
		big.NewInt(id),
	)
	parameter := hexutil.Encode(bytes)
	parameter = strings.Replace(parameter, "0x", "", 1)

	data := webThree.CallContractRequest{
		OwnerAddress:     b.OwnerAddress,
		ContractAddress:  b.ContractAddress,
		FunctionSelector: "adminSetChainId(uint32)",
		Parameter:        parameter,
		FeeLimit:         webThree.DefaultFeeLimit,
		Visible:          false,
	}
	deployBytes, _ := json.Marshal(data)
	res, err := utils.HttpPost(webThree.Web3Uri, map[string]string{}, deployBytes)
	if err != nil {
		return errors.New("server error")
	}
	if strings.Contains(string(res), "code") {
		var callContractResponseErr webThree.CallContractResponseErr
		_ = json.Unmarshal(res, &callContractResponseErr)
		return errors.New(callContractResponseErr.Result.Code)
	}
	return nil
}

func (b *BridgeContract) ChainId() (error, int64) {

	data := webThree.CallContractRequest{
		OwnerAddress:     b.OwnerAddress,
		ContractAddress:  b.ContractAddress,
		FunctionSelector: "ChainId()",
		FeeLimit:         webThree.DefaultFeeLimit,
		Visible:          false,
	}
	deployBytes, _ := json.Marshal(data)
	res, err := utils.HttpPost(webThree.Web3Uri, map[string]string{}, deployBytes)
	if err != nil {
		return errors.New("server error"), 0
	}
	if strings.Contains(string(res), "code") {
		var callContractResponseErr webThree.CallContractResponseErr
		_ = json.Unmarshal(res, &callContractResponseErr)
		return errors.New(callContractResponseErr.Result.Code), 0
	}
	var callContractResponse webThree.CallContractResponse
	_ = json.Unmarshal(res, &callContractResponse)
	fmt.Println(string(res), err, 9998)
	return nil, 0
}
