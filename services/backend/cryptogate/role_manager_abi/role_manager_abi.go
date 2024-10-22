// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package role_manager_abi

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

// RoleManagerAbiMetaData contains all meta data concerning the RoleManagerAbi contract.
var RoleManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isVoter\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"RoleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isVoter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"addRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfRoles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getRoleDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"admin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rIndex\",\"type\":\"uint256\"}],\"name\":\"getRoleDetailsFromIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"admin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"isAdminRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"isVoterRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"roleAccess\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"roleExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ultParent\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_typeOfTxn\",\"type\":\"uint256\"}],\"name\":\"transactionAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RoleManagerAbi is an auto generated Go binding around an Ethereum contract.
type RoleManagerAbi struct {
	RoleManagerAbiCaller     // Read-only binding to the contract
	RoleManagerAbiTransactor // Write-only binding to the contract
	RoleManagerAbiFilterer   // Log filterer for contract events
}

// RoleManagerAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleManagerAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleManagerAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleManagerAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleManagerAbiSession struct {
	Contract     *RoleManagerAbi   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleManagerAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleManagerAbiCallerSession struct {
	Contract *RoleManagerAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RoleManagerAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleManagerAbiTransactorSession struct {
	Contract     *RoleManagerAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RoleManagerAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleManagerAbiRaw struct {
	Contract *RoleManagerAbi // Generic contract binding to access the raw methods on
}

// RoleManagerAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleManagerAbiCallerRaw struct {
	Contract *RoleManagerAbiCaller // Generic read-only contract binding to access the raw methods on
}

// RoleManagerAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleManagerAbiTransactorRaw struct {
	Contract *RoleManagerAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleManagerAbi creates a new instance of RoleManagerAbi, bound to a specific deployed contract.
func NewRoleManagerAbi(address common.Address, backend bind.ContractBackend) (*RoleManagerAbi, error) {
	contract, err := bindRoleManagerAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleManagerAbi{RoleManagerAbiCaller: RoleManagerAbiCaller{contract: contract}, RoleManagerAbiTransactor: RoleManagerAbiTransactor{contract: contract}, RoleManagerAbiFilterer: RoleManagerAbiFilterer{contract: contract}}, nil
}

// NewRoleManagerAbiCaller creates a new read-only instance of RoleManagerAbi, bound to a specific deployed contract.
func NewRoleManagerAbiCaller(address common.Address, caller bind.ContractCaller) (*RoleManagerAbiCaller, error) {
	contract, err := bindRoleManagerAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagerAbiCaller{contract: contract}, nil
}

// NewRoleManagerAbiTransactor creates a new write-only instance of RoleManagerAbi, bound to a specific deployed contract.
func NewRoleManagerAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleManagerAbiTransactor, error) {
	contract, err := bindRoleManagerAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagerAbiTransactor{contract: contract}, nil
}

// NewRoleManagerAbiFilterer creates a new log filterer instance of RoleManagerAbi, bound to a specific deployed contract.
func NewRoleManagerAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleManagerAbiFilterer, error) {
	contract, err := bindRoleManagerAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleManagerAbiFilterer{contract: contract}, nil
}

