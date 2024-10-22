// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account_manager_abi

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

// AccountManagerAbiMetaData contains all meta data concerning the AccountManagerAbi contract.
var AccountManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_orgAdmin\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"AccountAccessModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_orgAdmin\",\"type\":\"bool\"}],\"name\":\"AccountAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"AccountStatusChanged\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addNewAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"voterUpdate\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_adminRole\",\"type\":\"bool\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"assignAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"checkOrgAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_aIndex\",\"type\":\"uint256\"}],\"name\":\"getAccountDetailsFromIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountOrgRole\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountRole\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"orgAdminExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeExistingAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"voterUpdate\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setDefaults\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateAccountStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccountManagerAbi is an auto generated Go binding around an Ethereum contract.
type AccountManagerAbi struct {
	AccountManagerAbiCaller     // Read-only binding to the contract
	AccountManagerAbiTransactor // Write-only binding to the contract
	AccountManagerAbiFilterer   // Log filterer for contract events
}

// AccountManagerAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerAbiSession struct {
	Contract     *AccountManagerAbi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountManagerAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerAbiCallerSession struct {
	Contract *AccountManagerAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AccountManagerAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerAbiTransactorSession struct {
	Contract     *AccountManagerAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AccountManagerAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerAbiRaw struct {
	Contract *AccountManagerAbi // Generic contract binding to access the raw methods on
}

// AccountManagerAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerAbiCallerRaw struct {
	Contract *AccountManagerAbiCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerAbiTransactorRaw struct {
	Contract *AccountManagerAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManagerAbi creates a new instance of AccountManagerAbi, bound to a specific deployed contract.
func NewAccountManagerAbi(address common.Address, backend bind.ContractBackend) (*AccountManagerAbi, error) {
	contract, err := bindAccountManagerAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbi{AccountManagerAbiCaller: AccountManagerAbiCaller{contract: contract}, AccountManagerAbiTransactor: AccountManagerAbiTransactor{contract: contract}, AccountManagerAbiFilterer: AccountManagerAbiFilterer{contract: contract}}, nil
}

// NewAccountManagerAbiCaller creates a new read-only instance of AccountManagerAbi, bound to a specific deployed contract.
func NewAccountManagerAbiCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerAbiCaller, error) {
	contract, err := bindAccountManagerAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbiCaller{contract: contract}, nil
}

// NewAccountManagerAbiTransactor creates a new write-only instance of AccountManagerAbi, bound to a specific deployed contract.
func NewAccountManagerAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerAbiTransactor, error) {
	contract, err := bindAccountManagerAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbiTransactor{contract: contract}, nil
}

// NewAccountManagerAbiFilterer creates a new log filterer instance of AccountManagerAbi, bound to a specific deployed contract.
func NewAccountManagerAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerAbiFilterer, error) {
	contract, err := bindAccountManagerAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbiFilterer{contract: contract}, nil
}

// bindAccountManagerAbi binds a generic wrapper to an already deployed contract.
func bindAccountManagerAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerAbi *AccountManagerAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerAbi.Contract.AccountManagerAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerAbi *AccountManagerAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AccountManagerAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerAbi *AccountManagerAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AccountManagerAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerAbi *AccountManagerAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerAbi *AccountManagerAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerAbi *AccountManagerAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.contract.Transact(opts, method, params...)
}

// CheckOrgAdmin is a free data retrieval call binding the contract method 0xe8b42bf4.
//
// Solidity: function checkOrgAdmin(address _account, string _orgId, string _ultParent) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiCaller) CheckOrgAdmin(opts *bind.CallOpts, _account common.Address, _orgId string, _ultParent string) (bool, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "checkOrgAdmin", _account, _orgId, _ultParent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOrgAdmin is a free data retrieval call binding the contract method 0xe8b42bf4.
//
// Solidity: function checkOrgAdmin(address _account, string _orgId, string _ultParent) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiSession) CheckOrgAdmin(_account common.Address, _orgId string, _ultParent string) (bool, error) {
	return _AccountManagerAbi.Contract.CheckOrgAdmin(&_AccountManagerAbi.CallOpts, _account, _orgId, _ultParent)
}

// CheckOrgAdmin is a free data retrieval call binding the contract method 0xe8b42bf4.
//
// Solidity: function checkOrgAdmin(address _account, string _orgId, string _ultParent) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) CheckOrgAdmin(_account common.Address, _orgId string, _ultParent string) (bool, error) {
	return _AccountManagerAbi.Contract.CheckOrgAdmin(&_AccountManagerAbi.CallOpts, _account, _orgId, _ultParent)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address _account) view returns(address, string, string, uint256, bool)
