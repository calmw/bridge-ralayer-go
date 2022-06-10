package tests

import (
	"bridge-relayer/log"
	"bridge-relayer/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestSetVoteThreshold(t *testing.T) {
	fmt.Println(uint32(len(ReLayer)), 33)
	_, err := ManagerContract.Contract.AdminSetVoteThreshold(ManagerContract.TransactOpts, uint32(len(ReLayer)))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	fmt.Println("SetVoteThreshold success")
}

func TestGrantRole(t *testing.T) {
	ValidatorRoleBytes, err := hex.DecodeString(ValidatorRole)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	for _, r := range ReLayer {
		_, err = ManagerContract.Contract.GrantRole(ManagerContract.TransactOpts, utils.ByteSliceToByte32(ValidatorRoleBytes),
			common.HexToAddress(r["address"]),
		)
		if err != nil {
			log.Logger.Error(err.Error())
			return
		}
	}

	fmt.Println("GrantRole success")
}

func TestRevokeRole(t *testing.T) {
	ValidatorRoleBytes, err := hex.DecodeString(ValidatorRole)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	for _, r := range ReLayer {
		_, err = ManagerContract.Contract.RevokeRole(ManagerContract.TransactOpts, utils.ByteSliceToByte32(ValidatorRoleBytes),
			common.HexToAddress(r["address"]),
		)
		if err != nil {
			log.Logger.Error(err.Error())
			return
		}
	}

	fmt.Println("RevokeRole success")
}

func TestSetConfigResource(t *testing.T) {
	automaticCall := true
	ResourceIdBytes, err := hex.DecodeString(ResourceId)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	_, err = ManagerContract.Contract.AdminSetConfigResource(ManagerContract.TransactOpts, utils.ByteSliceToByte32(ResourceIdBytes),
		automaticCall,
	)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	fmt.Println("SetConfigResource success")
}
