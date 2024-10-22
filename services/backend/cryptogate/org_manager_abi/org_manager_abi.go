// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package org_manager_abi

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
	_ = abi.ConvertType
)

// OrgManagerAbiMetaData contains all meta data concerning the OrgManagerAbi contract.
var OrgManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"OrgApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"OrgPendingApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"OrgSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"OrgSuspensionRevoked\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pOrgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"approveOrgStatusUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"checkOrgActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"checkOrgExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_orgStatus\",\"type\":\"uint256\"}],\"name\":\"checkOrgStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfOrgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getOrgDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orgIndex\",\"type\":\"uint256\"}],\"name\":\"getOrgInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getSubOrgIndexes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getUltimateParent\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_breadth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depth\",\"type\":\"uint256\"}],\"name\":\"setUpOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateOrg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrgManagerAbi is an auto generated Go binding around an Ethereum contract.
type OrgManagerAbi struct {
	OrgManagerAbiCaller     // Read-only binding to the contract
	OrgManagerAbiTransactor // Write-only binding to the contract
	OrgManagerAbiFilterer   // Log filterer for contract events
}

// OrgManagerAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrgManagerAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgManagerAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrgManagerAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgManagerAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrgManagerAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgManagerAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrgManagerAbiSession struct {
	Contract     *OrgManagerAbi    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgManagerAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrgManagerAbiCallerSession struct {
	Contract *OrgManagerAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OrgManagerAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrgManagerAbiTransactorSession struct {
	Contract     *OrgManagerAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OrgManagerAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrgManagerAbiRaw struct {
	Contract *OrgManagerAbi // Generic contract binding to access the raw methods on
}

// OrgManagerAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrgManagerAbiCallerRaw struct {
	Contract *OrgManagerAbiCaller // Generic read-only contract binding to access the raw methods on
}

// OrgManagerAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrgManagerAbiTransactorRaw struct {
	Contract *OrgManagerAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrgManagerAbi creates a new instance of OrgManagerAbi, bound to a specific deployed contract.
func NewOrgManagerAbi(address common.Address, backend bind.ContractBackend) (*OrgManagerAbi, error) {
	contract, err := bindOrgManagerAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbi{OrgManagerAbiCaller: OrgManagerAbiCaller{contract: contract}, OrgManagerAbiTransactor: OrgManagerAbiTransactor{contract: contract}, OrgManagerAbiFilterer: OrgManagerAbiFilterer{contract: contract}}, nil
}

// NewOrgManagerAbiCaller creates a new read-only instance of OrgManagerAbi, bound to a specific deployed contract.
func NewOrgManagerAbiCaller(address common.Address, caller bind.ContractCaller) (*OrgManagerAbiCaller, error) {
	contract, err := bindOrgManagerAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiCaller{contract: contract}, nil
}

// NewOrgManagerAbiTransactor creates a new write-only instance of OrgManagerAbi, bound to a specific deployed contract.
func NewOrgManagerAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*OrgManagerAbiTransactor, error) {
	contract, err := bindOrgManagerAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiTransactor{contract: contract}, nil
}

// NewOrgManagerAbiFilterer creates a new log filterer instance of OrgManagerAbi, bound to a specific deployed contract.
func NewOrgManagerAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*OrgManagerAbiFilterer, error) {
	contract, err := bindOrgManagerAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiFilterer{contract: contract}, nil
}

// bindOrgManagerAbi binds a generic wrapper to an already deployed contract.
func bindOrgManagerAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrgManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgManagerAbi *OrgManagerAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrgManagerAbi.Contract.OrgManagerAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgManagerAbi *OrgManagerAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.OrgManagerAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgManagerAbi *OrgManagerAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.OrgManagerAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgManagerAbi *OrgManagerAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrgManagerAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgManagerAbi *OrgManagerAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgManagerAbi *OrgManagerAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.contract.Transact(opts, method, params...)
}

