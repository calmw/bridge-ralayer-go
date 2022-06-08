package tests

import (
	"bridge-relayer/internal/relayer"
	"bridge-relayer/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestSetSignatureThreshold(t *testing.T) {
	for _, b := range BridgeContracts {
		_, err := b.Contract.AdminSetSignatureThreshold(relayer.ThisReLayer.TransactOpts, uint32(len(ReLayer)))
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
	}

	fmt.Println("SetVoteThreshold success")
}

func TestGrantRoleBridge(t *testing.T) {
	ValidatorRoleBytes, err := hex.DecodeString(ValidatorRole)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	for _, b := range BridgeContracts {
		for _, address := range ValidatorAddress {
			_, err = b.Contract.GrantRole(b.TransactOpts, utils.ByteSliceToByte32(ValidatorRoleBytes),
				common.HexToAddress(address),
			)
			if err != nil {
				fmt.Println("err: ", err)
				return
			}
		}
	}

	fmt.Println("GrantRole success")
}

func TestAdminSetConfigResource(t *testing.T) {
	ResourceIdBytes, err := hex.DecodeString(ResourceId)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	for _, b := range BridgeContracts {
		_, err = b.Contract.AdminSetConfigResource(b.TransactOpts, utils.ByteSliceToByte32(ResourceIdBytes), big.NewInt(0), []common.Address{})
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
	}

	fmt.Println("SetConfigResource success")
}