// bindRoleManagerAbi binds a generic wrapper to an already deployed contract.
func bindRoleManagerAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManagerAbi *RoleManagerAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleManagerAbi.Contract.RoleManagerAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManagerAbi *RoleManagerAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.RoleManagerAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManagerAbi *RoleManagerAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.RoleManagerAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManagerAbi *RoleManagerAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleManagerAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManagerAbi *RoleManagerAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManagerAbi *RoleManagerAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.contract.Transact(opts, method, params...)
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() view returns(uint256)
func (_RoleManagerAbi *RoleManagerAbiCaller) GetNumberOfRoles(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "getNumberOfRoles")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() view returns(uint256)
func (_RoleManagerAbi *RoleManagerAbiSession) GetNumberOfRoles() (*big.Int, error) {
	return _RoleManagerAbi.Contract.GetNumberOfRoles(&_RoleManagerAbi.CallOpts)
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() view returns(uint256)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) GetNumberOfRoles() (*big.Int, error) {
	return _RoleManagerAbi.Contract.GetNumberOfRoles(&_RoleManagerAbi.CallOpts)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(string _roleId, string _orgId) view returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManagerAbi *RoleManagerAbiCaller) GetRoleDetails(opts *bind.CallOpts, _roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "getRoleDetails", _roleId, _orgId)

	outstruct := new(struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Admin      bool
		Active     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoleId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.OrgId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AccessType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Voter = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Admin = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(string _roleId, string _orgId) view returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManagerAbi *RoleManagerAbiSession) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManagerAbi.Contract.GetRoleDetails(&_RoleManagerAbi.CallOpts, _roleId, _orgId)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(string _roleId, string _orgId) view returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManagerAbi.Contract.GetRoleDetails(&_RoleManagerAbi.CallOpts, _roleId, _orgId)
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(uint256 _rIndex) view returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManagerAbi *RoleManagerAbiCaller) GetRoleDetailsFromIndex(opts *bind.CallOpts, _rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "getRoleDetailsFromIndex", _rIndex)

	outstruct := new(struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Admin      bool
		Active     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoleId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.OrgId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AccessType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Voter = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Admin = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(uint256 _rIndex) view returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManagerAbi *RoleManagerAbiSession) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManagerAbi.Contract.GetRoleDetailsFromIndex(&_RoleManagerAbi.CallOpts, _rIndex)
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(uint256 _rIndex) view returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManagerAbi.Contract.GetRoleDetailsFromIndex(&_RoleManagerAbi.CallOpts, _rIndex)
}