// CheckOrgActive is a free data retrieval call binding the contract method 0x3fd62ae7.
//
// Solidity: function checkOrgActive(string _orgId) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiCaller) CheckOrgActive(opts *bind.CallOpts, _orgId string) (bool, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "checkOrgActive", _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOrgActive is a free data retrieval call binding the contract method 0x3fd62ae7.
//
// Solidity: function checkOrgActive(string _orgId) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiSession) CheckOrgActive(_orgId string) (bool, error) {
	return _OrgManagerAbi.Contract.CheckOrgActive(&_OrgManagerAbi.CallOpts, _orgId)
}

// CheckOrgActive is a free data retrieval call binding the contract method 0x3fd62ae7.
//
// Solidity: function checkOrgActive(string _orgId) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) CheckOrgActive(_orgId string) (bool, error) {
	return _OrgManagerAbi.Contract.CheckOrgActive(&_OrgManagerAbi.CallOpts, _orgId)
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(string _orgId) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiCaller) CheckOrgExists(opts *bind.CallOpts, _orgId string) (bool, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "checkOrgExists", _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(string _orgId) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiSession) CheckOrgExists(_orgId string) (bool, error) {
	return _OrgManagerAbi.Contract.CheckOrgExists(&_OrgManagerAbi.CallOpts, _orgId)
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(string _orgId) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) CheckOrgExists(_orgId string) (bool, error) {
	return _OrgManagerAbi.Contract.CheckOrgExists(&_OrgManagerAbi.CallOpts, _orgId)
}

// CheckOrgStatus is a free data retrieval call binding the contract method 0x8c8642df.
//
// Solidity: function checkOrgStatus(string _orgId, uint256 _orgStatus) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiCaller) CheckOrgStatus(opts *bind.CallOpts, _orgId string, _orgStatus *big.Int) (bool, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "checkOrgStatus", _orgId, _orgStatus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOrgStatus is a free data retrieval call binding the contract method 0x8c8642df.
//
// Solidity: function checkOrgStatus(string _orgId, uint256 _orgStatus) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiSession) CheckOrgStatus(_orgId string, _orgStatus *big.Int) (bool, error) {
	return _OrgManagerAbi.Contract.CheckOrgStatus(&_OrgManagerAbi.CallOpts, _orgId, _orgStatus)
}

// CheckOrgStatus is a free data retrieval call binding the contract method 0x8c8642df.
//
// Solidity: function checkOrgStatus(string _orgId, uint256 _orgStatus) view returns(bool)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) CheckOrgStatus(_orgId string, _orgStatus *big.Int) (bool, error) {
	return _OrgManagerAbi.Contract.CheckOrgStatus(&_OrgManagerAbi.CallOpts, _orgId, _orgStatus)
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() view returns(uint256)
func (_OrgManagerAbi *OrgManagerAbiCaller) GetNumberOfOrgs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "getNumberOfOrgs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() view returns(uint256)
func (_OrgManagerAbi *OrgManagerAbiSession) GetNumberOfOrgs() (*big.Int, error) {
	return _OrgManagerAbi.Contract.GetNumberOfOrgs(&_OrgManagerAbi.CallOpts)
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() view returns(uint256)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) GetNumberOfOrgs() (*big.Int, error) {
	return _OrgManagerAbi.Contract.GetNumberOfOrgs(&_OrgManagerAbi.CallOpts)
}

// GetOrgDetails is a free data retrieval call binding the contract method 0xf4d6d9f5.
//
// Solidity: function getOrgDetails(string _orgId) view returns(string, string, string, uint256, uint256)
func (_OrgManagerAbi *OrgManagerAbiCaller) GetOrgDetails(opts *bind.CallOpts, _orgId string) (string, string, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "getOrgDetails", _orgId)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetOrgDetails is a free data retrieval call binding the contract method 0xf4d6d9f5.
//
// Solidity: function getOrgDetails(string _orgId) view returns(string, string, string, uint256, uint256)
func (_OrgManagerAbi *OrgManagerAbiSession) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return _OrgManagerAbi.Contract.GetOrgDetails(&_OrgManagerAbi.CallOpts, _orgId)
}

