// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_chainId\",\"type\":\"uint32\"}],\"name\":\"adminSetChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blackList\",\"type\":\"address[]\"}],\"name\":\"adminSetConfigResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_signatureThreshold\",\"type\":\"uint32\"}],\"name\":\"adminSetSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sourceNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_chainId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_signatureThreshold\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ConfirmedRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_configResource\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"callRecords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumBridge.RecordState\",\"name\":\"recordState\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"confirmedRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executionRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) MANAGERROLE() ([32]byte, error) {
	return _Bridge.Contract.MANAGERROLE(&_Bridge.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) MANAGERROLE() ([32]byte, error) {
	return _Bridge.Contract.MANAGERROLE(&_Bridge.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) VALIDATORROLE() ([32]byte, error) {
	return _Bridge.Contract.VALIDATORROLE(&_Bridge.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _Bridge.Contract.VALIDATORROLE(&_Bridge.CallOpts)
}

// ConfigResource is a free data retrieval call binding the contract method 0x68d1b2a6.
//
// Solidity: function _configResource(bytes32 ) view returns(uint256 fee)
func (_Bridge *BridgeCaller) ConfigResource(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_configResource", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfigResource is a free data retrieval call binding the contract method 0x68d1b2a6.
//
// Solidity: function _configResource(bytes32 ) view returns(uint256 fee)
func (_Bridge *BridgeSession) ConfigResource(arg0 [32]byte) (*big.Int, error) {
	return _Bridge.Contract.ConfigResource(&_Bridge.CallOpts, arg0)
}

// ConfigResource is a free data retrieval call binding the contract method 0x68d1b2a6.
//
// Solidity: function _configResource(bytes32 ) view returns(uint256 fee)
func (_Bridge *BridgeCallerSession) ConfigResource(arg0 [32]byte) (*big.Int, error) {
	return _Bridge.Contract.ConfigResource(&_Bridge.CallOpts, arg0)
}

// CallRecords is a free data retrieval call binding the contract method 0x84b268d1.
//
// Solidity: function callRecords(bytes32 ) view returns(bytes32 resourceID, address caller, uint32 targetChainId, address target, bytes data, uint8 recordState)
func (_Bridge *BridgeCaller) CallRecords(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ResourceID    [32]byte
	Caller        common.Address
	TargetChainId uint32
	Target        common.Address
	Data          []byte
	RecordState   uint8
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "callRecords", arg0)

	outstruct := new(struct {
		ResourceID    [32]byte
		Caller        common.Address
		TargetChainId uint32
		Target        common.Address
		Data          []byte
		RecordState   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Caller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TargetChainId = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Target = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.RecordState = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// CallRecords is a free data retrieval call binding the contract method 0x84b268d1.
//
// Solidity: function callRecords(bytes32 ) view returns(bytes32 resourceID, address caller, uint32 targetChainId, address target, bytes data, uint8 recordState)
func (_Bridge *BridgeSession) CallRecords(arg0 [32]byte) (struct {
	ResourceID    [32]byte
	Caller        common.Address
	TargetChainId uint32
	Target        common.Address
	Data          []byte
	RecordState   uint8
}, error) {
	return _Bridge.Contract.CallRecords(&_Bridge.CallOpts, arg0)
}

// CallRecords is a free data retrieval call binding the contract method 0x84b268d1.
//
// Solidity: function callRecords(bytes32 ) view returns(bytes32 resourceID, address caller, uint32 targetChainId, address target, bytes data, uint8 recordState)
func (_Bridge *BridgeCallerSession) CallRecords(arg0 [32]byte) (struct {
	ResourceID    [32]byte
	Caller        common.Address
	TargetChainId uint32
	Target        common.Address
	Data          []byte
	RecordState   uint8
}, error) {
	return _Bridge.Contract.CallRecords(&_Bridge.CallOpts, arg0)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint32)
func (_Bridge *BridgeCaller) ChainId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint32)
func (_Bridge *BridgeSession) ChainId() (uint32, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint32)
func (_Bridge *BridgeCallerSession) ChainId() (uint32, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// ConfirmedRecord is a free data retrieval call binding the contract method 0x93e6ad97.
//
// Solidity: function confirmedRecord(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) ConfirmedRecord(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "confirmedRecord", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConfirmedRecord is a free data retrieval call binding the contract method 0x93e6ad97.
//
// Solidity: function confirmedRecord(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) ConfirmedRecord(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.ConfirmedRecord(&_Bridge.CallOpts, arg0)
}

// ConfirmedRecord is a free data retrieval call binding the contract method 0x93e6ad97.
//
// Solidity: function confirmedRecord(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) ConfirmedRecord(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.ConfirmedRecord(&_Bridge.CallOpts, arg0)
}

// ExecutionRecord is a free data retrieval call binding the contract method 0x4c906fea.
//
// Solidity: function executionRecord(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) ExecutionRecord(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "executionRecord", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecutionRecord is a free data retrieval call binding the contract method 0x4c906fea.
//
// Solidity: function executionRecord(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) ExecutionRecord(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.ExecutionRecord(&_Bridge.CallOpts, arg0)
}

// ExecutionRecord is a free data retrieval call binding the contract method 0x4c906fea.
//
// Solidity: function executionRecord(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) ExecutionRecord(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.ExecutionRecord(&_Bridge.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// LocalNonce is a free data retrieval call binding the contract method 0x13cb3591.
//
// Solidity: function localNonce() view returns(uint256)
func (_Bridge *BridgeCaller) LocalNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "localNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalNonce is a free data retrieval call binding the contract method 0x13cb3591.
//
// Solidity: function localNonce() view returns(uint256)
func (_Bridge *BridgeSession) LocalNonce() (*big.Int, error) {
	return _Bridge.Contract.LocalNonce(&_Bridge.CallOpts)
}

// LocalNonce is a free data retrieval call binding the contract method 0x13cb3591.
//
// Solidity: function localNonce() view returns(uint256)
func (_Bridge *BridgeCallerSession) LocalNonce() (*big.Int, error) {
	return _Bridge.Contract.LocalNonce(&_Bridge.CallOpts)
}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint32)
func (_Bridge *BridgeCaller) SignatureThreshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "signatureThreshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint32)
func (_Bridge *BridgeSession) SignatureThreshold() (uint32, error) {
	return _Bridge.Contract.SignatureThreshold(&_Bridge.CallOpts)
}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint32)
func (_Bridge *BridgeCallerSession) SignatureThreshold() (uint32, error) {
	return _Bridge.Contract.SignatureThreshold(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// AdminSetChainId is a paid mutator transaction binding the contract method 0x54549c4d.
//
// Solidity: function adminSetChainId(uint32 _chainId) returns()
func (_Bridge *BridgeTransactor) AdminSetChainId(opts *bind.TransactOpts, _chainId uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetChainId", _chainId)
}

// AdminSetChainId is a paid mutator transaction binding the contract method 0x54549c4d.
//
// Solidity: function adminSetChainId(uint32 _chainId) returns()
func (_Bridge *BridgeSession) AdminSetChainId(_chainId uint32) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetChainId(&_Bridge.TransactOpts, _chainId)
}

// AdminSetChainId is a paid mutator transaction binding the contract method 0x54549c4d.
//
// Solidity: function adminSetChainId(uint32 _chainId) returns()
func (_Bridge *BridgeTransactorSession) AdminSetChainId(_chainId uint32) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetChainId(&_Bridge.TransactOpts, _chainId)
}

