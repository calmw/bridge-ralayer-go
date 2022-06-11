package tests

import (
	"bridge-relayer/log"
	"bridge-relayer/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestSetSignatureThreshold(t *testing.T) {
	for _, b := range BridgeContracts {
		_, err := b.Contract.AdminSetSignatureThreshold(b.TransactOpts, uint32(len(ReLayer)))
		if err != nil {
			log.Logger.Error(err.Error())
			return
		}
	}

	fmt.Println("SetVoteThreshold success")
}

func TestGrantRoleBridge(t *testing.T) {
	ValidatorRoleBytes, err := hex.DecodeString(ValidatorRole)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	for _, b := range BridgeContracts {
		for _, r := range ReLayer {
			_, err = b.Contract.GrantRole(b.TransactOpts, utils.ByteSliceToByte32(ValidatorRoleBytes),
				common.HexToAddress(r["address"]),
			)
			if err != nil {
				log.Logger.Error(err.Error())
				return
			}
		}
	}

	fmt.Println("GrantRole success")
}

func TestRevokeRoleBridge(t *testing.T) {
	ValidatorRoleBytes, err := hex.DecodeString(ValidatorRole)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	for _, b := range BridgeContracts {
		for _, r := range ReLayer {
			_, err = b.Contract.RevokeRole(b.TransactOpts, utils.ByteSliceToByte32(ValidatorRoleBytes),
				common.HexToAddress(r["address"]),
			)
			if err != nil {
				log.Logger.Error(err.Error())
				return
			}
		}
	}

	fmt.Println("RevokeRole success")
}

func TestAdminSetConfigResource(t *testing.T) {
	ResourceIdBytes, err := hex.DecodeString(ResourceId)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	for _, b := range BridgeContracts {
		_, err = b.Contract.AdminSetConfigResource(
			b.TransactOpts,
			utils.ByteSliceToByte32(ResourceIdBytes),
			big.NewInt(0),
			[]common.Address{})
		if err != nil {
			log.Logger.Error(err.Error())
			return
		}
	}

	fmt.Println("SetConfigResource success")
}

func TestCallRemote(t *testing.T) {
	ResourceIdBytes, err := hex.DecodeString(ResourceId)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	data, _ := hex.DecodeString("3565375400000000000000000000000000000000000000000000000000000000")

	rand.Seed(time.Now().UnixNano())
	bridgeContractsIndex := rand.Intn(len(BridgeContracts))
	var targetBridgeContractsIndex int
	if bridgeContractsIndex-1 < 0 {
		targetBridgeContractsIndex = bridgeContractsIndex + 1
	} else {
		targetBridgeContractsIndex = bridgeContractsIndex - 1
	}
	var target common.Address
	if BridgeContracts[targetBridgeContractsIndex].Id == 2 {
		target = common.HexToAddress("0x0080b5ec2aaa531236df68bc91ff0d487a24eaeb")
	} else {
		target = common.HexToAddress("0xca98cd464a9cc69409de5bfec896a45985f670a3")
	}
	fmt.Println(bridgeContractsIndex, targetBridgeContractsIndex, BridgeContracts[targetBridgeContractsIndex].Id, target.String())
	tx, err := BridgeContracts[bridgeContractsIndex].Contract.CallRemote(
		BridgeContracts[bridgeContractsIndex].TransactOpts,
		utils.ByteSliceToByte32(ResourceIdBytes),
		uint32(BridgeContracts[targetBridgeContractsIndex].Id),
		target,
		data,
	)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	fmt.Println("CallRemote success", tx.ChainId())
}