// GetOrgDetails is a free data retrieval call binding the contract method 0xf4d6d9f5.
//
// Solidity: function getOrgDetails(string _orgId) view returns(string, string, string, uint256, uint256)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return _OrgManagerAbi.Contract.GetOrgDetails(&_OrgManagerAbi.CallOpts, _orgId)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(uint256 _orgIndex) view returns(string, string, string, uint256, uint256)
func (_OrgManagerAbi *OrgManagerAbiCaller) GetOrgInfo(opts *bind.CallOpts, _orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "getOrgInfo", _orgIndex)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(uint256 _orgIndex) view returns(string, string, string, uint256, uint256)
func (_OrgManagerAbi *OrgManagerAbiSession) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return _OrgManagerAbi.Contract.GetOrgInfo(&_OrgManagerAbi.CallOpts, _orgIndex)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(uint256 _orgIndex) view returns(string, string, string, uint256, uint256)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return _OrgManagerAbi.Contract.GetOrgInfo(&_OrgManagerAbi.CallOpts, _orgIndex)
}

// GetSubOrgIndexes is a free data retrieval call binding the contract method 0x5e99f6e5.
//
// Solidity: function getSubOrgIndexes(string _orgId) view returns(uint256[])
func (_OrgManagerAbi *OrgManagerAbiCaller) GetSubOrgIndexes(opts *bind.CallOpts, _orgId string) ([]*big.Int, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "getSubOrgIndexes", _orgId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSubOrgIndexes is a free data retrieval call binding the contract method 0x5e99f6e5.
//
// Solidity: function getSubOrgIndexes(string _orgId) view returns(uint256[])
func (_OrgManagerAbi *OrgManagerAbiSession) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return _OrgManagerAbi.Contract.GetSubOrgIndexes(&_OrgManagerAbi.CallOpts, _orgId)
}

// GetSubOrgIndexes is a free data retrieval call binding the contract method 0x5e99f6e5.
//
// Solidity: function getSubOrgIndexes(string _orgId) view returns(uint256[])
func (_OrgManagerAbi *OrgManagerAbiCallerSession) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return _OrgManagerAbi.Contract.GetSubOrgIndexes(&_OrgManagerAbi.CallOpts, _orgId)
}