// AdminSetConfigResource is a paid mutator transaction binding the contract method 0xc9cfe45a.
//
// Solidity: function adminSetConfigResource(bytes32 resourceID, uint256 _fee, address[] blackList) returns()
func (_Bridge *BridgeTransactor) AdminSetConfigResource(opts *bind.TransactOpts, resourceID [32]byte, _fee *big.Int, blackList []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetConfigResource", resourceID, _fee, blackList)
}

// AdminSetConfigResource is a paid mutator transaction binding the contract method 0xc9cfe45a.
//
// Solidity: function adminSetConfigResource(bytes32 resourceID, uint256 _fee, address[] blackList) returns()
func (_Bridge *BridgeSession) AdminSetConfigResource(resourceID [32]byte, _fee *big.Int, blackList []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetConfigResource(&_Bridge.TransactOpts, resourceID, _fee, blackList)
}

// AdminSetConfigResource is a paid mutator transaction binding the contract method 0xc9cfe45a.
//
// Solidity: function adminSetConfigResource(bytes32 resourceID, uint256 _fee, address[] blackList) returns()
func (_Bridge *BridgeTransactorSession) AdminSetConfigResource(resourceID [32]byte, _fee *big.Int, blackList []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetConfigResource(&_Bridge.TransactOpts, resourceID, _fee, blackList)
}

// AdminSetSignatureThreshold is a paid mutator transaction binding the contract method 0x669ce802.
//
// Solidity: function adminSetSignatureThreshold(uint32 _signatureThreshold) returns()
func (_Bridge *BridgeTransactor) AdminSetSignatureThreshold(opts *bind.TransactOpts, _signatureThreshold uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetSignatureThreshold", _signatureThreshold)
}

// AdminSetSignatureThreshold is a paid mutator transaction binding the contract method 0x669ce802.
//
// Solidity: function adminSetSignatureThreshold(uint32 _signatureThreshold) returns()
func (_Bridge *BridgeSession) AdminSetSignatureThreshold(_signatureThreshold uint32) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetSignatureThreshold(&_Bridge.TransactOpts, _signatureThreshold)
}

