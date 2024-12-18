// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdt_abi

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

// UsdtAbiMetaData contains all meta data concerning the UsdtAbi contract.
var UsdtABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_upgradedAddress\",\"type\":\"address\"}],\"name\":\"deprecate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deprecated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradedAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maximumFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBasisPoints\",\"type\":\"uint256\"},{\"name\":\"newMaxFee\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"basisPointsRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feeBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"Params\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// UsdtAbi is an auto generated Go binding around an Ethereum contract.
type UsdtAbi struct {
	UsdtAbiCaller     // Read-only binding to the contract
	UsdtAbiTransactor // Write-only binding to the contract
	UsdtAbiFilterer   // Log filterer for contract events
}

// UsdtAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdtAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdtAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdtAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdtAbiSession struct {
	Contract     *UsdtAbi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdtAbiCallerSession struct {
	Contract *UsdtAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UsdtAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdtAbiTransactorSession struct {
	Contract     *UsdtAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UsdtAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdtAbiRaw struct {
	Contract *UsdtAbi // Generic contract binding to access the raw methods on
}

// UsdtAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdtAbiCallerRaw struct {
	Contract *UsdtAbiCaller // Generic read-only contract binding to access the raw methods on
}

// UsdtAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdtAbiTransactorRaw struct {
	Contract *UsdtAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdtAbi creates a new instance of UsdtAbi, bound to a specific deployed contract.
func NewUsdtAbi(address common.Address, backend bind.ContractBackend) (*UsdtAbi, error) {
	contract, err := bindUsdtAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsdtAbi{UsdtAbiCaller: UsdtAbiCaller{contract: contract}, UsdtAbiTransactor: UsdtAbiTransactor{contract: contract}, UsdtAbiFilterer: UsdtAbiFilterer{contract: contract}}, nil
}

// NewUsdtAbiCaller creates a new read-only instance of UsdtAbi, bound to a specific deployed contract.
func NewUsdtAbiCaller(address common.Address, caller bind.ContractCaller) (*UsdtAbiCaller, error) {
	contract, err := bindUsdtAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtAbiCaller{contract: contract}, nil
}

// NewUsdtAbiTransactor creates a new write-only instance of UsdtAbi, bound to a specific deployed contract.
func NewUsdtAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdtAbiTransactor, error) {
	contract, err := bindUsdtAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtAbiTransactor{contract: contract}, nil
}

// NewUsdtAbiFilterer creates a new log filterer instance of UsdtAbi, bound to a specific deployed contract.
func NewUsdtAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdtAbiFilterer, error) {
	contract, err := bindUsdtAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdtAbiFilterer{contract: contract}, nil
}

// bindUsdtAbi binds a generic wrapper to an already deployed contract.
func bindUsdtAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsdtAbi *UsdtAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsdtAbi.Contract.UsdtAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsdtAbi *UsdtAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtAbi.Contract.UsdtAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsdtAbi *UsdtAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsdtAbi.Contract.UsdtAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsdtAbi *UsdtAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsdtAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsdtAbi *UsdtAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsdtAbi *UsdtAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsdtAbi.Contract.contract.Transact(opts, method, params...)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) MAXUINT() (*big.Int, error) {
	return _UsdtAbi.Contract.MAXUINT(&_UsdtAbi.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) MAXUINT() (*big.Int, error) {
	return _UsdtAbi.Contract.MAXUINT(&_UsdtAbi.CallOpts)
}

// TotalSupplyPrivate is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) TotalSupplyPrivate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "_totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyPrivate is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) TotalSupplyPrivate() (*big.Int, error) {
	return _UsdtAbi.Contract.TotalSupplyPrivate(&_UsdtAbi.CallOpts)
}

