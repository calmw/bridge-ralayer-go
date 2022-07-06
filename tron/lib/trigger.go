package lib

import (
	"bridge-relayer/tron/trigger"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

const DefaultFeeLimit = 100000000

func CallRemote(resourceId, targetChainId, target, data string) error {
	var parameter []map[string]string
	parameter = append(parameter, map[string]string{
		"bytes32": resourceId,
	})
	parameter = append(parameter, map[string]string{
		"uint32": targetChainId,
	})
	parameter = append(parameter, map[string]string{
		"address": target,
	})
	parameter = append(parameter, map[string]string{
		"bytes": data,
	})
	parameterBytes, err := json.Marshal(parameter)
	err = trigger.TriggerContract("TSUbUczwFjPt6vweMT4rrJ7yL1mzjaaEqk", "callRemote(bytes32,uint32,address,bytes)",
		string(parameterBytes), 100000000, 0,
		"", 0)
	return err
}
func Pack(methodName string, params ...interface{}) ([]byte, error) {
	myAbi := &abi.ABI{}
	if methodName == "" {
		// constructor
		arguments, err := myAbi.Constructor.Inputs.Pack(params...)
		if err != nil {
			return nil, err
		}
		return arguments, nil
	}
	method, exist := myAbi.Methods[methodName]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", methodName)
	}
	arguments, err := method.Inputs.Pack(params...)
	if err != nil {
		return nil, err
	}
	// Pack up the method ID too if not a constructor and return
	return append(method.ID, arguments...), nil
}

// signature [][]byte()
//func Execute(executRequest trigger.ExecuteRequest) error {
//	parameterBytes, err := json.Marshal(executRequest)
//	if err != nil {
//		return err
//	}
//	fmt.Println(string(parameterBytes), 9980)
//
//	err = trigger.TriggerContract("TSUbUczwFjPt6vweMT4rrJ7yL1mzjaaEqk", "execute(bytes32,uint32,uint256,bytes32,uint32,address,bytes,bytes[])",
//		string(parameterBytes), DefaultFeeLimit, 0,
//		"", 0)
//	return err
//}

//func Execute(jsonData string) error {
//
//	err := trigger.TriggerContract("TSUbUczwFjPt6vweMT4rrJ7yL1mzjaaEqk", "execute(bytes32,uint32,uint256,bytes32,uint32,address,bytes,bytes[])",
//		jsonData, DefaultFeeLimit, 0,
//		"", 0)
//	return err
//}

func Execute2(packed []byte) error {

	err := trigger.TriggerContract2("TSUbUczwFjPt6vweMT4rrJ7yL1mzjaaEqk", "execute(bytes32,uint32,uint256,bytes32,uint32,address,bytes,bytes[])",
		packed, DefaultFeeLimit, 0,
		"", 0)
	return err
}
