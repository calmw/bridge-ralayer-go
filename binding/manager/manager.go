// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package manager

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

// ManagerResourceConfig is an auto generated low-level Go binding around an user-defined struct.
type ManagerResourceConfig struct {
	RemoteCallType uint8
}

// ManagerMetaData contains all meta data concerning the Manager contract.
var ManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_voteThreshold\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_expiredBlockNum\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIManager.VoteStatus\",\"name\":\"voteStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"SignatureCollected\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ExpiredBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"automaticCall\",\"type\":\"bool\"}],\"name\":\"adminSetConfigResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_voteThreshold\",\"type\":\"uint32\"}],\"name\":\"adminSetVoteThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"chainConfigResource\",\"outputs\":[{\"internalType\":\"enumManager.RemoteCallType\",\"name\":\"remoteCallType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"getChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"enumManager.RemoteCallType\",\"name\":\"remoteCallType\",\"type\":\"uint8\"}],\"internalType\":\"structManager.ResourceConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"hasVotedOnMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sourceNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voteRecords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"enumIManager.VoteStatus\",\"name\":\"voteStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sourceNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"targetChainId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteThreshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ManagerMetaData.ABI instead.
var ManagerABI = ManagerMetaData.ABI

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Manager *ManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Manager *ManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Manager.Contract.DEFAULTADMINROLE(&_Manager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Manager *ManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Manager.Contract.DEFAULTADMINROLE(&_Manager.CallOpts)
}

// ExpiredBlockNum is a free data retrieval call binding the contract method 0x1f9dcb69.
//
// Solidity: function ExpiredBlockNum() view returns(uint256)
func (_Manager *ManagerCaller) ExpiredBlockNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "ExpiredBlockNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiredBlockNum is a free data retrieval call binding the contract method 0x1f9dcb69.
//
// Solidity: function ExpiredBlockNum() view returns(uint256)
func (_Manager *ManagerSession) ExpiredBlockNum() (*big.Int, error) {
	return _Manager.Contract.ExpiredBlockNum(&_Manager.CallOpts)
}

