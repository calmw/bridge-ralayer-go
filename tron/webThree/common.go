package webThree

import "github.com/ethereum/go-ethereum/accounts/abi"

var Uint256Ty abi.Type
var Uint32Ty abi.Type
var Uint8Ty abi.Type
var Bytes32Ty abi.Type
var AddressTy abi.Type
var StringTy abi.Type

const DefaultFeeLimit = 100000

func init() {
	Uint256Ty, _ = abi.NewType("uint256", "", nil)
	Uint32Ty, _ = abi.NewType("uint32", "", nil)
	Uint8Ty, _ = abi.NewType("uint8", "", nil)
	Bytes32Ty, _ = abi.NewType("bytes32", "", nil)
	AddressTy, _ = abi.NewType("address", "", nil)
	StringTy, _ = abi.NewType("string", "", nil)
}
