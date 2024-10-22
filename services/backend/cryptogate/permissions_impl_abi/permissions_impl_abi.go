// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissions_impl_abi

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

// PermissionsImplAbiMetaData contains all meta data concerning the PermissionsImplAbi contract.
var PermissionsImplABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_orgManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rolesManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_networkBootStatus\",\"type\":\"bool\"}],\"name\":\"PermissionsInitialized\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAdminAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"}],\"name\":\"addAdminNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_access\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_admin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNewRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pOrgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveBlacklistedAccountRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveBlacklistedNodeRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPolicyDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_breadth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depth\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_oAdminRole\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_networkBootStatus\",\"type\":\"bool\"}],\"name\":\"setMigrationPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"startBlacklistedAccountRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"startBlacklistedNodeRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"transactionAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateAccountStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PermissionsImplAbi is an auto generated Go binding around an Ethereum contract.
type PermissionsImplAbi struct {
	PermissionsImplAbiCaller     // Read-only binding to the contract
	PermissionsImplAbiTransactor // Write-only binding to the contract
	PermissionsImplAbiFilterer   // Log filterer for contract events
}

// PermissionsImplAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionsImplAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsImplAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionsImplAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsImplAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionsImplAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsImplAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionsImplAbiSession struct {
	Contract     *PermissionsImplAbi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermissionsImplAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionsImplAbiCallerSession struct {
	Contract *PermissionsImplAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PermissionsImplAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionsImplAbiTransactorSession struct {
	Contract     *PermissionsImplAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PermissionsImplAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionsImplAbiRaw struct {
	Contract *PermissionsImplAbi // Generic contract binding to access the raw methods on
}

// PermissionsImplAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionsImplAbiCallerRaw struct {
	Contract *PermissionsImplAbiCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionsImplAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionsImplAbiTransactorRaw struct {
	Contract *PermissionsImplAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionsImplAbi creates a new instance of PermissionsImplAbi, bound to a specific deployed contract.
func NewPermissionsImplAbi(address common.Address, backend bind.ContractBackend) (*PermissionsImplAbi, error) {
	contract, err := bindPermissionsImplAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionsImplAbi{PermissionsImplAbiCaller: PermissionsImplAbiCaller{contract: contract}, PermissionsImplAbiTransactor: PermissionsImplAbiTransactor{contract: contract}, PermissionsImplAbiFilterer: PermissionsImplAbiFilterer{contract: contract}}, nil
}

// NewPermissionsImplAbiCaller creates a new read-only instance of PermissionsImplAbi, bound to a specific deployed contract.
func NewPermissionsImplAbiCaller(address common.Address, caller bind.ContractCaller) (*PermissionsImplAbiCaller, error) {
	contract, err := bindPermissionsImplAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsImplAbiCaller{contract: contract}, nil
}

// NewPermissionsImplAbiTransactor creates a new write-only instance of PermissionsImplAbi, bound to a specific deployed contract.
func NewPermissionsImplAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionsImplAbiTransactor, error) {
	contract, err := bindPermissionsImplAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsImplAbiTransactor{contract: contract}, nil
}

// NewPermissionsImplAbiFilterer creates a new log filterer instance of PermissionsImplAbi, bound to a specific deployed contract.
func NewPermissionsImplAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionsImplAbiFilterer, error) {
	contract, err := bindPermissionsImplAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionsImplAbiFilterer{contract: contract}, nil
}

// bindPermissionsImplAbi binds a generic wrapper to an already deployed contract.
func bindPermissionsImplAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionsImplAbi *PermissionsImplAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionsImplAbi.Contract.PermissionsImplAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionsImplAbi *PermissionsImplAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.PermissionsImplAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionsImplAbi *PermissionsImplAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.PermissionsImplAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionsImplAbi *PermissionsImplAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionsImplAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionsImplAbi *PermissionsImplAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionsImplAbi *PermissionsImplAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.contract.Transact(opts, method, params...)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) ConnectionAllowed(opts *bind.CallOpts, _enodeId string, _ip string, _port uint16) (bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "connectionAllowed", _enodeId, _ip, _port)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) ConnectionAllowed(_enodeId string, _ip string, _port uint16) (bool, error) {
	return _PermissionsImplAbi.Contract.ConnectionAllowed(&_PermissionsImplAbi.CallOpts, _enodeId, _ip, _port)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) ConnectionAllowed(_enodeId string, _ip string, _port uint16) (bool, error) {
	return _PermissionsImplAbi.Contract.ConnectionAllowed(&_PermissionsImplAbi.CallOpts, _enodeId, _ip, _port)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) GetNetworkBootStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "getNetworkBootStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) GetNetworkBootStatus() (bool, error) {
	return _PermissionsImplAbi.Contract.GetNetworkBootStatus(&_PermissionsImplAbi.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) GetNetworkBootStatus() (bool, error) {
	return _PermissionsImplAbi.Contract.GetNetworkBootStatus(&_PermissionsImplAbi.CallOpts)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) view returns(string, string, address, uint256)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) GetPendingOp(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "getPendingOp", _orgId)

	if err != nil {
		return *new(string), *new(string), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) view returns(string, string, address, uint256)