func (_AccountManagerAbi *AccountManagerAbiCaller) GetAccountDetails(opts *bind.CallOpts, _account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "getAccountDetails", _account)

	if err != nil {
		return *new(common.Address), *new(string), *new(string), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address _account) view returns(address, string, string, uint256, bool)
func (_AccountManagerAbi *AccountManagerAbiSession) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return _AccountManagerAbi.Contract.GetAccountDetails(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address _account) view returns(address, string, string, uint256, bool)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return _AccountManagerAbi.Contract.GetAccountDetails(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountDetailsFromIndex is a free data retrieval call binding the contract method 0xb2018568.
//
// Solidity: function getAccountDetailsFromIndex(uint256 _aIndex) view returns(address, string, string, uint256, bool)
func (_AccountManagerAbi *AccountManagerAbiCaller) GetAccountDetailsFromIndex(opts *bind.CallOpts, _aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "getAccountDetailsFromIndex", _aIndex)

	if err != nil {
		return *new(common.Address), *new(string), *new(string), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// GetAccountDetailsFromIndex is a free data retrieval call binding the contract method 0xb2018568.
//
// Solidity: function getAccountDetailsFromIndex(uint256 _aIndex) view returns(address, string, string, uint256, bool)
func (_AccountManagerAbi *AccountManagerAbiSession) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return _AccountManagerAbi.Contract.GetAccountDetailsFromIndex(&_AccountManagerAbi.CallOpts, _aIndex)
}

// GetAccountDetailsFromIndex is a free data retrieval call binding the contract method 0xb2018568.
//
// Solidity: function getAccountDetailsFromIndex(uint256 _aIndex) view returns(address, string, string, uint256, bool)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return _AccountManagerAbi.Contract.GetAccountDetailsFromIndex(&_AccountManagerAbi.CallOpts, _aIndex)
}