// TotalSupplyPrivate is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) TotalSupplyPrivate() (*big.Int, error) {
	return _UsdtAbi.Contract.TotalSupplyPrivate(&_UsdtAbi.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_UsdtAbi *UsdtAbiCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_UsdtAbi *UsdtAbiSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.Allowance(&_UsdtAbi.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_UsdtAbi *UsdtAbiCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.Allowance(&_UsdtAbi.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.Allowed(&_UsdtAbi.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.Allowed(&_UsdtAbi.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.BalanceOf(&_UsdtAbi.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.BalanceOf(&_UsdtAbi.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.Balances(&_UsdtAbi.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _UsdtAbi.Contract.Balances(&_UsdtAbi.CallOpts, arg0)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) BasisPointsRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "basisPointsRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) BasisPointsRate() (*big.Int, error) {
	return _UsdtAbi.Contract.BasisPointsRate(&_UsdtAbi.CallOpts)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) BasisPointsRate() (*big.Int, error) {
	return _UsdtAbi.Contract.BasisPointsRate(&_UsdtAbi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) Decimals() (*big.Int, error) {
	return _UsdtAbi.Contract.Decimals(&_UsdtAbi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) Decimals() (*big.Int, error) {
	return _UsdtAbi.Contract.Decimals(&_UsdtAbi.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_UsdtAbi *UsdtAbiCaller) Deprecated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "deprecated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_UsdtAbi *UsdtAbiSession) Deprecated() (bool, error) {
	return _UsdtAbi.Contract.Deprecated(&_UsdtAbi.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_UsdtAbi *UsdtAbiCallerSession) Deprecated() (bool, error) {
	return _UsdtAbi.Contract.Deprecated(&_UsdtAbi.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_UsdtAbi *UsdtAbiCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_UsdtAbi *UsdtAbiSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _UsdtAbi.Contract.GetBlackListStatus(&_UsdtAbi.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_UsdtAbi *UsdtAbiCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _UsdtAbi.Contract.GetBlackListStatus(&_UsdtAbi.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_UsdtAbi *UsdtAbiCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_UsdtAbi *UsdtAbiSession) GetOwner() (common.Address, error) {
	return _UsdtAbi.Contract.GetOwner(&_UsdtAbi.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_UsdtAbi *UsdtAbiCallerSession) GetOwner() (common.Address, error) {
	return _UsdtAbi.Contract.GetOwner(&_UsdtAbi.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_UsdtAbi *UsdtAbiCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_UsdtAbi *UsdtAbiSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _UsdtAbi.Contract.IsBlackListed(&_UsdtAbi.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_UsdtAbi *UsdtAbiCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _UsdtAbi.Contract.IsBlackListed(&_UsdtAbi.CallOpts, arg0)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) MaximumFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "maximumFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) MaximumFee() (*big.Int, error) {
	return _UsdtAbi.Contract.MaximumFee(&_UsdtAbi.CallOpts)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) MaximumFee() (*big.Int, error) {
	return _UsdtAbi.Contract.MaximumFee(&_UsdtAbi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UsdtAbi *UsdtAbiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UsdtAbi *UsdtAbiSession) Name() (string, error) {
	return _UsdtAbi.Contract.Name(&_UsdtAbi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UsdtAbi *UsdtAbiCallerSession) Name() (string, error) {
	return _UsdtAbi.Contract.Name(&_UsdtAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UsdtAbi *UsdtAbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UsdtAbi *UsdtAbiSession) Owner() (common.Address, error) {
	return _UsdtAbi.Contract.Owner(&_UsdtAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UsdtAbi *UsdtAbiCallerSession) Owner() (common.Address, error) {
	return _UsdtAbi.Contract.Owner(&_UsdtAbi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UsdtAbi *UsdtAbiCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UsdtAbi *UsdtAbiSession) Paused() (bool, error) {
	return _UsdtAbi.Contract.Paused(&_UsdtAbi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UsdtAbi *UsdtAbiCallerSession) Paused() (bool, error) {
	return _UsdtAbi.Contract.Paused(&_UsdtAbi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UsdtAbi *UsdtAbiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UsdtAbi *UsdtAbiSession) Symbol() (string, error) {
	return _UsdtAbi.Contract.Symbol(&_UsdtAbi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UsdtAbi *UsdtAbiCallerSession) Symbol() (string, error) {
	return _UsdtAbi.Contract.Symbol(&_UsdtAbi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UsdtAbi *UsdtAbiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UsdtAbi *UsdtAbiSession) TotalSupply() (*big.Int, error) {
	return _UsdtAbi.Contract.TotalSupply(&_UsdtAbi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UsdtAbi *UsdtAbiCallerSession) TotalSupply() (*big.Int, error) {
	return _UsdtAbi.Contract.TotalSupply(&_UsdtAbi.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_UsdtAbi *UsdtAbiCaller) UpgradedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdtAbi.contract.Call(opts, &out, "upgradedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_UsdtAbi *UsdtAbiSession) UpgradedAddress() (common.Address, error) {
	return _UsdtAbi.Contract.UpgradedAddress(&_UsdtAbi.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_UsdtAbi *UsdtAbiCallerSession) UpgradedAddress() (common.Address, error) {
	return _UsdtAbi.Contract.UpgradedAddress(&_UsdtAbi.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_UsdtAbi *UsdtAbiTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_UsdtAbi *UsdtAbiSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.AddBlackList(&_UsdtAbi.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.AddBlackList(&_UsdtAbi.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Approve(&_UsdtAbi.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Approve(&_UsdtAbi.TransactOpts, _spender, _value)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_UsdtAbi *UsdtAbiTransactor) Deprecate(opts *bind.TransactOpts, _upgradedAddress common.Address) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "deprecate", _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_UsdtAbi *UsdtAbiSession) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Deprecate(&_UsdtAbi.TransactOpts, _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Deprecate(&_UsdtAbi.TransactOpts, _upgradedAddress)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_UsdtAbi *UsdtAbiTransactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_UsdtAbi *UsdtAbiSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.DestroyBlackFunds(&_UsdtAbi.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.DestroyBlackFunds(&_UsdtAbi.TransactOpts, _blackListedUser)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_UsdtAbi *UsdtAbiTransactor) Issue(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "issue", amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_UsdtAbi *UsdtAbiSession) Issue(amount *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Issue(&_UsdtAbi.TransactOpts, amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Issue(amount *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Issue(&_UsdtAbi.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UsdtAbi *UsdtAbiTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UsdtAbi *UsdtAbiSession) Pause() (*types.Transaction, error) {
	return _UsdtAbi.Contract.Pause(&_UsdtAbi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Pause() (*types.Transaction, error) {
	return _UsdtAbi.Contract.Pause(&_UsdtAbi.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_UsdtAbi *UsdtAbiTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_UsdtAbi *UsdtAbiSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Redeem(&_UsdtAbi.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Redeem(&_UsdtAbi.TransactOpts, amount)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_UsdtAbi *UsdtAbiTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_UsdtAbi *UsdtAbiSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.RemoveBlackList(&_UsdtAbi.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.RemoveBlackList(&_UsdtAbi.TransactOpts, _clearedUser)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_UsdtAbi *UsdtAbiTransactor) SetParams(opts *bind.TransactOpts, newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "setParams", newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_UsdtAbi *UsdtAbiSession) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.SetParams(&_UsdtAbi.TransactOpts, newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.SetParams(&_UsdtAbi.TransactOpts, newBasisPoints, newMaxFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Transfer(&_UsdtAbi.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.Transfer(&_UsdtAbi.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.TransferFrom(&_UsdtAbi.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UsdtAbi.Contract.TransferFrom(&_UsdtAbi.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UsdtAbi *UsdtAbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UsdtAbi *UsdtAbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.TransferOwnership(&_UsdtAbi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UsdtAbi *UsdtAbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UsdtAbi.Contract.TransferOwnership(&_UsdtAbi.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UsdtAbi *UsdtAbiTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtAbi.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UsdtAbi *UsdtAbiSession) Unpause() (*types.Transaction, error) {
	return _UsdtAbi.Contract.Unpause(&_UsdtAbi.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UsdtAbi *UsdtAbiTransactorSession) Unpause() (*types.Transaction, error) {
	return _UsdtAbi.Contract.Unpause(&_UsdtAbi.TransactOpts)
}

// UsdtAbiAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the UsdtAbi contract.
type UsdtAbiAddedBlackListIterator struct {
	Event *UsdtAbiAddedBlackList // Event containing the contract specifics and raw log

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
func (it *UsdtAbiAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiAddedBlackList)
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
		it.Event = new(UsdtAbiAddedBlackList)
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
func (it *UsdtAbiAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiAddedBlackList represents a AddedBlackList event raised by the UsdtAbi contract.
type UsdtAbiAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_UsdtAbi *UsdtAbiFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*UsdtAbiAddedBlackListIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiAddedBlackListIterator{contract: _UsdtAbi.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_UsdtAbi *UsdtAbiFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *UsdtAbiAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiAddedBlackList)
				if err := _UsdtAbi.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_UsdtAbi *UsdtAbiFilterer) ParseAddedBlackList(log types.Log) (*UsdtAbiAddedBlackList, error) {
	event := new(UsdtAbiAddedBlackList)
	if err := _UsdtAbi.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UsdtAbi contract.
type UsdtAbiApprovalIterator struct {
	Event *UsdtAbiApproval // Event containing the contract specifics and raw log

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
func (it *UsdtAbiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiApproval)
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
		it.Event = new(UsdtAbiApproval)
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
func (it *UsdtAbiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiApproval represents a Approval event raised by the UsdtAbi contract.
type UsdtAbiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UsdtAbi *UsdtAbiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UsdtAbiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UsdtAbiApprovalIterator{contract: _UsdtAbi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UsdtAbi *UsdtAbiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UsdtAbiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiApproval)
				if err := _UsdtAbi.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UsdtAbi *UsdtAbiFilterer) ParseApproval(log types.Log) (*UsdtAbiApproval, error) {
	event := new(UsdtAbiApproval)
	if err := _UsdtAbi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiDeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the UsdtAbi contract.
type UsdtAbiDeprecateIterator struct {
	Event *UsdtAbiDeprecate // Event containing the contract specifics and raw log

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
func (it *UsdtAbiDeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiDeprecate)
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
		it.Event = new(UsdtAbiDeprecate)
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
func (it *UsdtAbiDeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiDeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiDeprecate represents a Deprecate event raised by the UsdtAbi contract.
type UsdtAbiDeprecate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_UsdtAbi *UsdtAbiFilterer) FilterDeprecate(opts *bind.FilterOpts) (*UsdtAbiDeprecateIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiDeprecateIterator{contract: _UsdtAbi.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_UsdtAbi *UsdtAbiFilterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *UsdtAbiDeprecate) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiDeprecate)
				if err := _UsdtAbi.contract.UnpackLog(event, "Deprecate", log); err != nil {
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

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_UsdtAbi *UsdtAbiFilterer) ParseDeprecate(log types.Log) (*UsdtAbiDeprecate, error) {
	event := new(UsdtAbiDeprecate)
	if err := _UsdtAbi.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the UsdtAbi contract.
type UsdtAbiDestroyedBlackFundsIterator struct {
	Event *UsdtAbiDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *UsdtAbiDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiDestroyedBlackFunds)
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
		it.Event = new(UsdtAbiDestroyedBlackFunds)
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
func (it *UsdtAbiDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the UsdtAbi contract.
type UsdtAbiDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_UsdtAbi *UsdtAbiFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*UsdtAbiDestroyedBlackFundsIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiDestroyedBlackFundsIterator{contract: _UsdtAbi.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_UsdtAbi *UsdtAbiFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *UsdtAbiDestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiDestroyedBlackFunds)
				if err := _UsdtAbi.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_UsdtAbi *UsdtAbiFilterer) ParseDestroyedBlackFunds(log types.Log) (*UsdtAbiDestroyedBlackFunds, error) {
	event := new(UsdtAbiDestroyedBlackFunds)
	if err := _UsdtAbi.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the UsdtAbi contract.
type UsdtAbiIssueIterator struct {
	Event *UsdtAbiIssue // Event containing the contract specifics and raw log

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
func (it *UsdtAbiIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiIssue)
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
		it.Event = new(UsdtAbiIssue)
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
func (it *UsdtAbiIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiIssue represents a Issue event raised by the UsdtAbi contract.
type UsdtAbiIssue struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_UsdtAbi *UsdtAbiFilterer) FilterIssue(opts *bind.FilterOpts) (*UsdtAbiIssueIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiIssueIterator{contract: _UsdtAbi.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_UsdtAbi *UsdtAbiFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *UsdtAbiIssue) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiIssue)
				if err := _UsdtAbi.contract.UnpackLog(event, "Issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_UsdtAbi *UsdtAbiFilterer) ParseIssue(log types.Log) (*UsdtAbiIssue, error) {
	event := new(UsdtAbiIssue)
	if err := _UsdtAbi.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiParamsIterator is returned from FilterParams and is used to iterate over the raw logs and unpacked data for Params events raised by the UsdtAbi contract.
type UsdtAbiParamsIterator struct {
	Event *UsdtAbiParams // Event containing the contract specifics and raw log

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
func (it *UsdtAbiParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiParams)
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
		it.Event = new(UsdtAbiParams)
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
func (it *UsdtAbiParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiParams represents a Params event raised by the UsdtAbi contract.
type UsdtAbiParams struct {
	FeeBasisPoints *big.Int
	MaxFee         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterParams is a free log retrieval operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_UsdtAbi *UsdtAbiFilterer) FilterParams(opts *bind.FilterOpts) (*UsdtAbiParamsIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiParamsIterator{contract: _UsdtAbi.contract, event: "Params", logs: logs, sub: sub}, nil
}

// WatchParams is a free log subscription operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_UsdtAbi *UsdtAbiFilterer) WatchParams(opts *bind.WatchOpts, sink chan<- *UsdtAbiParams) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiParams)
				if err := _UsdtAbi.contract.UnpackLog(event, "Params", log); err != nil {
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

// ParseParams is a log parse operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_UsdtAbi *UsdtAbiFilterer) ParseParams(log types.Log) (*UsdtAbiParams, error) {
	event := new(UsdtAbiParams)
	if err := _UsdtAbi.contract.UnpackLog(event, "Params", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the UsdtAbi contract.
type UsdtAbiPauseIterator struct {
	Event *UsdtAbiPause // Event containing the contract specifics and raw log

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
func (it *UsdtAbiPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiPause)
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
		it.Event = new(UsdtAbiPause)
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
func (it *UsdtAbiPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiPause represents a Pause event raised by the UsdtAbi contract.
type UsdtAbiPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_UsdtAbi *UsdtAbiFilterer) FilterPause(opts *bind.FilterOpts) (*UsdtAbiPauseIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiPauseIterator{contract: _UsdtAbi.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_UsdtAbi *UsdtAbiFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *UsdtAbiPause) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiPause)
				if err := _UsdtAbi.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_UsdtAbi *UsdtAbiFilterer) ParsePause(log types.Log) (*UsdtAbiPause, error) {
	event := new(UsdtAbiPause)
	if err := _UsdtAbi.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the UsdtAbi contract.
type UsdtAbiRedeemIterator struct {
	Event *UsdtAbiRedeem // Event containing the contract specifics and raw log

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
func (it *UsdtAbiRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiRedeem)
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
		it.Event = new(UsdtAbiRedeem)
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
func (it *UsdtAbiRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiRedeem represents a Redeem event raised by the UsdtAbi contract.
type UsdtAbiRedeem struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_UsdtAbi *UsdtAbiFilterer) FilterRedeem(opts *bind.FilterOpts) (*UsdtAbiRedeemIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiRedeemIterator{contract: _UsdtAbi.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_UsdtAbi *UsdtAbiFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *UsdtAbiRedeem) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiRedeem)
				if err := _UsdtAbi.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_UsdtAbi *UsdtAbiFilterer) ParseRedeem(log types.Log) (*UsdtAbiRedeem, error) {
	event := new(UsdtAbiRedeem)
	if err := _UsdtAbi.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the UsdtAbi contract.
type UsdtAbiRemovedBlackListIterator struct {
	Event *UsdtAbiRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *UsdtAbiRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiRemovedBlackList)
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
		it.Event = new(UsdtAbiRemovedBlackList)
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
func (it *UsdtAbiRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiRemovedBlackList represents a RemovedBlackList event raised by the UsdtAbi contract.
type UsdtAbiRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_UsdtAbi *UsdtAbiFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*UsdtAbiRemovedBlackListIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiRemovedBlackListIterator{contract: _UsdtAbi.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_UsdtAbi *UsdtAbiFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *UsdtAbiRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiRemovedBlackList)
				if err := _UsdtAbi.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_UsdtAbi *UsdtAbiFilterer) ParseRemovedBlackList(log types.Log) (*UsdtAbiRemovedBlackList, error) {
	event := new(UsdtAbiRemovedBlackList)
	if err := _UsdtAbi.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UsdtAbi contract.
type UsdtAbiTransferIterator struct {
	Event *UsdtAbiTransfer // Event containing the contract specifics and raw log

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
func (it *UsdtAbiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiTransfer)
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
		it.Event = new(UsdtAbiTransfer)
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
func (it *UsdtAbiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiTransfer represents a Transfer event raised by the UsdtAbi contract.
type UsdtAbiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UsdtAbi *UsdtAbiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UsdtAbiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UsdtAbiTransferIterator{contract: _UsdtAbi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UsdtAbi *UsdtAbiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UsdtAbiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiTransfer)
				if err := _UsdtAbi.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UsdtAbi *UsdtAbiFilterer) ParseTransfer(log types.Log) (*UsdtAbiTransfer, error) {
	event := new(UsdtAbiTransfer)
	if err := _UsdtAbi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtAbiUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the UsdtAbi contract.
type UsdtAbiUnpauseIterator struct {
	Event *UsdtAbiUnpause // Event containing the contract specifics and raw log

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
func (it *UsdtAbiUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtAbiUnpause)
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
		it.Event = new(UsdtAbiUnpause)
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
func (it *UsdtAbiUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtAbiUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtAbiUnpause represents a Unpause event raised by the UsdtAbi contract.
type UsdtAbiUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_UsdtAbi *UsdtAbiFilterer) FilterUnpause(opts *bind.FilterOpts) (*UsdtAbiUnpauseIterator, error) {

	logs, sub, err := _UsdtAbi.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &UsdtAbiUnpauseIterator{contract: _UsdtAbi.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_UsdtAbi *UsdtAbiFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *UsdtAbiUnpause) (event.Subscription, error) {

	logs, sub, err := _UsdtAbi.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtAbiUnpause)
				if err := _UsdtAbi.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_UsdtAbi *UsdtAbiFilterer) ParseUnpause(log types.Log) (*UsdtAbiUnpause, error) {
	event := new(UsdtAbiUnpause)
	if err := _UsdtAbi.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