func (_PermissionsImplAbi *PermissionsImplAbiSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermissionsImplAbi.Contract.GetPendingOp(&_PermissionsImplAbi.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) view returns(string, string, address, uint256)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermissionsImplAbi.Contract.GetPendingOp(&_PermissionsImplAbi.CallOpts, _orgId)
}

// GetPolicyDetails is a free data retrieval call binding the contract method 0xcc9ba6fa.
//
// Solidity: function getPolicyDetails() view returns(string, string, string, bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) GetPolicyDetails(opts *bind.CallOpts) (string, string, string, bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "getPolicyDetails")

	if err != nil {
		return *new(string), *new(string), *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetPolicyDetails is a free data retrieval call binding the contract method 0xcc9ba6fa.
//
// Solidity: function getPolicyDetails() view returns(string, string, string, bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) GetPolicyDetails() (string, string, string, bool, error) {
	return _PermissionsImplAbi.Contract.GetPolicyDetails(&_PermissionsImplAbi.CallOpts)
}

// GetPolicyDetails is a free data retrieval call binding the contract method 0xcc9ba6fa.
//
// Solidity: function getPolicyDetails() view returns(string, string, string, bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) GetPolicyDetails() (string, string, string, bool, error) {
	return _PermissionsImplAbi.Contract.GetPolicyDetails(&_PermissionsImplAbi.CallOpts)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "isNetworkAdmin", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermissionsImplAbi.Contract.IsNetworkAdmin(&_PermissionsImplAbi.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermissionsImplAbi.Contract.IsNetworkAdmin(&_PermissionsImplAbi.CallOpts, _account)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) IsOrgAdmin(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "isOrgAdmin", _account, _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermissionsImplAbi.Contract.IsOrgAdmin(&_PermissionsImplAbi.CallOpts, _account, _orgId)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermissionsImplAbi.Contract.IsOrgAdmin(&_PermissionsImplAbi.CallOpts, _account, _orgId)
}

// TransactionAllowed is a free data retrieval call binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address _sender, address _target, uint256 _value, uint256 _gasPrice, uint256 _gasLimit, bytes _payload) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) TransactionAllowed(opts *bind.CallOpts, _sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "transactionAllowed", _sender, _target, _value, _gasPrice, _gasLimit, _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransactionAllowed is a free data retrieval call binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address _sender, address _target, uint256 _value, uint256 _gasPrice, uint256 _gasLimit, bytes _payload) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	return _PermissionsImplAbi.Contract.TransactionAllowed(&_PermissionsImplAbi.CallOpts, _sender, _target, _value, _gasPrice, _gasLimit, _payload)
}