// ExpiredBlockNum is a free data retrieval call binding the contract method 0x1f9dcb69.
//
// Solidity: function ExpiredBlockNum() view returns(uint256)
func (_Manager *ManagerCallerSession) ExpiredBlockNum() (*big.Int, error) {
	return _Manager.Contract.ExpiredBlockNum(&_Manager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Manager *ManagerCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Manager *ManagerSession) MANAGERROLE() ([32]byte, error) {
	return _Manager.Contract.MANAGERROLE(&_Manager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Manager *ManagerCallerSession) MANAGERROLE() ([32]byte, error) {
	return _Manager.Contract.MANAGERROLE(&_Manager.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Manager *ManagerCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Manager *ManagerSession) VALIDATORROLE() ([32]byte, error) {
	return _Manager.Contract.VALIDATORROLE(&_Manager.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Manager *ManagerCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _Manager.Contract.VALIDATORROLE(&_Manager.CallOpts)
}

// ChainConfigResource is a free data retrieval call binding the contract method 0x099eaed6.
//
// Solidity: function chainConfigResource(bytes32 ) view returns(uint8 remoteCallType)
func (_Manager *ManagerCaller) ChainConfigResource(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "chainConfigResource", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChainConfigResource is a free data retrieval call binding the contract method 0x099eaed6.
//
// Solidity: function chainConfigResource(bytes32 ) view returns(uint8 remoteCallType)
func (_Manager *ManagerSession) ChainConfigResource(arg0 [32]byte) (uint8, error) {
	return _Manager.Contract.ChainConfigResource(&_Manager.CallOpts, arg0)
}

// ChainConfigResource is a free data retrieval call binding the contract method 0x099eaed6.
//
// Solidity: function chainConfigResource(bytes32 ) view returns(uint8 remoteCallType)
func (_Manager *ManagerCallerSession) ChainConfigResource(arg0 [32]byte) (uint8, error) {
	return _Manager.Contract.ChainConfigResource(&_Manager.CallOpts, arg0)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x2c55235e.
//
// Solidity: function getChainConfig(bytes32 resourceID) view returns((uint8))
func (_Manager *ManagerCaller) GetChainConfig(opts *bind.CallOpts, resourceID [32]byte) (ManagerResourceConfig, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "getChainConfig", resourceID)

	if err != nil {
		return *new(ManagerResourceConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ManagerResourceConfig)).(*ManagerResourceConfig)

	return out0, err

}

// GetChainConfig is a free data retrieval call binding the contract method 0x2c55235e.
//
// Solidity: function getChainConfig(bytes32 resourceID) view returns((uint8))
func (_Manager *ManagerSession) GetChainConfig(resourceID [32]byte) (ManagerResourceConfig, error) {
	return _Manager.Contract.GetChainConfig(&_Manager.CallOpts, resourceID)
}

// GetChainConfig is a free data retrieval call binding the contract method 0x2c55235e.
//
// Solidity: function getChainConfig(bytes32 resourceID) view returns((uint8))
func (_Manager *ManagerCallerSession) GetChainConfig(resourceID [32]byte) (ManagerResourceConfig, error) {
	return _Manager.Contract.GetChainConfig(&_Manager.CallOpts, resourceID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Manager *ManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Manager *ManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Manager.Contract.GetRoleAdmin(&_Manager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Manager *ManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Manager.Contract.GetRoleAdmin(&_Manager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Manager *ManagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Manager *ManagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Manager.Contract.GetRoleMember(&_Manager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Manager *ManagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Manager.Contract.GetRoleMember(&_Manager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Manager *ManagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Manager *ManagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Manager.Contract.GetRoleMemberCount(&_Manager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Manager *ManagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Manager.Contract.GetRoleMemberCount(&_Manager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Manager *ManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Manager *ManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Manager.Contract.HasRole(&_Manager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Manager *ManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Manager.Contract.HasRole(&_Manager.CallOpts, role, account)
}

// HasVotedOnMessage is a free data retrieval call binding the contract method 0x8eb77eeb.
//
// Solidity: function hasVotedOnMessage(bytes32 messageId) view returns(bool)
func (_Manager *ManagerCaller) HasVotedOnMessage(opts *bind.CallOpts, messageId [32]byte) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "hasVotedOnMessage", messageId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnMessage is a free data retrieval call binding the contract method 0x8eb77eeb.
//
// Solidity: function hasVotedOnMessage(bytes32 messageId) view returns(bool)
func (_Manager *ManagerSession) HasVotedOnMessage(messageId [32]byte) (bool, error) {
	return _Manager.Contract.HasVotedOnMessage(&_Manager.CallOpts, messageId)
}

// HasVotedOnMessage is a free data retrieval call binding the contract method 0x8eb77eeb.
//
// Solidity: function hasVotedOnMessage(bytes32 messageId) view returns(bool)
func (_Manager *ManagerCallerSession) HasVotedOnMessage(messageId [32]byte) (bool, error) {
	return _Manager.Contract.HasVotedOnMessage(&_Manager.CallOpts, messageId)
}

// NonceMap is a free data retrieval call binding the contract method 0x96d5e2bd.
//
// Solidity: function nonceMap(uint32 ) view returns(uint256)
func (_Manager *ManagerCaller) NonceMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "nonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceMap is a free data retrieval call binding the contract method 0x96d5e2bd.
//
// Solidity: function nonceMap(uint32 ) view returns(uint256)
func (_Manager *ManagerSession) NonceMap(arg0 uint32) (*big.Int, error) {
	return _Manager.Contract.NonceMap(&_Manager.CallOpts, arg0)
}

// NonceMap is a free data retrieval call binding the contract method 0x96d5e2bd.
//
// Solidity: function nonceMap(uint32 ) view returns(uint256)
func (_Manager *ManagerCallerSession) NonceMap(arg0 uint32) (*big.Int, error) {
	return _Manager.Contract.NonceMap(&_Manager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Manager *ManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Manager *ManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Manager.Contract.SupportsInterface(&_Manager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Manager *ManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Manager.Contract.SupportsInterface(&_Manager.CallOpts, interfaceId)
}

// VoteRecords is a free data retrieval call binding the contract method 0x79916eb4.
//
// Solidity: function voteRecords(bytes32 ) view returns(bytes32 resourceID, uint8 voteStatus, uint256 startBlock, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, bytes32 dataHash)
func (_Manager *ManagerCaller) VoteRecords(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ResourceID    [32]byte
	VoteStatus    uint8
	StartBlock    *big.Int
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	DataHash      [32]byte
}, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "voteRecords", arg0)

	outstruct := new(struct {
		ResourceID    [32]byte
		VoteStatus    uint8
		StartBlock    *big.Int
		SourceChainId uint32
		SourceNonce   *big.Int
		TargetChainId uint32
		DataHash      [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteStatus = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.StartBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SourceChainId = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.SourceNonce = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TargetChainId = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.DataHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// VoteRecords is a free data retrieval call binding the contract method 0x79916eb4.
//
// Solidity: function voteRecords(bytes32 ) view returns(bytes32 resourceID, uint8 voteStatus, uint256 startBlock, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, bytes32 dataHash)
func (_Manager *ManagerSession) VoteRecords(arg0 [32]byte) (struct {
	ResourceID    [32]byte
	VoteStatus    uint8
	StartBlock    *big.Int
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	DataHash      [32]byte
}, error) {
	return _Manager.Contract.VoteRecords(&_Manager.CallOpts, arg0)
}

// VoteRecords is a free data retrieval call binding the contract method 0x79916eb4.
//
// Solidity: function voteRecords(bytes32 ) view returns(bytes32 resourceID, uint8 voteStatus, uint256 startBlock, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, bytes32 dataHash)
func (_Manager *ManagerCallerSession) VoteRecords(arg0 [32]byte) (struct {
	ResourceID    [32]byte
	VoteStatus    uint8
	StartBlock    *big.Int
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	DataHash      [32]byte
}, error) {
	return _Manager.Contract.VoteRecords(&_Manager.CallOpts, arg0)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint32)
func (_Manager *ManagerCaller) VoteThreshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "voteThreshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint32)
func (_Manager *ManagerSession) VoteThreshold() (uint32, error) {
	return _Manager.Contract.VoteThreshold(&_Manager.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint32)
func (_Manager *ManagerCallerSession) VoteThreshold() (uint32, error) {
	return _Manager.Contract.VoteThreshold(&_Manager.CallOpts)
}

// AdminSetConfigResource is a paid mutator transaction binding the contract method 0x2ae08b59.
//
// Solidity: function adminSetConfigResource(bytes32 resourceID, bool automaticCall) returns()
func (_Manager *ManagerTransactor) AdminSetConfigResource(opts *bind.TransactOpts, resourceID [32]byte, automaticCall bool) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "adminSetConfigResource", resourceID, automaticCall)
}

// AdminSetConfigResource is a paid mutator transaction binding the contract method 0x2ae08b59.
//
// Solidity: function adminSetConfigResource(bytes32 resourceID, bool automaticCall) returns()
func (_Manager *ManagerSession) AdminSetConfigResource(resourceID [32]byte, automaticCall bool) (*types.Transaction, error) {
	return _Manager.Contract.AdminSetConfigResource(&_Manager.TransactOpts, resourceID, automaticCall)
}

// AdminSetConfigResource is a paid mutator transaction binding the contract method 0x2ae08b59.
//
// Solidity: function adminSetConfigResource(bytes32 resourceID, bool automaticCall) returns()
func (_Manager *ManagerTransactorSession) AdminSetConfigResource(resourceID [32]byte, automaticCall bool) (*types.Transaction, error) {
	return _Manager.Contract.AdminSetConfigResource(&_Manager.TransactOpts, resourceID, automaticCall)
}

// AdminSetVoteThreshold is a paid mutator transaction binding the contract method 0x44d38d0b.
//
// Solidity: function adminSetVoteThreshold(uint32 _voteThreshold) returns()
func (_Manager *ManagerTransactor) AdminSetVoteThreshold(opts *bind.TransactOpts, _voteThreshold uint32) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "adminSetVoteThreshold", _voteThreshold)
}

// AdminSetVoteThreshold is a paid mutator transaction binding the contract method 0x44d38d0b.
//
// Solidity: function adminSetVoteThreshold(uint32 _voteThreshold) returns()
func (_Manager *ManagerSession) AdminSetVoteThreshold(_voteThreshold uint32) (*types.Transaction, error) {
	return _Manager.Contract.AdminSetVoteThreshold(&_Manager.TransactOpts, _voteThreshold)
}

// AdminSetVoteThreshold is a paid mutator transaction binding the contract method 0x44d38d0b.
//
// Solidity: function adminSetVoteThreshold(uint32 _voteThreshold) returns()
func (_Manager *ManagerTransactorSession) AdminSetVoteThreshold(_voteThreshold uint32) (*types.Transaction, error) {
	return _Manager.Contract.AdminSetVoteThreshold(&_Manager.TransactOpts, _voteThreshold)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Manager *ManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Manager *ManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.Contract.GrantRole(&_Manager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Manager *ManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.Contract.GrantRole(&_Manager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Manager *ManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Manager *ManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.Contract.RenounceRole(&_Manager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Manager *ManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.Contract.RenounceRole(&_Manager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Manager *ManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Manager *ManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.Contract.RevokeRole(&_Manager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Manager *ManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Manager.Contract.RevokeRole(&_Manager.TransactOpts, role, account)
}

// Vote is a paid mutator transaction binding the contract method 0x35fdedbe.
//
// Solidity: function vote(bytes32 _resourceID, bytes32 messageId, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, address target, bytes32 dataHash, bytes signature) returns()
func (_Manager *ManagerTransactor) Vote(opts *bind.TransactOpts, _resourceID [32]byte, messageId [32]byte, sourceChainId uint32, sourceNonce *big.Int, targetChainId uint32, target common.Address, dataHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "vote", _resourceID, messageId, sourceChainId, sourceNonce, targetChainId, target, dataHash, signature)
}

// Vote is a paid mutator transaction binding the contract method 0x35fdedbe.
//
// Solidity: function vote(bytes32 _resourceID, bytes32 messageId, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, address target, bytes32 dataHash, bytes signature) returns()
func (_Manager *ManagerSession) Vote(_resourceID [32]byte, messageId [32]byte, sourceChainId uint32, sourceNonce *big.Int, targetChainId uint32, target common.Address, dataHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _Manager.Contract.Vote(&_Manager.TransactOpts, _resourceID, messageId, sourceChainId, sourceNonce, targetChainId, target, dataHash, signature)
}

// Vote is a paid mutator transaction binding the contract method 0x35fdedbe.
//
// Solidity: function vote(bytes32 _resourceID, bytes32 messageId, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, address target, bytes32 dataHash, bytes signature) returns()
func (_Manager *ManagerTransactorSession) Vote(_resourceID [32]byte, messageId [32]byte, sourceChainId uint32, sourceNonce *big.Int, targetChainId uint32, target common.Address, dataHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _Manager.Contract.Vote(&_Manager.TransactOpts, _resourceID, messageId, sourceChainId, sourceNonce, targetChainId, target, dataHash, signature)
}

// ManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Manager contract.
type ManagerRoleAdminChangedIterator struct {
	Event *ManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerRoleAdminChanged)
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
		it.Event = new(ManagerRoleAdminChanged)
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
func (it *ManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerRoleAdminChanged represents a RoleAdminChanged event raised by the Manager contract.
type ManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Manager *ManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Manager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ManagerRoleAdminChangedIterator{contract: _Manager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Manager *ManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Manager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerRoleAdminChanged)
				if err := _Manager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseRoleAdminChanged(log types.Log) (*ManagerRoleAdminChanged, error) {
	event := new(ManagerRoleAdminChanged)
	if err := _Manager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Manager contract.
type ManagerRoleGrantedIterator struct {
	Event *ManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *ManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerRoleGranted)
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
		it.Event = new(ManagerRoleGranted)
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
func (it *ManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerRoleGranted represents a RoleGranted event raised by the Manager contract.
type ManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Manager *ManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _Manager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ManagerRoleGrantedIterator{contract: _Manager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Manager *ManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Manager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerRoleGranted)
				if err := _Manager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseRoleGranted(log types.Log) (*ManagerRoleGranted, error) {
	event := new(ManagerRoleGranted)
	if err := _Manager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Manager contract.
type ManagerRoleRevokedIterator struct {
	Event *ManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerRoleRevoked)
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
		it.Event = new(ManagerRoleRevoked)
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
func (it *ManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerRoleRevoked represents a RoleRevoked event raised by the Manager contract.
type ManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Manager *ManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _Manager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ManagerRoleRevokedIterator{contract: _Manager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Manager *ManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Manager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerRoleRevoked)
				if err := _Manager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseRoleRevoked(log types.Log) (*ManagerRoleRevoked, error) {
	event := new(ManagerRoleRevoked)
	if err := _Manager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerSignatureCollectedIterator is returned from FilterSignatureCollected and is used to iterate over the raw logs and unpacked data for SignatureCollected events raised by the Manager contract.
type ManagerSignatureCollectedIterator struct {
	Event *ManagerSignatureCollected // Event containing the contract specifics and raw log

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
func (it *ManagerSignatureCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerSignatureCollected)
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
		it.Event = new(ManagerSignatureCollected)
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
func (it *ManagerSignatureCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerSignatureCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerSignatureCollected represents a SignatureCollected event raised by the Manager contract.
type ManagerSignatureCollected struct {
	ResourceID    [32]byte
	VoteStatus    uint8
	SourceChainId uint32
	SourceNonce   *big.Int
	TargetChainId uint32
	Target        common.Address
	MessageId     [32]byte
	DataHash      [32]byte
	Signatures    [][]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignatureCollected is a free log retrieval operation binding the contract event 0x5b7df0ed48b3444dd9580aff59ede15690a58a3ddaa3bba555b5fd20be3c7e0b.
//
// Solidity: event SignatureCollected(bytes32 resourceID, uint8 voteStatus, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, address target, bytes32 indexed messageId, bytes32 indexed dataHash, bytes[] signatures)
func (_Manager *ManagerFilterer) FilterSignatureCollected(opts *bind.FilterOpts, messageId [][32]byte, dataHash [][32]byte) (*ManagerSignatureCollectedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "SignatureCollected", messageIdRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return &ManagerSignatureCollectedIterator{contract: _Manager.contract, event: "SignatureCollected", logs: logs, sub: sub}, nil
}

// WatchSignatureCollected is a free log subscription operation binding the contract event 0x5b7df0ed48b3444dd9580aff59ede15690a58a3ddaa3bba555b5fd20be3c7e0b.
//
// Solidity: event SignatureCollected(bytes32 resourceID, uint8 voteStatus, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, address target, bytes32 indexed messageId, bytes32 indexed dataHash, bytes[] signatures)
func (_Manager *ManagerFilterer) WatchSignatureCollected(opts *bind.WatchOpts, sink chan<- *ManagerSignatureCollected, messageId [][32]byte, dataHash [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "SignatureCollected", messageIdRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerSignatureCollected)
				if err := _Manager.contract.UnpackLog(event, "SignatureCollected", log); err != nil {
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

// ParseSignatureCollected is a log parse operation binding the contract event 0x5b7df0ed48b3444dd9580aff59ede15690a58a3ddaa3bba555b5fd20be3c7e0b.
//
// Solidity: event SignatureCollected(bytes32 resourceID, uint8 voteStatus, uint32 sourceChainId, uint256 sourceNonce, uint32 targetChainId, address target, bytes32 indexed messageId, bytes32 indexed dataHash, bytes[] signatures)
func (_Manager *ManagerFilterer) ParseSignatureCollected(log types.Log) (*ManagerSignatureCollected, error) {
	event := new(ManagerSignatureCollected)
	if err := _Manager.contract.UnpackLog(event, "SignatureCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