// GetUltimateParent is a free data retrieval call binding the contract method 0x177c8d8a.
//
// Solidity: function getUltimateParent(string _orgId) view returns(string)
func (_OrgManagerAbi *OrgManagerAbiCaller) GetUltimateParent(opts *bind.CallOpts, _orgId string) (string, error) {
	var out []interface{}
	err := _OrgManagerAbi.contract.Call(opts, &out, "getUltimateParent", _orgId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUltimateParent is a free data retrieval call binding the contract method 0x177c8d8a.
//
// Solidity: function getUltimateParent(string _orgId) view returns(string)
func (_OrgManagerAbi *OrgManagerAbiSession) GetUltimateParent(_orgId string) (string, error) {
	return _OrgManagerAbi.Contract.GetUltimateParent(&_OrgManagerAbi.CallOpts, _orgId)
}

// GetUltimateParent is a free data retrieval call binding the contract method 0x177c8d8a.
//
// Solidity: function getUltimateParent(string _orgId) view returns(string)
func (_OrgManagerAbi *OrgManagerAbiCallerSession) GetUltimateParent(_orgId string) (string, error) {
	return _OrgManagerAbi.Contract.GetUltimateParent(&_OrgManagerAbi.CallOpts, _orgId)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf9953de5.
//
// Solidity: function addOrg(string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactor) AddOrg(opts *bind.TransactOpts, _orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.contract.Transact(opts, "addOrg", _orgId)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf9953de5.
//
// Solidity: function addOrg(string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiSession) AddOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.AddOrg(&_OrgManagerAbi.TransactOpts, _orgId)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf9953de5.
//
// Solidity: function addOrg(string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactorSession) AddOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.AddOrg(&_OrgManagerAbi.TransactOpts, _orgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactor) AddSubOrg(opts *bind.TransactOpts, _pOrgId string, _orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.contract.Transact(opts, "addSubOrg", _pOrgId, _orgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiSession) AddSubOrg(_pOrgId string, _orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.AddSubOrg(&_OrgManagerAbi.TransactOpts, _pOrgId, _orgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactorSession) AddSubOrg(_pOrgId string, _orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.AddSubOrg(&_OrgManagerAbi.TransactOpts, _pOrgId, _orgId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xe3028316.
//
// Solidity: function approveOrg(string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.contract.Transact(opts, "approveOrg", _orgId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xe3028316.
//
// Solidity: function approveOrg(string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiSession) ApproveOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.ApproveOrg(&_OrgManagerAbi.TransactOpts, _orgId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xe3028316.
//
// Solidity: function approveOrg(string _orgId) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactorSession) ApproveOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.ApproveOrg(&_OrgManagerAbi.TransactOpts, _orgId)
}

// ApproveOrgStatusUpdate is a paid mutator transaction binding the contract method 0x14f775f9.
//
// Solidity: function approveOrgStatusUpdate(string _orgId, uint256 _action) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactor) ApproveOrgStatusUpdate(opts *bind.TransactOpts, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.contract.Transact(opts, "approveOrgStatusUpdate", _orgId, _action)
}

// ApproveOrgStatusUpdate is a paid mutator transaction binding the contract method 0x14f775f9.
//
// Solidity: function approveOrgStatusUpdate(string _orgId, uint256 _action) returns()
func (_OrgManagerAbi *OrgManagerAbiSession) ApproveOrgStatusUpdate(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.ApproveOrgStatusUpdate(&_OrgManagerAbi.TransactOpts, _orgId, _action)
}

// ApproveOrgStatusUpdate is a paid mutator transaction binding the contract method 0x14f775f9.
//
// Solidity: function approveOrgStatusUpdate(string _orgId, uint256 _action) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactorSession) ApproveOrgStatusUpdate(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.ApproveOrgStatusUpdate(&_OrgManagerAbi.TransactOpts, _orgId, _action)
}

// SetUpOrg is a paid mutator transaction binding the contract method 0x9e58eb9f.
//
// Solidity: function setUpOrg(string _orgId, uint256 _breadth, uint256 _depth) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactor) SetUpOrg(opts *bind.TransactOpts, _orgId string, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.contract.Transact(opts, "setUpOrg", _orgId, _breadth, _depth)
}

// SetUpOrg is a paid mutator transaction binding the contract method 0x9e58eb9f.
//
// Solidity: function setUpOrg(string _orgId, uint256 _breadth, uint256 _depth) returns()
func (_OrgManagerAbi *OrgManagerAbiSession) SetUpOrg(_orgId string, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.SetUpOrg(&_OrgManagerAbi.TransactOpts, _orgId, _breadth, _depth)
}

// SetUpOrg is a paid mutator transaction binding the contract method 0x9e58eb9f.
//
// Solidity: function setUpOrg(string _orgId, uint256 _breadth, uint256 _depth) returns()
func (_OrgManagerAbi *OrgManagerAbiTransactorSession) SetUpOrg(_orgId string, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.SetUpOrg(&_OrgManagerAbi.TransactOpts, _orgId, _breadth, _depth)
}

// UpdateOrg is a paid mutator transaction binding the contract method 0x0cc27493.
//
// Solidity: function updateOrg(string _orgId, uint256 _action) returns(uint256)
func (_OrgManagerAbi *OrgManagerAbiTransactor) UpdateOrg(opts *bind.TransactOpts, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.contract.Transact(opts, "updateOrg", _orgId, _action)
}