// AdminSetSignatureThreshold is a paid mutator transaction binding the contract method 0x669ce802.
//
// Solidity: function adminSetSignatureThreshold(uint32 _signatureThreshold) returns()
func (_Bridge *BridgeTransactorSession) AdminSetSignatureThreshold(_signatureThreshold uint32) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetSignatureThreshold(&_Bridge.TransactOpts, _signatureThreshold)
}

// CallRemote is a paid mutator transaction binding the contract method 0xb0aaa153.
//
// Solidity: function callRemote(bytes32 resourceID, uint32 targetChainId, address target, bytes data) payable returns(bytes32)
func (_Bridge *BridgeTransactor) CallRemote(opts *bind.TransactOpts, resourceID [32]byte, targetChainId uint32, target common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "callRemote", resourceID, targetChainId, target, data)
}

// CallRemote is a paid mutator transaction binding the contract method 0xb0aaa153.
//
// Solidity: function callRemote(bytes32 resourceID, uint32 targetChainId, address target, bytes data) payable returns(bytes32)
func (_Bridge *BridgeSession) CallRemote(resourceID [32]byte, targetChainId uint32, target common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.CallRemote(&_Bridge.TransactOpts, resourceID, targetChainId, target, data)
}

// CallRemote is a paid mutator transaction binding the contract method 0xb0aaa153.
//
// Solidity: function callRemote(bytes32 resourceID, uint32 targetChainId, address target, bytes data) payable returns(bytes32)
func (_Bridge *BridgeTransactorSession) CallRemote(resourceID [32]byte, targetChainId uint32, target common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.CallRemote(&_Bridge.TransactOpts, resourceID, targetChainId, target, data)
}

// Execute is a paid mutator transaction binding the contract method 0xe24c4203.
//
// Solidity: function execute(bytes32 resourceID, uint32 sourceChainId, uint256 sourceNonce, bytes32 messageId, uint32 targetChainId, address target, bytes data, bytes[] signatures) returns()
func (_Bridge *BridgeTransactor) Execute(opts *bind.TransactOpts, resourceID [32]byte, sourceChainId uint32, sourceNonce *big.Int, messageId [32]byte, targetChainId uint32, target common.Address, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "execute", resourceID, sourceChainId, sourceNonce, messageId, targetChainId, target, data, signatures)
}

// Execute is a paid mutator transaction binding the contract method 0xe24c4203.
//
// Solidity: function execute(bytes32 resourceID, uint32 sourceChainId, uint256 sourceNonce, bytes32 messageId, uint32 targetChainId, address target, bytes data, bytes[] signatures) returns()
func (_Bridge *BridgeSession) Execute(resourceID [32]byte, sourceChainId uint32, sourceNonce *big.Int, messageId [32]byte, targetChainId uint32, target common.Address, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Execute(&_Bridge.TransactOpts, resourceID, sourceChainId, sourceNonce, messageId, targetChainId, target, data, signatures)
}

