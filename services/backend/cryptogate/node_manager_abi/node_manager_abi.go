// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node_manager_abi

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

// NodeManagerAbiMetaData contains all meta data concerning the NodeManagerAbi contract.
var NodeManagerAbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeRecoveryCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeRecoveryInitiated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addAdminNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addOrgNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"enodeId\",\"type\":\"string\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getNodeDetailsFromIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NodeManagerAbiABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeManagerAbiMetaData.ABI instead.
var NodeManagerAbiABI = NodeManagerAbiMetaData.ABI

// NodeManagerAbi is an auto generated Go binding around an Ethereum contract.
type NodeManagerAbi struct {
	NodeManagerAbiCaller     // Read-only binding to the contract
	NodeManagerAbiTransactor // Write-only binding to the contract
	NodeManagerAbiFilterer   // Log filterer for contract events
}

// NodeManagerAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerAbiSession struct {
	Contract     *NodeManagerAbi   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerAbiCallerSession struct {
	Contract *NodeManagerAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodeManagerAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerAbiTransactorSession struct {
	Contract     *NodeManagerAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodeManagerAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerAbiRaw struct {
	Contract *NodeManagerAbi // Generic contract binding to access the raw methods on
}

// NodeManagerAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerAbiCallerRaw struct {
	Contract *NodeManagerAbiCaller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerAbiTransactorRaw struct {
	Contract *NodeManagerAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManagerAbi creates a new instance of NodeManagerAbi, bound to a specific deployed contract.
func NewNodeManagerAbi(address common.Address, backend bind.ContractBackend) (*NodeManagerAbi, error) {
	contract, err := bindNodeManagerAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbi{NodeManagerAbiCaller: NodeManagerAbiCaller{contract: contract}, NodeManagerAbiTransactor: NodeManagerAbiTransactor{contract: contract}, NodeManagerAbiFilterer: NodeManagerAbiFilterer{contract: contract}}, nil
}

// NewNodeManagerAbiCaller creates a new read-only instance of NodeManagerAbi, bound to a specific deployed contract.
func NewNodeManagerAbiCaller(address common.Address, caller bind.ContractCaller) (*NodeManagerAbiCaller, error) {
	contract, err := bindNodeManagerAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiCaller{contract: contract}, nil
}

// NewNodeManagerAbiTransactor creates a new write-only instance of NodeManagerAbi, bound to a specific deployed contract.
func NewNodeManagerAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerAbiTransactor, error) {
	contract, err := bindNodeManagerAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiTransactor{contract: contract}, nil
}

// NewNodeManagerAbiFilterer creates a new log filterer instance of NodeManagerAbi, bound to a specific deployed contract.
func NewNodeManagerAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerAbiFilterer, error) {
	contract, err := bindNodeManagerAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiFilterer{contract: contract}, nil
}

// bindNodeManagerAbi binds a generic wrapper to an already deployed contract.
func bindNodeManagerAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeManagerAbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManagerAbi *NodeManagerAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManagerAbi.Contract.NodeManagerAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManagerAbi *NodeManagerAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.NodeManagerAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManagerAbi *NodeManagerAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.NodeManagerAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManagerAbi *NodeManagerAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManagerAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManagerAbi *NodeManagerAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManagerAbi *NodeManagerAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.contract.Transact(opts, method, params...)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_NodeManagerAbi *NodeManagerAbiCaller) ConnectionAllowed(opts *bind.CallOpts, _enodeId string, _ip string, _port uint16) (bool, error) {
	var out []interface{}
	err := _NodeManagerAbi.contract.Call(opts, &out, "connectionAllowed", _enodeId, _ip, _port)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_NodeManagerAbi *NodeManagerAbiSession) ConnectionAllowed(_enodeId string, _ip string, _port uint16) (bool, error) {
	return _NodeManagerAbi.Contract.ConnectionAllowed(&_NodeManagerAbi.CallOpts, _enodeId, _ip, _port)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_NodeManagerAbi *NodeManagerAbiCallerSession) ConnectionAllowed(_enodeId string, _ip string, _port uint16) (bool, error) {
	return _NodeManagerAbi.Contract.ConnectionAllowed(&_NodeManagerAbi.CallOpts, _enodeId, _ip, _port)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(string enodeId) view returns(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus)
func (_NodeManagerAbi *NodeManagerAbiCaller) GetNodeDetails(opts *bind.CallOpts, enodeId string) (struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	var out []interface{}
	err := _NodeManagerAbi.contract.Call(opts, &out, "getNodeDetails", enodeId)

	outstruct := new(struct {
		OrgId      string
		EnodeId    string
		Ip         string
		Port       uint16
		Raftport   uint16
		NodeStatus *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrgId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EnodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Ip = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Port = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Raftport = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.NodeStatus = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(string enodeId) view returns(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus)
func (_NodeManagerAbi *NodeManagerAbiSession) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	return _NodeManagerAbi.Contract.GetNodeDetails(&_NodeManagerAbi.CallOpts, enodeId)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(string enodeId) view returns(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus)
func (_NodeManagerAbi *NodeManagerAbiCallerSession) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	return _NodeManagerAbi.Contract.GetNodeDetails(&_NodeManagerAbi.CallOpts, enodeId)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(uint256 _nodeIndex) view returns(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus)
func (_NodeManagerAbi *NodeManagerAbiCaller) GetNodeDetailsFromIndex(opts *bind.CallOpts, _nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	var out []interface{}
	err := _NodeManagerAbi.contract.Call(opts, &out, "getNodeDetailsFromIndex", _nodeIndex)

	outstruct := new(struct {
		OrgId      string
		EnodeId    string
		Ip         string
		Port       uint16
		Raftport   uint16
		NodeStatus *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrgId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EnodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Ip = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Port = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Raftport = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.NodeStatus = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(uint256 _nodeIndex) view returns(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus)
func (_NodeManagerAbi *NodeManagerAbiSession) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	return _NodeManagerAbi.Contract.GetNodeDetailsFromIndex(&_NodeManagerAbi.CallOpts, _nodeIndex)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(uint256 _nodeIndex) view returns(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus)
func (_NodeManagerAbi *NodeManagerAbiCallerSession) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	return _NodeManagerAbi.Contract.GetNodeDetailsFromIndex(&_NodeManagerAbi.CallOpts, _nodeIndex)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() view returns(uint256)
func (_NodeManagerAbi *NodeManagerAbiCaller) GetNumberOfNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerAbi.contract.Call(opts, &out, "getNumberOfNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() view returns(uint256)
func (_NodeManagerAbi *NodeManagerAbiSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManagerAbi.Contract.GetNumberOfNodes(&_NodeManagerAbi.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() view returns(uint256)
func (_NodeManagerAbi *NodeManagerAbiCallerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManagerAbi.Contract.GetNumberOfNodes(&_NodeManagerAbi.CallOpts)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x4530abe1.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactor) AddAdminNode(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.contract.Transact(opts, "addAdminNode", _enodeId, _ip, _port, _raftport, _orgId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x4530abe1.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiSession) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.AddAdminNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x4530abe1.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactorSession) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.AddAdminNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0x549583df.
//
// Solidity: function addNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactor) AddNode(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.contract.Transact(opts, "addNode", _enodeId, _ip, _port, _raftport, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0x549583df.
//
// Solidity: function addNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiSession) AddNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.AddNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0x549583df.
//
// Solidity: function addNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactorSession) AddNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.AddNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x4c573311.
//
// Solidity: function addOrgNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactor) AddOrgNode(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.contract.Transact(opts, "addOrgNode", _enodeId, _ip, _port, _raftport, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x4c573311.
//
// Solidity: function addOrgNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiSession) AddOrgNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.AddOrgNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x4c573311.
//
// Solidity: function addOrgNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactorSession) AddOrgNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.AddOrgNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xf82e08ac.
//
// Solidity: function approveNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactor) ApproveNode(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.contract.Transact(opts, "approveNode", _enodeId, _ip, _port, _raftport, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xf82e08ac.
//
// Solidity: function approveNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiSession) ApproveNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.ApproveNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xf82e08ac.
//
// Solidity: function approveNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactorSession) ApproveNode(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.ApproveNode(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x37d50b27.
//
// Solidity: function updateNodeStatus(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId, uint256 _action) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManagerAbi.contract.Transact(opts, "updateNodeStatus", _enodeId, _ip, _port, _raftport, _orgId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x37d50b27.
//
// Solidity: function updateNodeStatus(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId, uint256 _action) returns()
func (_NodeManagerAbi *NodeManagerAbiSession) UpdateNodeStatus(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.UpdateNodeStatus(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x37d50b27.
//
// Solidity: function updateNodeStatus(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId, uint256 _action) returns()
func (_NodeManagerAbi *NodeManagerAbiTransactorSession) UpdateNodeStatus(_enodeId string, _ip string, _port uint16, _raftport uint16, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManagerAbi.Contract.UpdateNodeStatus(&_NodeManagerAbi.TransactOpts, _enodeId, _ip, _port, _raftport, _orgId, _action)
}

// NodeManagerAbiNodeActivatedIterator is returned from FilterNodeActivated and is used to iterate over the raw logs and unpacked data for NodeActivated events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeActivatedIterator struct {
	Event *NodeManagerAbiNodeActivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeActivated)
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
		it.Event = new(NodeManagerAbiNodeActivated)
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
func (it *NodeManagerAbiNodeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeActivated represents a NodeActivated event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeActivated struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeActivated is a free log retrieval operation binding the contract event 0xfb98f62dea866f0c373574c8523f611d0db1d8f19cc1b95d07a221d36a6a45de.
//
// Solidity: event NodeActivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeActivated(opts *bind.FilterOpts) (*NodeManagerAbiNodeActivatedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeActivatedIterator{contract: _NodeManagerAbi.contract, event: "NodeActivated", logs: logs, sub: sub}, nil
}

// WatchNodeActivated is a free log subscription operation binding the contract event 0xfb98f62dea866f0c373574c8523f611d0db1d8f19cc1b95d07a221d36a6a45de.
//
// Solidity: event NodeActivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeActivated(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeActivated) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeActivated)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeActivated", log); err != nil {
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

// ParseNodeActivated is a log parse operation binding the contract event 0xfb98f62dea866f0c373574c8523f611d0db1d8f19cc1b95d07a221d36a6a45de.
//
// Solidity: event NodeActivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeActivated(log types.Log) (*NodeManagerAbiNodeActivated, error) {
	event := new(NodeManagerAbiNodeActivated)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerAbiNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeApprovedIterator struct {
	Event *NodeManagerAbiNodeApproved // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeApproved)
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
		it.Event = new(NodeManagerAbiNodeApproved)
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
func (it *NodeManagerAbiNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeApproved represents a NodeApproved event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeApproved struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0x9394c836a3325586270659f6aa3b9f835abca9afe7fec5abfc69760bb12bce0d.
//
// Solidity: event NodeApproved(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*NodeManagerAbiNodeApprovedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeApprovedIterator{contract: _NodeManagerAbi.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0x9394c836a3325586270659f6aa3b9f835abca9afe7fec5abfc69760bb12bce0d.
//
// Solidity: event NodeApproved(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeApproved) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeApproved)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeApproved", log); err != nil {
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

// ParseNodeApproved is a log parse operation binding the contract event 0x9394c836a3325586270659f6aa3b9f835abca9afe7fec5abfc69760bb12bce0d.
//
// Solidity: event NodeApproved(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeApproved(log types.Log) (*NodeManagerAbiNodeApproved, error) {
	event := new(NodeManagerAbiNodeApproved)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerAbiNodeBlacklistedIterator is returned from FilterNodeBlacklisted and is used to iterate over the raw logs and unpacked data for NodeBlacklisted events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeBlacklistedIterator struct {
	Event *NodeManagerAbiNodeBlacklisted // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeBlacklisted)
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
		it.Event = new(NodeManagerAbiNodeBlacklisted)
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
func (it *NodeManagerAbiNodeBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeBlacklisted represents a NodeBlacklisted event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeBlacklisted struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeBlacklisted is a free log retrieval operation binding the contract event 0x25300d4d785e654bc9b7979700cfa0fdc9ace890a46841fecfce661fd2c41a33.
//
// Solidity: event NodeBlacklisted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeBlacklisted(opts *bind.FilterOpts) (*NodeManagerAbiNodeBlacklistedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeBlacklistedIterator{contract: _NodeManagerAbi.contract, event: "NodeBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNodeBlacklisted is a free log subscription operation binding the contract event 0x25300d4d785e654bc9b7979700cfa0fdc9ace890a46841fecfce661fd2c41a33.
//
// Solidity: event NodeBlacklisted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeBlacklisted(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeBlacklisted) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeBlacklisted)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
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

// ParseNodeBlacklisted is a log parse operation binding the contract event 0x25300d4d785e654bc9b7979700cfa0fdc9ace890a46841fecfce661fd2c41a33.
//
// Solidity: event NodeBlacklisted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeBlacklisted(log types.Log) (*NodeManagerAbiNodeBlacklisted, error) {
	event := new(NodeManagerAbiNodeBlacklisted)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerAbiNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeDeactivatedIterator struct {
	Event *NodeManagerAbiNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeDeactivated)
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
		it.Event = new(NodeManagerAbiNodeDeactivated)
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
func (it *NodeManagerAbiNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeDeactivated represents a NodeDeactivated event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeDeactivated struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xf631019be71bc682c59150635d714061185232e98e60de8bdd87bbee239cc5c8.
//
// Solidity: event NodeDeactivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeDeactivated(opts *bind.FilterOpts) (*NodeManagerAbiNodeDeactivatedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeDeactivatedIterator{contract: _NodeManagerAbi.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xf631019be71bc682c59150635d714061185232e98e60de8bdd87bbee239cc5c8.
//
// Solidity: event NodeDeactivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeDeactivated) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeDeactivated)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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

// ParseNodeDeactivated is a log parse operation binding the contract event 0xf631019be71bc682c59150635d714061185232e98e60de8bdd87bbee239cc5c8.
//
// Solidity: event NodeDeactivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeDeactivated(log types.Log) (*NodeManagerAbiNodeDeactivated, error) {
	event := new(NodeManagerAbiNodeDeactivated)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerAbiNodeProposedIterator is returned from FilterNodeProposed and is used to iterate over the raw logs and unpacked data for NodeProposed events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeProposedIterator struct {
	Event *NodeManagerAbiNodeProposed // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeProposed)
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
		it.Event = new(NodeManagerAbiNodeProposed)
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
func (it *NodeManagerAbiNodeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeProposed represents a NodeProposed event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeProposed struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeProposed is a free log retrieval operation binding the contract event 0xf9bad9f8a2dccc52fad61273a7fd673335b420319506c19b87df9ce7a19732da.
//
// Solidity: event NodeProposed(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeProposed(opts *bind.FilterOpts) (*NodeManagerAbiNodeProposedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeProposedIterator{contract: _NodeManagerAbi.contract, event: "NodeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeProposed is a free log subscription operation binding the contract event 0xf9bad9f8a2dccc52fad61273a7fd673335b420319506c19b87df9ce7a19732da.
//
// Solidity: event NodeProposed(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeProposed(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeProposed) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeProposed)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeProposed", log); err != nil {
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

// ParseNodeProposed is a log parse operation binding the contract event 0xf9bad9f8a2dccc52fad61273a7fd673335b420319506c19b87df9ce7a19732da.
//
// Solidity: event NodeProposed(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeProposed(log types.Log) (*NodeManagerAbiNodeProposed, error) {
	event := new(NodeManagerAbiNodeProposed)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerAbiNodeRecoveryCompletedIterator is returned from FilterNodeRecoveryCompleted and is used to iterate over the raw logs and unpacked data for NodeRecoveryCompleted events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeRecoveryCompletedIterator struct {
	Event *NodeManagerAbiNodeRecoveryCompleted // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeRecoveryCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeRecoveryCompleted)
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
		it.Event = new(NodeManagerAbiNodeRecoveryCompleted)
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
func (it *NodeManagerAbiNodeRecoveryCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeRecoveryCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeRecoveryCompleted represents a NodeRecoveryCompleted event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeRecoveryCompleted struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeRecoveryCompleted is a free log retrieval operation binding the contract event 0x60aac8c36efdaabf125dc9ec2124bde8b3ceafe5c8b4fc8063fc4ac9017eb0be.
//
// Solidity: event NodeRecoveryCompleted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeRecoveryCompleted(opts *bind.FilterOpts) (*NodeManagerAbiNodeRecoveryCompletedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeRecoveryCompleted")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeRecoveryCompletedIterator{contract: _NodeManagerAbi.contract, event: "NodeRecoveryCompleted", logs: logs, sub: sub}, nil
}

// WatchNodeRecoveryCompleted is a free log subscription operation binding the contract event 0x60aac8c36efdaabf125dc9ec2124bde8b3ceafe5c8b4fc8063fc4ac9017eb0be.
//
// Solidity: event NodeRecoveryCompleted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeRecoveryCompleted(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeRecoveryCompleted) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeRecoveryCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeRecoveryCompleted)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeRecoveryCompleted", log); err != nil {
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

// ParseNodeRecoveryCompleted is a log parse operation binding the contract event 0x60aac8c36efdaabf125dc9ec2124bde8b3ceafe5c8b4fc8063fc4ac9017eb0be.
//
// Solidity: event NodeRecoveryCompleted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeRecoveryCompleted(log types.Log) (*NodeManagerAbiNodeRecoveryCompleted, error) {
	event := new(NodeManagerAbiNodeRecoveryCompleted)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeRecoveryCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerAbiNodeRecoveryInitiatedIterator is returned from FilterNodeRecoveryInitiated and is used to iterate over the raw logs and unpacked data for NodeRecoveryInitiated events raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeRecoveryInitiatedIterator struct {
	Event *NodeManagerAbiNodeRecoveryInitiated // Event containing the contract specifics and raw log

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
func (it *NodeManagerAbiNodeRecoveryInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerAbiNodeRecoveryInitiated)
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
		it.Event = new(NodeManagerAbiNodeRecoveryInitiated)
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
func (it *NodeManagerAbiNodeRecoveryInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerAbiNodeRecoveryInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerAbiNodeRecoveryInitiated represents a NodeRecoveryInitiated event raised by the NodeManagerAbi contract.
type NodeManagerAbiNodeRecoveryInitiated struct {
	EnodeId  string
	Ip       string
	Port     uint16
	Raftport uint16
	OrgId    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeRecoveryInitiated is a free log retrieval operation binding the contract event 0x72779f66ea90e28bae76fbfe03eaef5ae01699976c7493f93186ab9560ccfaa4.
//
// Solidity: event NodeRecoveryInitiated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) FilterNodeRecoveryInitiated(opts *bind.FilterOpts) (*NodeManagerAbiNodeRecoveryInitiatedIterator, error) {

	logs, sub, err := _NodeManagerAbi.contract.FilterLogs(opts, "NodeRecoveryInitiated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerAbiNodeRecoveryInitiatedIterator{contract: _NodeManagerAbi.contract, event: "NodeRecoveryInitiated", logs: logs, sub: sub}, nil
}

// WatchNodeRecoveryInitiated is a free log subscription operation binding the contract event 0x72779f66ea90e28bae76fbfe03eaef5ae01699976c7493f93186ab9560ccfaa4.
//
// Solidity: event NodeRecoveryInitiated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) WatchNodeRecoveryInitiated(opts *bind.WatchOpts, sink chan<- *NodeManagerAbiNodeRecoveryInitiated) (event.Subscription, error) {

	logs, sub, err := _NodeManagerAbi.contract.WatchLogs(opts, "NodeRecoveryInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerAbiNodeRecoveryInitiated)
				if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeRecoveryInitiated", log); err != nil {
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

// ParseNodeRecoveryInitiated is a log parse operation binding the contract event 0x72779f66ea90e28bae76fbfe03eaef5ae01699976c7493f93186ab9560ccfaa4.
//
// Solidity: event NodeRecoveryInitiated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId)
func (_NodeManagerAbi *NodeManagerAbiFilterer) ParseNodeRecoveryInitiated(log types.Log) (*NodeManagerAbiNodeRecoveryInitiated, error) {
	event := new(NodeManagerAbiNodeRecoveryInitiated)
	if err := _NodeManagerAbi.contract.UnpackLog(event, "NodeRecoveryInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
