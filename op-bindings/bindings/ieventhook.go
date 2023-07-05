// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IEventHookMetaData contains all meta data concerning the IEventHook contract.
var IEventHookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"topics\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEventHookABI is the input ABI used to generate the binding from.
// Deprecated: Use IEventHookMetaData.ABI instead.
var IEventHookABI = IEventHookMetaData.ABI

// IEventHook is an auto generated Go binding around an Ethereum contract.
type IEventHook struct {
	IEventHookCaller     // Read-only binding to the contract
	IEventHookTransactor // Write-only binding to the contract
	IEventHookFilterer   // Log filterer for contract events
}

// IEventHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEventHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEventHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEventHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEventHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEventHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEventHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEventHookSession struct {
	Contract     *IEventHook       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEventHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEventHookCallerSession struct {
	Contract *IEventHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IEventHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEventHookTransactorSession struct {
	Contract     *IEventHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IEventHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEventHookRaw struct {
	Contract *IEventHook // Generic contract binding to access the raw methods on
}

// IEventHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEventHookCallerRaw struct {
	Contract *IEventHookCaller // Generic read-only contract binding to access the raw methods on
}

// IEventHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEventHookTransactorRaw struct {
	Contract *IEventHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEventHook creates a new instance of IEventHook, bound to a specific deployed contract.
func NewIEventHook(address common.Address, backend bind.ContractBackend) (*IEventHook, error) {
	contract, err := bindIEventHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEventHook{IEventHookCaller: IEventHookCaller{contract: contract}, IEventHookTransactor: IEventHookTransactor{contract: contract}, IEventHookFilterer: IEventHookFilterer{contract: contract}}, nil
}

// NewIEventHookCaller creates a new read-only instance of IEventHook, bound to a specific deployed contract.
func NewIEventHookCaller(address common.Address, caller bind.ContractCaller) (*IEventHookCaller, error) {
	contract, err := bindIEventHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEventHookCaller{contract: contract}, nil
}

// NewIEventHookTransactor creates a new write-only instance of IEventHook, bound to a specific deployed contract.
func NewIEventHookTransactor(address common.Address, transactor bind.ContractTransactor) (*IEventHookTransactor, error) {
	contract, err := bindIEventHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEventHookTransactor{contract: contract}, nil
}

// NewIEventHookFilterer creates a new log filterer instance of IEventHook, bound to a specific deployed contract.
func NewIEventHookFilterer(address common.Address, filterer bind.ContractFilterer) (*IEventHookFilterer, error) {
	contract, err := bindIEventHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEventHookFilterer{contract: contract}, nil
}

// bindIEventHook binds a generic wrapper to an already deployed contract.
func bindIEventHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEventHookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEventHook *IEventHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEventHook.Contract.IEventHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEventHook *IEventHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEventHook.Contract.IEventHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEventHook *IEventHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEventHook.Contract.IEventHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEventHook *IEventHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEventHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEventHook *IEventHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEventHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEventHook *IEventHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEventHook.Contract.contract.Transact(opts, method, params...)
}

// Handle is a paid mutator transaction binding the contract method 0x87aeac00.
//
// Solidity: function handle(address origin, bytes topics, bytes data) returns()
func (_IEventHook *IEventHookTransactor) Handle(opts *bind.TransactOpts, origin common.Address, topics []byte, data []byte) (*types.Transaction, error) {
	return _IEventHook.contract.Transact(opts, "handle", origin, topics, data)
}

// Handle is a paid mutator transaction binding the contract method 0x87aeac00.
//
// Solidity: function handle(address origin, bytes topics, bytes data) returns()
func (_IEventHook *IEventHookSession) Handle(origin common.Address, topics []byte, data []byte) (*types.Transaction, error) {
	return _IEventHook.Contract.Handle(&_IEventHook.TransactOpts, origin, topics, data)
}

// Handle is a paid mutator transaction binding the contract method 0x87aeac00.
//
// Solidity: function handle(address origin, bytes topics, bytes data) returns()
func (_IEventHook *IEventHookTransactorSession) Handle(origin common.Address, topics []byte, data []byte) (*types.Transaction, error) {
	return _IEventHook.Contract.Handle(&_IEventHook.TransactOpts, origin, topics, data)
}