// Execute is a paid mutator transaction binding the contract method 0xe24c4203.
//
// Solidity: function execute(bytes32 resourceID, uint32 sourceChainId, uint256 sourceNonce, bytes32 messageId, uint32 targetChainId, address target, bytes data, bytes[] signatures) returns()
func (_Bridge *BridgeTransactorSession) Execute(resourceID [32]byte, sourceChainId uint32, sourceNonce *big.Int, messageId [32]byte, targetChainId uint32, target common.Address, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Execute(&_Bridge.TransactOpts, resourceID, sourceChainId, sourceNonce, messageId, targetChainId, target, data, signatures)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// BridgeCallRequestIterator is returned from FilterCallRequest and is used to iterate over the raw logs and unpacked data for CallRequest events raised by the Bridge contract.
type BridgeCallRequestIterator struct {
	Event *BridgeCallRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeCallRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCallRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeCallRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeCallRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCallRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCallRequest represents a CallRequest event raised by the Bridge contract.
type BridgeCallRequest struct {
	ResourceID    [32]byte
	Caller        common.Address
	SourceChainId uint32
	SourceNonce   *big.Int
	MessageId     [32]byte
	TargetChainId uint32
	Target        common.Address
	Data          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCallRequest is a free log retrieval operation binding the contract event 0x402ecf32b0067916f0ca429f48aa67733a56593a33df268550b2d21129cc991e.
//
// Solidity: event CallRequest(bytes32 resourceID, address caller, uint32 sourceChainId, uint256 sourceNonce, bytes32 indexed messageId, uint32 indexed targetChainId, address indexed target, bytes data)
func (_Bridge *BridgeFilterer) FilterCallRequest(opts *bind.FilterOpts, messageId [][32]byte, targetChainId []uint32, target []common.Address) (*BridgeCallRequestIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CallRequest", messageIdRule, targetChainIdRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCallRequestIterator{contract: _Bridge.contract, event: "CallRequest", logs: logs, sub: sub}, nil
}

// WatchCallRequest is a free log subscription operation binding the contract event 0x402ecf32b0067916f0ca429f48aa67733a56593a33df268550b2d21129cc991e.
//
// Solidity: event CallRequest(bytes32 resourceID, address caller, uint32 sourceChainId, uint256 sourceNonce, bytes32 indexed messageId, uint32 indexed targetChainId, address indexed target, bytes data)
func (_Bridge *BridgeFilterer) WatchCallRequest(opts *bind.WatchOpts, sink chan<- *BridgeCallRequest, messageId [][32]byte, targetChainId []uint32, target []common.Address) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CallRequest", messageIdRule, targetChainIdRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCallRequest)
				if err := _Bridge.contract.UnpackLog(event, "CallRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallRequest is a log parse operation binding the contract event 0x402ecf32b0067916f0ca429f48aa67733a56593a33df268550b2d21129cc991e.
//
// Solidity: event CallRequest(bytes32 resourceID, address caller, uint32 sourceChainId, uint256 sourceNonce, bytes32 indexed messageId, uint32 indexed targetChainId, address indexed target, bytes data)
func (_Bridge *BridgeFilterer) ParseCallRequest(log types.Log) (*BridgeCallRequest, error) {
	event := new(BridgeCallRequest)
	if err := _Bridge.contract.UnpackLog(event, "CallRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeConfirmedRequestIterator is returned from FilterConfirmedRequest and is used to iterate over the raw logs and unpacked data for ConfirmedRequest events raised by the Bridge contract.
type BridgeConfirmedRequestIterator struct {
	Event *BridgeConfirmedRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeConfirmedRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeConfirmedRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeConfirmedRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeConfirmedRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeConfirmedRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeConfirmedRequest represents a ConfirmedRequest event raised by the Bridge contract.
type BridgeConfirmedRequest struct {
	ResourceID    [32]byte
	Caller        common.Address
	SourceChainId uint32
	SourceNonce   *big.Int
	MessageId     [32]byte
	TargetChainId uint32
	Target        common.Address
	Data          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmedRequest is a free log retrieval operation binding the contract event 0xe1678c0796b64c840da197b535dfd9e7e989768fbf5afd267d31c664a1f30e32.
//
// Solidity: event ConfirmedRequest(bytes32 resourceID, address caller, uint32 sourceChainId, uint256 sourceNonce, bytes32 indexed messageId, uint32 indexed targetChainId, address indexed target, bytes data)
func (_Bridge *BridgeFilterer) FilterConfirmedRequest(opts *bind.FilterOpts, messageId [][32]byte, targetChainId []uint32, target []common.Address) (*BridgeConfirmedRequestIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ConfirmedRequest", messageIdRule, targetChainIdRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeConfirmedRequestIterator{contract: _Bridge.contract, event: "ConfirmedRequest", logs: logs, sub: sub}, nil
}

// WatchConfirmedRequest is a free log subscription operation binding the contract event 0xe1678c0796b64c840da197b535dfd9e7e989768fbf5afd267d31c664a1f30e32.
//
// Solidity: event ConfirmedRequest(bytes32 resourceID, address caller, uint32 sourceChainId, uint256 sourceNonce, bytes32 indexed messageId, uint32 indexed targetChainId, address indexed target, bytes data)
func (_Bridge *BridgeFilterer) WatchConfirmedRequest(opts *bind.WatchOpts, sink chan<- *BridgeConfirmedRequest, messageId [][32]byte, targetChainId []uint32, target []common.Address) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var targetChainIdRule []interface{}
	for _, targetChainIdItem := range targetChainId {
		targetChainIdRule = append(targetChainIdRule, targetChainIdItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ConfirmedRequest", messageIdRule, targetChainIdRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeConfirmedRequest)
				if err := _Bridge.contract.UnpackLog(event, "ConfirmedRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedRequest is a log parse operation binding the contract event 0xe1678c0796b64c840da197b535dfd9e7e989768fbf5afd267d31c664a1f30e32.
//
// Solidity: event ConfirmedRequest(bytes32 resourceID, address caller, uint32 sourceChainId, uint256 sourceNonce, bytes32 indexed messageId, uint32 indexed targetChainId, address indexed target, bytes data)
func (_Bridge *BridgeFilterer) ParseConfirmedRequest(log types.Log) (*BridgeConfirmedRequest, error) {
	event := new(BridgeConfirmedRequest)
	if err := _Bridge.contract.UnpackLog(event, "ConfirmedRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridge contract.
type BridgeRoleAdminChangedIterator struct {
	Event *BridgeRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Bridge contract.
type BridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleAdminChangedIterator{contract: _Bridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRoleAdminChanged, error) {
	event := new(BridgeRoleAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