// TransactionAllowed is a free data retrieval call binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address _sender, address _target, uint256 _value, uint256 _gasPrice, uint256 _gasLimit, bytes _payload) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	return _PermissionsImplAbi.Contract.TransactionAllowed(&_PermissionsImplAbi.CallOpts, _sender, _target, _value, _gasPrice, _gasLimit, _payload)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCaller) ValidateAccount(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var out []interface{}
	err := _PermissionsImplAbi.contract.Call(opts, &out, "validateAccount", _account, _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermissionsImplAbi.Contract.ValidateAccount(&_PermissionsImplAbi.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiCallerSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermissionsImplAbi.Contract.ValidateAccount(&_PermissionsImplAbi.CallOpts, _account, _orgId)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _account) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AddAdminAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "addAdminAccount", _account)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _account) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AddAdminAccount(_account common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddAdminAccount(&_PermissionsImplAbi.TransactOpts, _account)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _account) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AddAdminAccount(_account common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddAdminAccount(&_PermissionsImplAbi.TransactOpts, _account)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x8683c7fe.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AddAdminNode(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "addAdminNode", _enodeId, _ip, _port, _raftport)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x8683c7fe.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddAdminNode(&_PermissionsImplAbi.TransactOpts, _enodeId, _ip, _port, _raftport)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x8683c7fe.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddAdminNode(&_PermissionsImplAbi.TransactOpts, _enodeId, _ip, _port, _raftport)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AddNewRole(opts *bind.TransactOpts, _roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "addNewRole", _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddNewRole(&_PermissionsImplAbi.TransactOpts, _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddNewRole(&_PermissionsImplAbi.TransactOpts, _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0xecad01d5.
//
// Solidity: function addNode(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AddNode(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "addNode", _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0xecad01d5.
//
// Solidity: function addNode(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddNode(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0xecad01d5.
//
// Solidity: function addNode(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddNode(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xe91b0e19.
//
// Solidity: function addOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AddOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "addOrg", _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xe91b0e19.
//
// Solidity: function addOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddOrg(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xe91b0e19.
//
// Solidity: function addOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddOrg(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x68a61273.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AddSubOrg(opts *bind.TransactOpts, _pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "addSubOrg", _pOrgId, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x68a61273.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddSubOrg(&_PermissionsImplAbi.TransactOpts, _pOrgId, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x68a61273.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AddSubOrg(&_PermissionsImplAbi.TransactOpts, _pOrgId, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) ApproveAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "approveAdminRole", _orgId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) ApproveAdminRole(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveAdminRole(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) ApproveAdminRole(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveAdminRole(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _caller)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x4b20f45f.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) ApproveBlacklistedAccountRecovery(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "approveBlacklistedAccountRecovery", _orgId, _account, _caller)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x4b20f45f.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveBlacklistedAccountRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _caller)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x4b20f45f.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveBlacklistedAccountRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _caller)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xa042bf40.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) ApproveBlacklistedNodeRecovery(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "approveBlacklistedNodeRecovery", _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xa042bf40.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveBlacklistedNodeRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xa042bf40.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveBlacklistedNodeRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xf75f0a06.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "approveOrg", _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xf75f0a06.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveOrg(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xf75f0a06.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveOrg(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) ApproveOrgStatus(opts *bind.TransactOpts, _orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "approveOrgStatus", _orgId, _action, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) ApproveOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveOrgStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _action, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) ApproveOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.ApproveOrgStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _action, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AssignAccountRole(opts *bind.TransactOpts, _account common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "assignAccountRole", _account, _orgId, _roleId, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AssignAccountRole(&_PermissionsImplAbi.TransactOpts, _account, _orgId, _roleId, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AssignAccountRole(&_PermissionsImplAbi.TransactOpts, _account, _orgId, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) AssignAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "assignAdminRole", _orgId, _account, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AssignAdminRole(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.AssignAdminRole(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _roleId, _caller)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) Init(opts *bind.TransactOpts, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "init", _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.Init(&_PermissionsImplAbi.TransactOpts, _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.Init(&_PermissionsImplAbi.TransactOpts, _breadth, _depth)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(string _roleId, string _orgId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "removeRole", _roleId, _orgId, _caller)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(string _roleId, string _orgId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) RemoveRole(_roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.RemoveRole(&_PermissionsImplAbi.TransactOpts, _roleId, _orgId, _caller)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(string _roleId, string _orgId, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) RemoveRole(_roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.RemoveRole(&_PermissionsImplAbi.TransactOpts, _roleId, _orgId, _caller)
}

// SetMigrationPolicy is a paid mutator transaction binding the contract method 0xf5ad584a.
//
// Solidity: function setMigrationPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole, bool _networkBootStatus) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) SetMigrationPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string, _networkBootStatus bool) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "setMigrationPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole, _networkBootStatus)
}

// SetMigrationPolicy is a paid mutator transaction binding the contract method 0xf5ad584a.
//
// Solidity: function setMigrationPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole, bool _networkBootStatus) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) SetMigrationPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string, _networkBootStatus bool) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.SetMigrationPolicy(&_PermissionsImplAbi.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole, _networkBootStatus)
}

// SetMigrationPolicy is a paid mutator transaction binding the contract method 0xf5ad584a.
//
// Solidity: function setMigrationPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole, bool _networkBootStatus) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) SetMigrationPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string, _networkBootStatus bool) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.SetMigrationPolicy(&_PermissionsImplAbi.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole, _networkBootStatus)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) SetPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "setPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.SetPolicy(&_PermissionsImplAbi.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.SetPolicy(&_PermissionsImplAbi.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x1c249912.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) StartBlacklistedAccountRecovery(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "startBlacklistedAccountRecovery", _orgId, _account, _caller)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x1c249912.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) StartBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.StartBlacklistedAccountRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _caller)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x1c249912.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) StartBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.StartBlacklistedAccountRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _caller)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xd621d957.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) StartBlacklistedNodeRecovery(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "startBlacklistedNodeRecovery", _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xd621d957.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.StartBlacklistedNodeRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xd621d957.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.StartBlacklistedNodeRecovery(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) UpdateAccountStatus(opts *bind.TransactOpts, _orgId string, _account common.Address, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "updateAccountStatus", _orgId, _account, _action, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateAccountStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _action, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateAccountStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _account, _action, _caller)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) UpdateNetworkBootStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "updateNetworkBootStatus")
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateNetworkBootStatus(&_PermissionsImplAbi.TransactOpts)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateNetworkBootStatus(&_PermissionsImplAbi.TransactOpts)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xb9b7fe6c.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "updateNodeStatus", _orgId, _enodeId, _ip, _port, _raftport, _action, _caller)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xb9b7fe6c.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateNodeStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _action, _caller)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xb9b7fe6c.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateNodeStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _action, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactor) UpdateOrgStatus(opts *bind.TransactOpts, _orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.contract.Transact(opts, "updateOrgStatus", _orgId, _action, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiSession) UpdateOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateOrgStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _action, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermissionsImplAbi *PermissionsImplAbiTransactorSession) UpdateOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermissionsImplAbi.Contract.UpdateOrgStatus(&_PermissionsImplAbi.TransactOpts, _orgId, _action, _caller)
}