// UpdateOrg is a paid mutator transaction binding the contract method 0x0cc27493.
//
// Solidity: function updateOrg(string _orgId, uint256 _action) returns(uint256)
func (_OrgManagerAbi *OrgManagerAbiSession) UpdateOrg(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.UpdateOrg(&_OrgManagerAbi.TransactOpts, _orgId, _action)
}

// UpdateOrg is a paid mutator transaction binding the contract method 0x0cc27493.
//
// Solidity: function updateOrg(string _orgId, uint256 _action) returns(uint256)
func (_OrgManagerAbi *OrgManagerAbiTransactorSession) UpdateOrg(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManagerAbi.Contract.UpdateOrg(&_OrgManagerAbi.TransactOpts, _orgId, _action)
}

// OrgManagerAbiOrgApprovedIterator is returned from FilterOrgApproved and is used to iterate over the raw logs and unpacked data for OrgApproved events raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgApprovedIterator struct {
	Event *OrgManagerAbiOrgApproved // Event containing the contract specifics and raw log

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
func (it *OrgManagerAbiOrgApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerAbiOrgApproved)
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
		it.Event = new(OrgManagerAbiOrgApproved)
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
func (it *OrgManagerAbiOrgApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerAbiOrgApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerAbiOrgApproved represents a OrgApproved event raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgApproved struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgApproved is a free log retrieval operation binding the contract event 0xd705723a50859c9cc1d3953e10b8b9478820e7a62927ad3215897ed87b20591c.
//
// Solidity: event OrgApproved(string _orgId, string _porgId, string _ultParent, uint256 _level, uint256 _status)
func (_OrgManagerAbi *OrgManagerAbiFilterer) FilterOrgApproved(opts *bind.FilterOpts) (*OrgManagerAbiOrgApprovedIterator, error) {

	logs, sub, err := _OrgManagerAbi.contract.FilterLogs(opts, "OrgApproved")
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiOrgApprovedIterator{contract: _OrgManagerAbi.contract, event: "OrgApproved", logs: logs, sub: sub}, nil
}

// WatchOrgApproved is a free log subscription operation binding the contract event 0xd705723a50859c9cc1d3953e10b8b9478820e7a62927ad3215897ed87b20591c.
//
// Solidity: event OrgApproved(string _orgId, string _porgId, string _ultParent, uint256 _level, uint256 _status)
func (_OrgManagerAbi *OrgManagerAbiFilterer) WatchOrgApproved(opts *bind.WatchOpts, sink chan<- *OrgManagerAbiOrgApproved) (event.Subscription, error) {

	logs, sub, err := _OrgManagerAbi.contract.WatchLogs(opts, "OrgApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerAbiOrgApproved)
				if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgApproved", log); err != nil {
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

// ParseOrgApproved is a log parse operation binding the contract event 0xd705723a50859c9cc1d3953e10b8b9478820e7a62927ad3215897ed87b20591c.
//
// Solidity: event OrgApproved(string _orgId, string _porgId, string _ultParent, uint256 _level, uint256 _status)
func (_OrgManagerAbi *OrgManagerAbiFilterer) ParseOrgApproved(log types.Log) (*OrgManagerAbiOrgApproved, error) {
	event := new(OrgManagerAbiOrgApproved)
	if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrgManagerAbiOrgPendingApprovalIterator is returned from FilterOrgPendingApproval and is used to iterate over the raw logs and unpacked data for OrgPendingApproval events raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgPendingApprovalIterator struct {
	Event *OrgManagerAbiOrgPendingApproval // Event containing the contract specifics and raw log

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
func (it *OrgManagerAbiOrgPendingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerAbiOrgPendingApproval)
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
		it.Event = new(OrgManagerAbiOrgPendingApproval)
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
func (it *OrgManagerAbiOrgPendingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerAbiOrgPendingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerAbiOrgPendingApproval represents a OrgPendingApproval event raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgPendingApproval struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgPendingApproval is a free log retrieval operation binding the contract event 0x0e8b7be64e0c730234ba2cd252b227fb481d7a247ba806d1941144c535bf054b.
//
// Solidity: event OrgPendingApproval(string _orgId, string _porgId, string _ultParent, uint256 _level, uint256 _status)
func (_OrgManagerAbi *OrgManagerAbiFilterer) FilterOrgPendingApproval(opts *bind.FilterOpts) (*OrgManagerAbiOrgPendingApprovalIterator, error) {

	logs, sub, err := _OrgManagerAbi.contract.FilterLogs(opts, "OrgPendingApproval")
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiOrgPendingApprovalIterator{contract: _OrgManagerAbi.contract, event: "OrgPendingApproval", logs: logs, sub: sub}, nil
}

// WatchOrgPendingApproval is a free log subscription operation binding the contract event 0x0e8b7be64e0c730234ba2cd252b227fb481d7a247ba806d1941144c535bf054b.
//
// Solidity: event OrgPendingApproval(string _orgId, string _porgId, string _ultParent, uint256 _level, uint256 _status)
func (_OrgManagerAbi *OrgManagerAbiFilterer) WatchOrgPendingApproval(opts *bind.WatchOpts, sink chan<- *OrgManagerAbiOrgPendingApproval) (event.Subscription, error) {

	logs, sub, err := _OrgManagerAbi.contract.WatchLogs(opts, "OrgPendingApproval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerAbiOrgPendingApproval)
				if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgPendingApproval", log); err != nil {
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

// ParseOrgPendingApproval is a log parse operation binding the contract event 0x0e8b7be64e0c730234ba2cd252b227fb481d7a247ba806d1941144c535bf054b.
//
// Solidity: event OrgPendingApproval(string _orgId, string _porgId, string _ultParent, uint256 _level, uint256 _status)
func (_OrgManagerAbi *OrgManagerAbiFilterer) ParseOrgPendingApproval(log types.Log) (*OrgManagerAbiOrgPendingApproval, error) {
	event := new(OrgManagerAbiOrgPendingApproval)
	if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgPendingApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrgManagerAbiOrgSuspendedIterator is returned from FilterOrgSuspended and is used to iterate over the raw logs and unpacked data for OrgSuspended events raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgSuspendedIterator struct {
	Event *OrgManagerAbiOrgSuspended // Event containing the contract specifics and raw log

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
func (it *OrgManagerAbiOrgSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerAbiOrgSuspended)
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
		it.Event = new(OrgManagerAbiOrgSuspended)
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
func (it *OrgManagerAbiOrgSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerAbiOrgSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerAbiOrgSuspended represents a OrgSuspended event raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgSuspended struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgSuspended is a free log retrieval operation binding the contract event 0x73ccf8d6c8385bf5347269bd59712da33183c1a5e1702494bcdb87d0f4674d96.
//
// Solidity: event OrgSuspended(string _orgId, string _porgId, string _ultParent, uint256 _level)
func (_OrgManagerAbi *OrgManagerAbiFilterer) FilterOrgSuspended(opts *bind.FilterOpts) (*OrgManagerAbiOrgSuspendedIterator, error) {

	logs, sub, err := _OrgManagerAbi.contract.FilterLogs(opts, "OrgSuspended")
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiOrgSuspendedIterator{contract: _OrgManagerAbi.contract, event: "OrgSuspended", logs: logs, sub: sub}, nil
}

// WatchOrgSuspended is a free log subscription operation binding the contract event 0x73ccf8d6c8385bf5347269bd59712da33183c1a5e1702494bcdb87d0f4674d96.
//
// Solidity: event OrgSuspended(string _orgId, string _porgId, string _ultParent, uint256 _level)
func (_OrgManagerAbi *OrgManagerAbiFilterer) WatchOrgSuspended(opts *bind.WatchOpts, sink chan<- *OrgManagerAbiOrgSuspended) (event.Subscription, error) {

	logs, sub, err := _OrgManagerAbi.contract.WatchLogs(opts, "OrgSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerAbiOrgSuspended)
				if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgSuspended", log); err != nil {
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

// ParseOrgSuspended is a log parse operation binding the contract event 0x73ccf8d6c8385bf5347269bd59712da33183c1a5e1702494bcdb87d0f4674d96.
//
// Solidity: event OrgSuspended(string _orgId, string _porgId, string _ultParent, uint256 _level)
func (_OrgManagerAbi *OrgManagerAbiFilterer) ParseOrgSuspended(log types.Log) (*OrgManagerAbiOrgSuspended, error) {
	event := new(OrgManagerAbiOrgSuspended)
	if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrgManagerAbiOrgSuspensionRevokedIterator is returned from FilterOrgSuspensionRevoked and is used to iterate over the raw logs and unpacked data for OrgSuspensionRevoked events raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgSuspensionRevokedIterator struct {
	Event *OrgManagerAbiOrgSuspensionRevoked // Event containing the contract specifics and raw log

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
func (it *OrgManagerAbiOrgSuspensionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerAbiOrgSuspensionRevoked)
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
		it.Event = new(OrgManagerAbiOrgSuspensionRevoked)
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
func (it *OrgManagerAbiOrgSuspensionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerAbiOrgSuspensionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerAbiOrgSuspensionRevoked represents a OrgSuspensionRevoked event raised by the OrgManagerAbi contract.
type OrgManagerAbiOrgSuspensionRevoked struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgSuspensionRevoked is a free log retrieval operation binding the contract event 0x882f030c609566cd82918a97d457fd48f9cfcefd11282e2654cde3f94579c15f.
//
// Solidity: event OrgSuspensionRevoked(string _orgId, string _porgId, string _ultParent, uint256 _level)
func (_OrgManagerAbi *OrgManagerAbiFilterer) FilterOrgSuspensionRevoked(opts *bind.FilterOpts) (*OrgManagerAbiOrgSuspensionRevokedIterator, error) {

	logs, sub, err := _OrgManagerAbi.contract.FilterLogs(opts, "OrgSuspensionRevoked")
	if err != nil {
		return nil, err
	}
	return &OrgManagerAbiOrgSuspensionRevokedIterator{contract: _OrgManagerAbi.contract, event: "OrgSuspensionRevoked", logs: logs, sub: sub}, nil
}

// WatchOrgSuspensionRevoked is a free log subscription operation binding the contract event 0x882f030c609566cd82918a97d457fd48f9cfcefd11282e2654cde3f94579c15f.
//
// Solidity: event OrgSuspensionRevoked(string _orgId, string _porgId, string _ultParent, uint256 _level)
func (_OrgManagerAbi *OrgManagerAbiFilterer) WatchOrgSuspensionRevoked(opts *bind.WatchOpts, sink chan<- *OrgManagerAbiOrgSuspensionRevoked) (event.Subscription, error) {

	logs, sub, err := _OrgManagerAbi.contract.WatchLogs(opts, "OrgSuspensionRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerAbiOrgSuspensionRevoked)
				if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgSuspensionRevoked", log); err != nil {
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

// ParseOrgSuspensionRevoked is a log parse operation binding the contract event 0x882f030c609566cd82918a97d457fd48f9cfcefd11282e2654cde3f94579c15f.
//
// Solidity: event OrgSuspensionRevoked(string _orgId, string _porgId, string _ultParent, uint256 _level)
func (_OrgManagerAbi *OrgManagerAbiFilterer) ParseOrgSuspensionRevoked(log types.Log) (*OrgManagerAbiOrgSuspensionRevoked, error) {
	event := new(OrgManagerAbiOrgSuspensionRevoked)
	if err := _OrgManagerAbi.contract.UnpackLog(event, "OrgSuspensionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
