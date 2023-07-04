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

// BurnMetaData contains all meta data concerning the Burn contract.
var BurnMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x602d6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea164736f6c634300080f000a",
}

// BurnABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnMetaData.ABI instead.
var BurnABI = BurnMetaData.ABI

// BurnBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BurnMetaData.Bin instead.
var BurnBin = BurnMetaData.Bin

// DeployBurn deploys a new Ethereum contract, binding an instance of Burn to it.
func DeployBurn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Burn, error) {
	parsed, err := BurnMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Burn{BurnCaller: BurnCaller{contract: contract}, BurnTransactor: BurnTransactor{contract: contract}, BurnFilterer: BurnFilterer{contract: contract}}, nil
}

// Burn is an auto generated Go binding around an Ethereum contract.
type Burn struct {
	BurnCaller     // Read-only binding to the contract
	BurnTransactor // Write-only binding to the contract
	BurnFilterer   // Log filterer for contract events
}

// BurnCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnSession struct {
	Contract     *Burn             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnCallerSession struct {
	Contract *BurnCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BurnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnTransactorSession struct {
	Contract     *BurnTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnRaw struct {
	Contract *Burn // Generic contract binding to access the raw methods on
}

// BurnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnCallerRaw struct {
	Contract *BurnCaller // Generic read-only contract binding to access the raw methods on
}

// BurnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnTransactorRaw struct {
	Contract *BurnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurn creates a new instance of Burn, bound to a specific deployed contract.
func NewBurn(address common.Address, backend bind.ContractBackend) (*Burn, error) {
	contract, err := bindBurn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Burn{BurnCaller: BurnCaller{contract: contract}, BurnTransactor: BurnTransactor{contract: contract}, BurnFilterer: BurnFilterer{contract: contract}}, nil
}

// NewBurnCaller creates a new read-only instance of Burn, bound to a specific deployed contract.
func NewBurnCaller(address common.Address, caller bind.ContractCaller) (*BurnCaller, error) {
	contract, err := bindBurn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnCaller{contract: contract}, nil
}

// NewBurnTransactor creates a new write-only instance of Burn, bound to a specific deployed contract.
func NewBurnTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnTransactor, error) {
	contract, err := bindBurn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnTransactor{contract: contract}, nil
}

// NewBurnFilterer creates a new log filterer instance of Burn, bound to a specific deployed contract.
func NewBurnFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnFilterer, error) {
	contract, err := bindBurn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnFilterer{contract: contract}, nil
}

// bindBurn binds a generic wrapper to an already deployed contract.
func bindBurn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burn *BurnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Burn.Contract.BurnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burn *BurnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burn.Contract.BurnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burn *BurnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burn.Contract.BurnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burn *BurnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Burn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burn *BurnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burn *BurnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burn.Contract.contract.Transact(opts, method, params...)
}