// GetAccountOrgRole is a free data retrieval call binding the contract method 0x6acee5fd.
//
// Solidity: function getAccountOrgRole(address _account) view returns(string, string)
func (_AccountManagerAbi *AccountManagerAbiCaller) GetAccountOrgRole(opts *bind.CallOpts, _account common.Address) (string, string, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "getAccountOrgRole", _account)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetAccountOrgRole is a free data retrieval call binding the contract method 0x6acee5fd.
//
// Solidity: function getAccountOrgRole(address _account) view returns(string, string)
func (_AccountManagerAbi *AccountManagerAbiSession) GetAccountOrgRole(_account common.Address) (string, string, error) {
	return _AccountManagerAbi.Contract.GetAccountOrgRole(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountOrgRole is a free data retrieval call binding the contract method 0x6acee5fd.
//
// Solidity: function getAccountOrgRole(address _account) view returns(string, string)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) GetAccountOrgRole(_account common.Address) (string, string, error) {
	return _AccountManagerAbi.Contract.GetAccountOrgRole(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountRole is a free data retrieval call binding the contract method 0x81d66b23.
//
// Solidity: function getAccountRole(address _account) view returns(string)
func (_AccountManagerAbi *AccountManagerAbiCaller) GetAccountRole(opts *bind.CallOpts, _account common.Address) (string, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "getAccountRole", _account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAccountRole is a free data retrieval call binding the contract method 0x81d66b23.
//
// Solidity: function getAccountRole(address _account) view returns(string)
func (_AccountManagerAbi *AccountManagerAbiSession) GetAccountRole(_account common.Address) (string, error) {
	return _AccountManagerAbi.Contract.GetAccountRole(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountRole is a free data retrieval call binding the contract method 0x81d66b23.
//
// Solidity: function getAccountRole(address _account) view returns(string)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) GetAccountRole(_account common.Address) (string, error) {
	return _AccountManagerAbi.Contract.GetAccountRole(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xfd4fa05a.
//
// Solidity: function getAccountStatus(address _account) view returns(uint256)
func (_AccountManagerAbi *AccountManagerAbiCaller) GetAccountStatus(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "getAccountStatus", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountStatus is a free data retrieval call binding the contract method 0xfd4fa05a.
//
// Solidity: function getAccountStatus(address _account) view returns(uint256)
func (_AccountManagerAbi *AccountManagerAbiSession) GetAccountStatus(_account common.Address) (*big.Int, error) {
	return _AccountManagerAbi.Contract.GetAccountStatus(&_AccountManagerAbi.CallOpts, _account)
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xfd4fa05a.
//
// Solidity: function getAccountStatus(address _account) view returns(uint256)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) GetAccountStatus(_account common.Address) (*big.Int, error) {
	return _AccountManagerAbi.Contract.GetAccountStatus(&_AccountManagerAbi.CallOpts, _account)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() view returns(uint256)
func (_AccountManagerAbi *AccountManagerAbiCaller) GetNumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "getNumberOfAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() view returns(uint256)
func (_AccountManagerAbi *AccountManagerAbiSession) GetNumberOfAccounts() (*big.Int, error) {
	return _AccountManagerAbi.Contract.GetNumberOfAccounts(&_AccountManagerAbi.CallOpts)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() view returns(uint256)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) GetNumberOfAccounts() (*big.Int, error) {
	return _AccountManagerAbi.Contract.GetNumberOfAccounts(&_AccountManagerAbi.CallOpts)
}

// OrgAdminExists is a free data retrieval call binding the contract method 0x950145cf.
//
// Solidity: function orgAdminExists(string _orgId) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiCaller) OrgAdminExists(opts *bind.CallOpts, _orgId string) (bool, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "orgAdminExists", _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrgAdminExists is a free data retrieval call binding the contract method 0x950145cf.
//
// Solidity: function orgAdminExists(string _orgId) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiSession) OrgAdminExists(_orgId string) (bool, error) {
	return _AccountManagerAbi.Contract.OrgAdminExists(&_AccountManagerAbi.CallOpts, _orgId)
}

// OrgAdminExists is a free data retrieval call binding the contract method 0x950145cf.
//
// Solidity: function orgAdminExists(string _orgId) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) OrgAdminExists(_orgId string) (bool, error) {
	return _AccountManagerAbi.Contract.OrgAdminExists(&_AccountManagerAbi.CallOpts, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiCaller) ValidateAccount(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var out []interface{}
	err := _AccountManagerAbi.contract.Call(opts, &out, "validateAccount", _account, _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _AccountManagerAbi.Contract.ValidateAccount(&_AccountManagerAbi.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_AccountManagerAbi *AccountManagerAbiCallerSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _AccountManagerAbi.Contract.ValidateAccount(&_AccountManagerAbi.CallOpts, _account, _orgId)
}

// AddNewAdmin is a paid mutator transaction binding the contract method 0xc214e5e5.
//
// Solidity: function addNewAdmin(string _orgId, address _account) returns(bool voterUpdate)
func (_AccountManagerAbi *AccountManagerAbiTransactor) AddNewAdmin(opts *bind.TransactOpts, _orgId string, _account common.Address) (*types.Transaction, error) {
	return _AccountManagerAbi.contract.Transact(opts, "addNewAdmin", _orgId, _account)
}

// AddNewAdmin is a paid mutator transaction binding the contract method 0xc214e5e5.
//
// Solidity: function addNewAdmin(string _orgId, address _account) returns(bool voterUpdate)
func (_AccountManagerAbi *AccountManagerAbiSession) AddNewAdmin(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AddNewAdmin(&_AccountManagerAbi.TransactOpts, _orgId, _account)
}

// AddNewAdmin is a paid mutator transaction binding the contract method 0xc214e5e5.
//
// Solidity: function addNewAdmin(string _orgId, address _account) returns(bool voterUpdate)
func (_AccountManagerAbi *AccountManagerAbiTransactorSession) AddNewAdmin(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AddNewAdmin(&_AccountManagerAbi.TransactOpts, _orgId, _account)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x143a5604.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, bool _adminRole) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactor) AssignAccountRole(opts *bind.TransactOpts, _account common.Address, _orgId string, _roleId string, _adminRole bool) (*types.Transaction, error) {
	return _AccountManagerAbi.contract.Transact(opts, "assignAccountRole", _account, _orgId, _roleId, _adminRole)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x143a5604.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, bool _adminRole) returns()
func (_AccountManagerAbi *AccountManagerAbiSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string, _adminRole bool) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AssignAccountRole(&_AccountManagerAbi.TransactOpts, _account, _orgId, _roleId, _adminRole)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x143a5604.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, bool _adminRole) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactorSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string, _adminRole bool) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AssignAccountRole(&_AccountManagerAbi.TransactOpts, _account, _orgId, _roleId, _adminRole)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0xe3483a9d.
//
// Solidity: function assignAdminRole(address _account, string _orgId, string _roleId, uint256 _status) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactor) AssignAdminRole(opts *bind.TransactOpts, _account common.Address, _orgId string, _roleId string, _status *big.Int) (*types.Transaction, error) {
	return _AccountManagerAbi.contract.Transact(opts, "assignAdminRole", _account, _orgId, _roleId, _status)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0xe3483a9d.
//
// Solidity: function assignAdminRole(address _account, string _orgId, string _roleId, uint256 _status) returns()
func (_AccountManagerAbi *AccountManagerAbiSession) AssignAdminRole(_account common.Address, _orgId string, _roleId string, _status *big.Int) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AssignAdminRole(&_AccountManagerAbi.TransactOpts, _account, _orgId, _roleId, _status)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0xe3483a9d.
//
// Solidity: function assignAdminRole(address _account, string _orgId, string _roleId, uint256 _status) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactorSession) AssignAdminRole(_account common.Address, _orgId string, _roleId string, _status *big.Int) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.AssignAdminRole(&_AccountManagerAbi.TransactOpts, _account, _orgId, _roleId, _status)
}

// RemoveExistingAdmin is a paid mutator transaction binding the contract method 0x1d09dc93.
//
// Solidity: function removeExistingAdmin(string _orgId) returns(bool voterUpdate, address account)
func (_AccountManagerAbi *AccountManagerAbiTransactor) RemoveExistingAdmin(opts *bind.TransactOpts, _orgId string) (*types.Transaction, error) {
	return _AccountManagerAbi.contract.Transact(opts, "removeExistingAdmin", _orgId)
}

// RemoveExistingAdmin is a paid mutator transaction binding the contract method 0x1d09dc93.
//
// Solidity: function removeExistingAdmin(string _orgId) returns(bool voterUpdate, address account)
func (_AccountManagerAbi *AccountManagerAbiSession) RemoveExistingAdmin(_orgId string) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.RemoveExistingAdmin(&_AccountManagerAbi.TransactOpts, _orgId)
}

// RemoveExistingAdmin is a paid mutator transaction binding the contract method 0x1d09dc93.
//
// Solidity: function removeExistingAdmin(string _orgId) returns(bool voterUpdate, address account)
func (_AccountManagerAbi *AccountManagerAbiTransactorSession) RemoveExistingAdmin(_orgId string) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.RemoveExistingAdmin(&_AccountManagerAbi.TransactOpts, _orgId)
}

// SetDefaults is a paid mutator transaction binding the contract method 0xcef7f6af.
//
// Solidity: function setDefaults(string _nwAdminRole, string _oAdminRole) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactor) SetDefaults(opts *bind.TransactOpts, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _AccountManagerAbi.contract.Transact(opts, "setDefaults", _nwAdminRole, _oAdminRole)
}

// SetDefaults is a paid mutator transaction binding the contract method 0xcef7f6af.
//
// Solidity: function setDefaults(string _nwAdminRole, string _oAdminRole) returns()
func (_AccountManagerAbi *AccountManagerAbiSession) SetDefaults(_nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.SetDefaults(&_AccountManagerAbi.TransactOpts, _nwAdminRole, _oAdminRole)
}

// SetDefaults is a paid mutator transaction binding the contract method 0xcef7f6af.
//
// Solidity: function setDefaults(string _nwAdminRole, string _oAdminRole) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactorSession) SetDefaults(_nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.SetDefaults(&_AccountManagerAbi.TransactOpts, _nwAdminRole, _oAdminRole)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactor) UpdateAccountStatus(opts *bind.TransactOpts, _orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return _AccountManagerAbi.contract.Transact(opts, "updateAccountStatus", _orgId, _account, _action)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action) returns()
func (_AccountManagerAbi *AccountManagerAbiSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.UpdateAccountStatus(&_AccountManagerAbi.TransactOpts, _orgId, _account, _action)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action) returns()
func (_AccountManagerAbi *AccountManagerAbiTransactorSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return _AccountManagerAbi.Contract.UpdateAccountStatus(&_AccountManagerAbi.TransactOpts, _orgId, _account, _action)
}

// AccountManagerAbiAccountAccessModifiedIterator is returned from FilterAccountAccessModified and is used to iterate over the raw logs and unpacked data for AccountAccessModified events raised by the AccountManagerAbi contract.
type AccountManagerAbiAccountAccessModifiedIterator struct {
	Event *AccountManagerAbiAccountAccessModified // Event containing the contract specifics and raw log

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
func (it *AccountManagerAbiAccountAccessModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAbiAccountAccessModified)
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
		it.Event = new(AccountManagerAbiAccountAccessModified)
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
func (it *AccountManagerAbiAccountAccessModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAbiAccountAccessModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAbiAccountAccessModified represents a AccountAccessModified event raised by the AccountManagerAbi contract.
type AccountManagerAbiAccountAccessModified struct {
	Account  common.Address
	OrgId    string
	RoleId   string
	OrgAdmin bool
	Status   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountAccessModified is a free log retrieval operation binding the contract event 0x68e62a03aeb0a125c2fc869eed72f2fca473680987bdd680c093a534e17cc776.
//
// Solidity: event AccountAccessModified(address _account, string _orgId, string _roleId, bool _orgAdmin, uint256 _status)
func (_AccountManagerAbi *AccountManagerAbiFilterer) FilterAccountAccessModified(opts *bind.FilterOpts) (*AccountManagerAbiAccountAccessModifiedIterator, error) {

	logs, sub, err := _AccountManagerAbi.contract.FilterLogs(opts, "AccountAccessModified")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbiAccountAccessModifiedIterator{contract: _AccountManagerAbi.contract, event: "AccountAccessModified", logs: logs, sub: sub}, nil
}

// WatchAccountAccessModified is a free log subscription operation binding the contract event 0x68e62a03aeb0a125c2fc869eed72f2fca473680987bdd680c093a534e17cc776.
//
// Solidity: event AccountAccessModified(address _account, string _orgId, string _roleId, bool _orgAdmin, uint256 _status)
func (_AccountManagerAbi *AccountManagerAbiFilterer) WatchAccountAccessModified(opts *bind.WatchOpts, sink chan<- *AccountManagerAbiAccountAccessModified) (event.Subscription, error) {

	logs, sub, err := _AccountManagerAbi.contract.WatchLogs(opts, "AccountAccessModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAbiAccountAccessModified)
				if err := _AccountManagerAbi.contract.UnpackLog(event, "AccountAccessModified", log); err != nil {
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

// ParseAccountAccessModified is a log parse operation binding the contract event 0x68e62a03aeb0a125c2fc869eed72f2fca473680987bdd680c093a534e17cc776.
//
// Solidity: event AccountAccessModified(address _account, string _orgId, string _roleId, bool _orgAdmin, uint256 _status)
func (_AccountManagerAbi *AccountManagerAbiFilterer) ParseAccountAccessModified(log types.Log) (*AccountManagerAbiAccountAccessModified, error) {
	event := new(AccountManagerAbiAccountAccessModified)
	if err := _AccountManagerAbi.contract.UnpackLog(event, "AccountAccessModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerAbiAccountAccessRevokedIterator is returned from FilterAccountAccessRevoked and is used to iterate over the raw logs and unpacked data for AccountAccessRevoked events raised by the AccountManagerAbi contract.
type AccountManagerAbiAccountAccessRevokedIterator struct {
	Event *AccountManagerAbiAccountAccessRevoked // Event containing the contract specifics and raw log

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
func (it *AccountManagerAbiAccountAccessRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAbiAccountAccessRevoked)
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
		it.Event = new(AccountManagerAbiAccountAccessRevoked)
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
func (it *AccountManagerAbiAccountAccessRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAbiAccountAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAbiAccountAccessRevoked represents a AccountAccessRevoked event raised by the AccountManagerAbi contract.
type AccountManagerAbiAccountAccessRevoked struct {
	Account  common.Address
	OrgId    string
	RoleId   string
	OrgAdmin bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountAccessRevoked is a free log retrieval operation binding the contract event 0x6b5105396435a8a139aeed682dd573cd2a7e6279de77f8c11f95a30399212ad1.
//
// Solidity: event AccountAccessRevoked(address _account, string _orgId, string _roleId, bool _orgAdmin)
func (_AccountManagerAbi *AccountManagerAbiFilterer) FilterAccountAccessRevoked(opts *bind.FilterOpts) (*AccountManagerAbiAccountAccessRevokedIterator, error) {

	logs, sub, err := _AccountManagerAbi.contract.FilterLogs(opts, "AccountAccessRevoked")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbiAccountAccessRevokedIterator{contract: _AccountManagerAbi.contract, event: "AccountAccessRevoked", logs: logs, sub: sub}, nil
}

// WatchAccountAccessRevoked is a free log subscription operation binding the contract event 0x6b5105396435a8a139aeed682dd573cd2a7e6279de77f8c11f95a30399212ad1.
//
// Solidity: event AccountAccessRevoked(address _account, string _orgId, string _roleId, bool _orgAdmin)
func (_AccountManagerAbi *AccountManagerAbiFilterer) WatchAccountAccessRevoked(opts *bind.WatchOpts, sink chan<- *AccountManagerAbiAccountAccessRevoked) (event.Subscription, error) {

	logs, sub, err := _AccountManagerAbi.contract.WatchLogs(opts, "AccountAccessRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAbiAccountAccessRevoked)
				if err := _AccountManagerAbi.contract.UnpackLog(event, "AccountAccessRevoked", log); err != nil {
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

// ParseAccountAccessRevoked is a log parse operation binding the contract event 0x6b5105396435a8a139aeed682dd573cd2a7e6279de77f8c11f95a30399212ad1.
//
// Solidity: event AccountAccessRevoked(address _account, string _orgId, string _roleId, bool _orgAdmin)
func (_AccountManagerAbi *AccountManagerAbiFilterer) ParseAccountAccessRevoked(log types.Log) (*AccountManagerAbiAccountAccessRevoked, error) {
	event := new(AccountManagerAbiAccountAccessRevoked)
	if err := _AccountManagerAbi.contract.UnpackLog(event, "AccountAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerAbiAccountStatusChangedIterator is returned from FilterAccountStatusChanged and is used to iterate over the raw logs and unpacked data for AccountStatusChanged events raised by the AccountManagerAbi contract.
type AccountManagerAbiAccountStatusChangedIterator struct {
	Event *AccountManagerAbiAccountStatusChanged // Event containing the contract specifics and raw log

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
func (it *AccountManagerAbiAccountStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAbiAccountStatusChanged)
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
		it.Event = new(AccountManagerAbiAccountStatusChanged)
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
func (it *AccountManagerAbiAccountStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAbiAccountStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAbiAccountStatusChanged represents a AccountStatusChanged event raised by the AccountManagerAbi contract.
type AccountManagerAbiAccountStatusChanged struct {
	Account common.Address
	OrgId   string
	Status  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountStatusChanged is a free log retrieval operation binding the contract event 0x36b0ea38154dec5e98b6bf928b971a9db5e8cd4b6946350e9e43fb9848c70b25.
//
// Solidity: event AccountStatusChanged(address _account, string _orgId, uint256 _status)
func (_AccountManagerAbi *AccountManagerAbiFilterer) FilterAccountStatusChanged(opts *bind.FilterOpts) (*AccountManagerAbiAccountStatusChangedIterator, error) {

	logs, sub, err := _AccountManagerAbi.contract.FilterLogs(opts, "AccountStatusChanged")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAbiAccountStatusChangedIterator{contract: _AccountManagerAbi.contract, event: "AccountStatusChanged", logs: logs, sub: sub}, nil
}

// WatchAccountStatusChanged is a free log subscription operation binding the contract event 0x36b0ea38154dec5e98b6bf928b971a9db5e8cd4b6946350e9e43fb9848c70b25.
//
// Solidity: event AccountStatusChanged(address _account, string _orgId, uint256 _status)
func (_AccountManagerAbi *AccountManagerAbiFilterer) WatchAccountStatusChanged(opts *bind.WatchOpts, sink chan<- *AccountManagerAbiAccountStatusChanged) (event.Subscription, error) {

	logs, sub, err := _AccountManagerAbi.contract.WatchLogs(opts, "AccountStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAbiAccountStatusChanged)
				if err := _AccountManagerAbi.contract.UnpackLog(event, "AccountStatusChanged", log); err != nil {
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

// ParseAccountStatusChanged is a log parse operation binding the contract event 0x36b0ea38154dec5e98b6bf928b971a9db5e8cd4b6946350e9e43fb9848c70b25.
//
// Solidity: event AccountStatusChanged(address _account, string _orgId, uint256 _status)
func (_AccountManagerAbi *AccountManagerAbiFilterer) ParseAccountStatusChanged(log types.Log) (*AccountManagerAbiAccountStatusChanged, error) {
	event := new(AccountManagerAbiAccountStatusChanged)
	if err := _AccountManagerAbi.contract.UnpackLog(event, "AccountStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