// PermissionsImplAbiPermissionsInitializedIterator is returned from FilterPermissionsInitialized and is used to iterate over the raw logs and unpacked data for PermissionsInitialized events raised by the PermissionsImplAbi contract.
type PermissionsImplAbiPermissionsInitializedIterator struct {
	Event *PermissionsImplAbiPermissionsInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionsImplAbiPermissionsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsImplAbiPermissionsInitialized)
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
		it.Event = new(PermissionsImplAbiPermissionsInitialized)
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
func (it *PermissionsImplAbiPermissionsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsImplAbiPermissionsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsImplAbiPermissionsInitialized represents a PermissionsInitialized event raised by the PermissionsImplAbi contract.
type PermissionsImplAbiPermissionsInitialized struct {
	NetworkBootStatus bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPermissionsInitialized is a free log retrieval operation binding the contract event 0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf.
//
// Solidity: event PermissionsInitialized(bool _networkBootStatus)
func (_PermissionsImplAbi *PermissionsImplAbiFilterer) FilterPermissionsInitialized(opts *bind.FilterOpts) (*PermissionsImplAbiPermissionsInitializedIterator, error) {

	logs, sub, err := _PermissionsImplAbi.contract.FilterLogs(opts, "PermissionsInitialized")
	if err != nil {
		return nil, err
	}
	return &PermissionsImplAbiPermissionsInitializedIterator{contract: _PermissionsImplAbi.contract, event: "PermissionsInitialized", logs: logs, sub: sub}, nil
}

// WatchPermissionsInitialized is a free log subscription operation binding the contract event 0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf.
//
// Solidity: event PermissionsInitialized(bool _networkBootStatus)
func (_PermissionsImplAbi *PermissionsImplAbiFilterer) WatchPermissionsInitialized(opts *bind.WatchOpts, sink chan<- *PermissionsImplAbiPermissionsInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionsImplAbi.contract.WatchLogs(opts, "PermissionsInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsImplAbiPermissionsInitialized)
				if err := _PermissionsImplAbi.contract.UnpackLog(event, "PermissionsInitialized", log); err != nil {
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

// ParsePermissionsInitialized is a log parse operation binding the contract event 0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf.
//
// Solidity: event PermissionsInitialized(bool _networkBootStatus)
func (_PermissionsImplAbi *PermissionsImplAbiFilterer) ParsePermissionsInitialized(log types.Log) (*PermissionsImplAbiPermissionsInitialized, error) {
	event := new(PermissionsImplAbiPermissionsInitialized)
	if err := _PermissionsImplAbi.contract.UnpackLog(event, "PermissionsInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