// IsAdminRole is a free data retrieval call binding the contract method 0xbe322e54.
//
// Solidity: function isAdminRole(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCaller) IsAdminRole(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (bool, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "isAdminRole", _roleId, _orgId, _ultParent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdminRole is a free data retrieval call binding the contract method 0xbe322e54.
//
// Solidity: function isAdminRole(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiSession) IsAdminRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManagerAbi.Contract.IsAdminRole(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// IsAdminRole is a free data retrieval call binding the contract method 0xbe322e54.
//
// Solidity: function isAdminRole(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) IsAdminRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManagerAbi.Contract.IsAdminRole(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// IsVoterRole is a free data retrieval call binding the contract method 0xdeb16ba7.
//
// Solidity: function isVoterRole(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCaller) IsVoterRole(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (bool, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "isVoterRole", _roleId, _orgId, _ultParent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterRole is a free data retrieval call binding the contract method 0xdeb16ba7.
//
// Solidity: function isVoterRole(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiSession) IsVoterRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManagerAbi.Contract.IsVoterRole(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// IsVoterRole is a free data retrieval call binding the contract method 0xdeb16ba7.
//
// Solidity: function isVoterRole(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) IsVoterRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManagerAbi.Contract.IsVoterRole(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// RoleAccess is a free data retrieval call binding the contract method 0xcfc83dfa.
//
// Solidity: function roleAccess(string _roleId, string _orgId, string _ultParent) view returns(uint256)
func (_RoleManagerAbi *RoleManagerAbiCaller) RoleAccess(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (*big.Int, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "roleAccess", _roleId, _orgId, _ultParent)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoleAccess is a free data retrieval call binding the contract method 0xcfc83dfa.
//
// Solidity: function roleAccess(string _roleId, string _orgId, string _ultParent) view returns(uint256)
func (_RoleManagerAbi *RoleManagerAbiSession) RoleAccess(_roleId string, _orgId string, _ultParent string) (*big.Int, error) {
	return _RoleManagerAbi.Contract.RoleAccess(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// RoleAccess is a free data retrieval call binding the contract method 0xcfc83dfa.
//
// Solidity: function roleAccess(string _roleId, string _orgId, string _ultParent) view returns(uint256)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) RoleAccess(_roleId string, _orgId string, _ultParent string) (*big.Int, error) {
	return _RoleManagerAbi.Contract.RoleAccess(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// RoleExists is a free data retrieval call binding the contract method 0xabf5739f.
//
// Solidity: function roleExists(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCaller) RoleExists(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (bool, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "roleExists", _roleId, _orgId, _ultParent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RoleExists is a free data retrieval call binding the contract method 0xabf5739f.
//
// Solidity: function roleExists(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiSession) RoleExists(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManagerAbi.Contract.RoleExists(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// RoleExists is a free data retrieval call binding the contract method 0xabf5739f.
//
// Solidity: function roleExists(string _roleId, string _orgId, string _ultParent) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) RoleExists(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManagerAbi.Contract.RoleExists(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent)
}

// TransactionAllowed is a free data retrieval call binding the contract method 0xd1f77866.
//
// Solidity: function transactionAllowed(string _roleId, string _orgId, string _ultParent, uint256 _typeOfTxn) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCaller) TransactionAllowed(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string, _typeOfTxn *big.Int) (bool, error) {
	var out []interface{}
	err := _RoleManagerAbi.contract.Call(opts, &out, "transactionAllowed", _roleId, _orgId, _ultParent, _typeOfTxn)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransactionAllowed is a free data retrieval call binding the contract method 0xd1f77866.
//
// Solidity: function transactionAllowed(string _roleId, string _orgId, string _ultParent, uint256 _typeOfTxn) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiSession) TransactionAllowed(_roleId string, _orgId string, _ultParent string, _typeOfTxn *big.Int) (bool, error) {
	return _RoleManagerAbi.Contract.TransactionAllowed(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent, _typeOfTxn)
}

// TransactionAllowed is a free data retrieval call binding the contract method 0xd1f77866.
//
// Solidity: function transactionAllowed(string _roleId, string _orgId, string _ultParent, uint256 _typeOfTxn) view returns(bool)
func (_RoleManagerAbi *RoleManagerAbiCallerSession) TransactionAllowed(_roleId string, _orgId string, _ultParent string, _typeOfTxn *big.Int) (bool, error) {
	return _RoleManagerAbi.Contract.TransactionAllowed(&_RoleManagerAbi.CallOpts, _roleId, _orgId, _ultParent, _typeOfTxn)
}

// AddRole is a paid mutator transaction binding the contract method 0x7b713579.
//
// Solidity: function addRole(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin) returns()
func (_RoleManagerAbi *RoleManagerAbiTransactor) AddRole(opts *bind.TransactOpts, _roleId string, _orgId string, _baseAccess *big.Int, _isVoter bool, _isAdmin bool) (*types.Transaction, error) {
	return _RoleManagerAbi.contract.Transact(opts, "addRole", _roleId, _orgId, _baseAccess, _isVoter, _isAdmin)
}

// AddRole is a paid mutator transaction binding the contract method 0x7b713579.
//
// Solidity: function addRole(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin) returns()
func (_RoleManagerAbi *RoleManagerAbiSession) AddRole(_roleId string, _orgId string, _baseAccess *big.Int, _isVoter bool, _isAdmin bool) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.AddRole(&_RoleManagerAbi.TransactOpts, _roleId, _orgId, _baseAccess, _isVoter, _isAdmin)
}

// AddRole is a paid mutator transaction binding the contract method 0x7b713579.
//
// Solidity: function addRole(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin) returns()
func (_RoleManagerAbi *RoleManagerAbiTransactorSession) AddRole(_roleId string, _orgId string, _baseAccess *big.Int, _isVoter bool, _isAdmin bool) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.AddRole(&_RoleManagerAbi.TransactOpts, _roleId, _orgId, _baseAccess, _isVoter, _isAdmin)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_RoleManagerAbi *RoleManagerAbiTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManagerAbi.contract.Transact(opts, "removeRole", _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_RoleManagerAbi *RoleManagerAbiSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.RemoveRole(&_RoleManagerAbi.TransactOpts, _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_RoleManagerAbi *RoleManagerAbiTransactorSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManagerAbi.Contract.RemoveRole(&_RoleManagerAbi.TransactOpts, _roleId, _orgId)
}

// RoleManagerAbiRoleCreatedIterator is returned from FilterRoleCreated and is used to iterate over the raw logs and unpacked data for RoleCreated events raised by the RoleManagerAbi contract.
type RoleManagerAbiRoleCreatedIterator struct {
	Event *RoleManagerAbiRoleCreated // Event containing the contract specifics and raw log

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
func (it *RoleManagerAbiRoleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagerAbiRoleCreated)
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
		it.Event = new(RoleManagerAbiRoleCreated)
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
func (it *RoleManagerAbiRoleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagerAbiRoleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagerAbiRoleCreated represents a RoleCreated event raised by the RoleManagerAbi contract.
type RoleManagerAbiRoleCreated struct {
	RoleId     string
	OrgId      string
	BaseAccess *big.Int
	IsVoter    bool
	IsAdmin    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoleCreated is a free log retrieval operation binding the contract event 0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c.
//
// Solidity: event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin)
func (_RoleManagerAbi *RoleManagerAbiFilterer) FilterRoleCreated(opts *bind.FilterOpts) (*RoleManagerAbiRoleCreatedIterator, error) {

	logs, sub, err := _RoleManagerAbi.contract.FilterLogs(opts, "RoleCreated")
	if err != nil {
		return nil, err
	}
	return &RoleManagerAbiRoleCreatedIterator{contract: _RoleManagerAbi.contract, event: "RoleCreated", logs: logs, sub: sub}, nil
}

// WatchRoleCreated is a free log subscription operation binding the contract event 0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c.
//
// Solidity: event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin)
func (_RoleManagerAbi *RoleManagerAbiFilterer) WatchRoleCreated(opts *bind.WatchOpts, sink chan<- *RoleManagerAbiRoleCreated) (event.Subscription, error) {

	logs, sub, err := _RoleManagerAbi.contract.WatchLogs(opts, "RoleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagerAbiRoleCreated)
				if err := _RoleManagerAbi.contract.UnpackLog(event, "RoleCreated", log); err != nil {
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

// ParseRoleCreated is a log parse operation binding the contract event 0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c.
//
// Solidity: event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin)
func (_RoleManagerAbi *RoleManagerAbiFilterer) ParseRoleCreated(log types.Log) (*RoleManagerAbiRoleCreated, error) {
	event := new(RoleManagerAbiRoleCreated)
	if err := _RoleManagerAbi.contract.UnpackLog(event, "RoleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleManagerAbiRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RoleManagerAbi contract.
type RoleManagerAbiRoleRevokedIterator struct {
	Event *RoleManagerAbiRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoleManagerAbiRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagerAbiRoleRevoked)
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
		it.Event = new(RoleManagerAbiRoleRevoked)
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
func (it *RoleManagerAbiRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagerAbiRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagerAbiRoleRevoked represents a RoleRevoked event raised by the RoleManagerAbi contract.
type RoleManagerAbiRoleRevoked struct {
	RoleId string
	OrgId  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: event RoleRevoked(string _roleId, string _orgId)
func (_RoleManagerAbi *RoleManagerAbiFilterer) FilterRoleRevoked(opts *bind.FilterOpts) (*RoleManagerAbiRoleRevokedIterator, error) {

	logs, sub, err := _RoleManagerAbi.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &RoleManagerAbiRoleRevokedIterator{contract: _RoleManagerAbi.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: event RoleRevoked(string _roleId, string _orgId)
func (_RoleManagerAbi *RoleManagerAbiFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoleManagerAbiRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _RoleManagerAbi.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagerAbiRoleRevoked)
				if err := _RoleManagerAbi.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: event RoleRevoked(string _roleId, string _orgId)
func (_RoleManagerAbi *RoleManagerAbiFilterer) ParseRoleRevoked(log types.Log) (*RoleManagerAbiRoleRevoked, error) {
	event := new(RoleManagerAbiRoleRevoked)
	if err := _RoleManagerAbi.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
