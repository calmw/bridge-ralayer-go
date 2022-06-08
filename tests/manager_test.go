package tests

import (
	"bridge-relayer/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestSetVoteThreshold(t *testing.T) {
	_, err := ManagerContract.Contract.AdminSetVoteThreshold(DeployTransactOpts, uint32(len(ReLayer)))
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("SetVoteThreshold success")
}

func TestGrantRole(t *testing.T) {
	ValidatorRoleBytes, err := hex.DecodeString(ValidatorRole)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	for _, r := range ReLayer {
		fmt.Println(r["address"], 66)
		_, err = ManagerContract.Contract.GrantRole(DeployTransactOpts, utils.ByteSliceToByte32(ValidatorRoleBytes),
			common.HexToAddress(r["address"]),
		)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
	}

	fmt.Println("GrantRole success")
}

func TestSetConfigResource(t *testing.T) {
	ResourceIdBytes, err := hex.DecodeString(ResourceId)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	_, err = ManagerContract.Contract.AdminSetConfigResource(DeployTransactOpts, utils.ByteSliceToByte32(ResourceIdBytes),
		true,
	)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("SetConfigResource success")
}
