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

// EventHookRegistryEventHookItem is an auto generated low-level Go binding around an user-defined struct.
type EventHookRegistryEventHookItem struct {
	Topic    [32]byte
	Origin   common.Address
	Receiver common.Address
}

// EventHookRegistryMetaData contains all meta data concerning the EventHookRegistry contract.
var EventHookRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEPOSITOR_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"addEventHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEventHooks\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structEventHookRegistry.EventHookItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goodbye\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hooks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600560809081526468656c6c6f60d81b60a05260019061002490826100d6565b5034801561003157600080fd5b50610195565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061006157607f821691505b60208210810361008157634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156100d157600081815260208120601f850160051c810160208610156100ae5750805b601f850160051c820191505b818110156100cd578281556001016100ba565b5050505b505050565b81516001600160401b038111156100ef576100ef610037565b610103816100fd845461004d565b84610087565b602080601f83116001811461013857600084156101205750858301515b600019600386901b1c1916600185901b1785556100cd565b600085815260208120601f198616915b8281101561016757888601518255948401946001909101908401610148565b50858210156101855787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610700806101a46000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063864c17d711610050578063864c17d7146101d8578063dc85b4991461021d578063e591b2821461023257600080fd5b806319ff1d211461007757806375fc8e3c146100955780637f5c4868146100ce575b600080fd5b61007f610272565b60405161008c9190610485565b60405180910390f35b60408051808201909152600781527f676f6f6462796500000000000000000000000000000000000000000000000000602082015261007f565b6101d66100dc366004610521565b6040805160608101825293845273ffffffffffffffffffffffffffffffffffffffff9283166020850190815291831690840190815260028054600181018255600091909152935160039094027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81019490945590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf840180549184167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617905590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad09093018054939092169216919091179055565b005b6101eb6101e636600461055d565b610300565b6040805193845273ffffffffffffffffffffffffffffffffffffffff928316602085015291169082015260600161008c565b61022561034d565b60405161008c9190610576565b61024d73deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008c565b6001805461027f906105e9565b80601f01602080910402602001604051908101604052809291908181526020018280546102ab906105e9565b80156102f85780601f106102cd576101008083540402835291602001916102f8565b820191906000526020600020905b8154815290600101906020018083116102db57829003601f168201915b505050505081565b6002818154811061031057600080fd5b600091825260209091206003909102018054600182015460029092015490925073ffffffffffffffffffffffffffffffffffffffff918216911683565b60025460609060009067ffffffffffffffff81111561036e5761036e610636565b6040519080825280602002602001820160405280156103d757816020015b60408051606081018252600080825260208083018290529282015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90920191018161038c5790505b50905060005b60025481101561047f57600281815481106103fa576103fa610665565b600091825260209182902060408051606081018252600390930290910180548352600181015473ffffffffffffffffffffffffffffffffffffffff908116948401949094526002015490921691810191909152825183908390811061046157610461610665565b6020026020010181905250808061047790610694565b9150506103dd565b50919050565b600060208083528351808285015260005b818110156104b257858101830151858201604001528201610496565b818111156104c4576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461051c57600080fd5b919050565b60008060006060848603121561053657600080fd5b83359250610546602085016104f8565b9150610554604085016104f8565b90509250925092565b60006020828403121561056f57600080fd5b5035919050565b602080825282518282018190526000919060409081850190868401855b828110156105dc578151805185528681015173ffffffffffffffffffffffffffffffffffffffff9081168887015290860151168585015260609093019290850190600101610593565b5091979650505050505050565b600181811c908216806105fd57607f821691505b60208210810361047f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036106ec577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

// EventHookRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EventHookRegistryMetaData.ABI instead.
var EventHookRegistryABI = EventHookRegistryMetaData.ABI

// EventHookRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventHookRegistryMetaData.Bin instead.
var EventHookRegistryBin = EventHookRegistryMetaData.Bin

// DeployEventHookRegistry deploys a new Ethereum contract, binding an instance of EventHookRegistry to it.
func DeployEventHookRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EventHookRegistry, error) {
	parsed, err := EventHookRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventHookRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventHookRegistry{EventHookRegistryCaller: EventHookRegistryCaller{contract: contract}, EventHookRegistryTransactor: EventHookRegistryTransactor{contract: contract}, EventHookRegistryFilterer: EventHookRegistryFilterer{contract: contract}}, nil
}

// EventHookRegistry is an auto generated Go binding around an Ethereum contract.
type EventHookRegistry struct {
	EventHookRegistryCaller     // Read-only binding to the contract
	EventHookRegistryTransactor // Write-only binding to the contract
	EventHookRegistryFilterer   // Log filterer for contract events
}

// EventHookRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventHookRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventHookRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventHookRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventHookRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventHookRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventHookRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventHookRegistrySession struct {
	Contract     *EventHookRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EventHookRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventHookRegistryCallerSession struct {
	Contract *EventHookRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EventHookRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventHookRegistryTransactorSession struct {
	Contract     *EventHookRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EventHookRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventHookRegistryRaw struct {
	Contract *EventHookRegistry // Generic contract binding to access the raw methods on
}

// EventHookRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventHookRegistryCallerRaw struct {
	Contract *EventHookRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EventHookRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventHookRegistryTransactorRaw struct {
	Contract *EventHookRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventHookRegistry creates a new instance of EventHookRegistry, bound to a specific deployed contract.
func NewEventHookRegistry(address common.Address, backend bind.ContractBackend) (*EventHookRegistry, error) {
	contract, err := bindEventHookRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistry{EventHookRegistryCaller: EventHookRegistryCaller{contract: contract}, EventHookRegistryTransactor: EventHookRegistryTransactor{contract: contract}, EventHookRegistryFilterer: EventHookRegistryFilterer{contract: contract}}, nil
}

// NewEventHookRegistryCaller creates a new read-only instance of EventHookRegistry, bound to a specific deployed contract.
func NewEventHookRegistryCaller(address common.Address, caller bind.ContractCaller) (*EventHookRegistryCaller, error) {
	contract, err := bindEventHookRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistryCaller{contract: contract}, nil
}

// NewEventHookRegistryTransactor creates a new write-only instance of EventHookRegistry, bound to a specific deployed contract.
func NewEventHookRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EventHookRegistryTransactor, error) {
	contract, err := bindEventHookRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistryTransactor{contract: contract}, nil
}

// NewEventHookRegistryFilterer creates a new log filterer instance of EventHookRegistry, bound to a specific deployed contract.
func NewEventHookRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EventHookRegistryFilterer, error) {
	contract, err := bindEventHookRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistryFilterer{contract: contract}, nil
}

// bindEventHookRegistry binds a generic wrapper to an already deployed contract.
func bindEventHookRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventHookRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventHookRegistry *EventHookRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventHookRegistry.Contract.EventHookRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventHookRegistry *EventHookRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.EventHookRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventHookRegistry *EventHookRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.EventHookRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventHookRegistry *EventHookRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventHookRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventHookRegistry *EventHookRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventHookRegistry *EventHookRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_EventHookRegistry *EventHookRegistryCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventHookRegistry.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_EventHookRegistry *EventHookRegistrySession) DEPOSITORACCOUNT() (common.Address, error) {
	return _EventHookRegistry.Contract.DEPOSITORACCOUNT(&_EventHookRegistry.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_EventHookRegistry *EventHookRegistryCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _EventHookRegistry.Contract.DEPOSITORACCOUNT(&_EventHookRegistry.CallOpts)
}

// GetEventHooks is a free data retrieval call binding the contract method 0xdc85b499.
//
// Solidity: function getEventHooks() view returns((bytes32,address,address)[])
func (_EventHookRegistry *EventHookRegistryCaller) GetEventHooks(opts *bind.CallOpts) ([]EventHookRegistryEventHookItem, error) {
	var out []interface{}
	err := _EventHookRegistry.contract.Call(opts, &out, "getEventHooks")

	if err != nil {
		return *new([]EventHookRegistryEventHookItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventHookRegistryEventHookItem)).(*[]EventHookRegistryEventHookItem)

	return out0, err

}

// GetEventHooks is a free data retrieval call binding the contract method 0xdc85b499.
//
// Solidity: function getEventHooks() view returns((bytes32,address,address)[])
func (_EventHookRegistry *EventHookRegistrySession) GetEventHooks() ([]EventHookRegistryEventHookItem, error) {
	return _EventHookRegistry.Contract.GetEventHooks(&_EventHookRegistry.CallOpts)
}

// GetEventHooks is a free data retrieval call binding the contract method 0xdc85b499.
//
// Solidity: function getEventHooks() view returns((bytes32,address,address)[])
func (_EventHookRegistry *EventHookRegistryCallerSession) GetEventHooks() ([]EventHookRegistryEventHookItem, error) {
	return _EventHookRegistry.Contract.GetEventHooks(&_EventHookRegistry.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() view returns(string)
func (_EventHookRegistry *EventHookRegistryCaller) Hello(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventHookRegistry.contract.Call(opts, &out, "hello")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() view returns(string)
func (_EventHookRegistry *EventHookRegistrySession) Hello() (string, error) {
	return _EventHookRegistry.Contract.Hello(&_EventHookRegistry.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() view returns(string)
func (_EventHookRegistry *EventHookRegistryCallerSession) Hello() (string, error) {
	return _EventHookRegistry.Contract.Hello(&_EventHookRegistry.CallOpts)
}

// Hooks is a free data retrieval call binding the contract method 0x864c17d7.
//
// Solidity: function hooks(uint256 ) view returns(bytes32 topic, address origin, address receiver)
func (_EventHookRegistry *EventHookRegistryCaller) Hooks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Topic    [32]byte
	Origin   common.Address
	Receiver common.Address
}, error) {
	var out []interface{}
	err := _EventHookRegistry.contract.Call(opts, &out, "hooks", arg0)

	outstruct := new(struct {
		Topic    [32]byte
		Origin   common.Address
		Receiver common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Topic = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Origin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Hooks is a free data retrieval call binding the contract method 0x864c17d7.
//
// Solidity: function hooks(uint256 ) view returns(bytes32 topic, address origin, address receiver)
func (_EventHookRegistry *EventHookRegistrySession) Hooks(arg0 *big.Int) (struct {
	Topic    [32]byte
	Origin   common.Address
	Receiver common.Address
}, error) {
	return _EventHookRegistry.Contract.Hooks(&_EventHookRegistry.CallOpts, arg0)
}

// Hooks is a free data retrieval call binding the contract method 0x864c17d7.
//
// Solidity: function hooks(uint256 ) view returns(bytes32 topic, address origin, address receiver)
func (_EventHookRegistry *EventHookRegistryCallerSession) Hooks(arg0 *big.Int) (struct {
	Topic    [32]byte
	Origin   common.Address
	Receiver common.Address
}, error) {
	return _EventHookRegistry.Contract.Hooks(&_EventHookRegistry.CallOpts, arg0)
}

// AddEventHook is a paid mutator transaction binding the contract method 0x7f5c4868.
//
// Solidity: function addEventHook(bytes32 _topic, address _origin, address _receiver) returns()
func (_EventHookRegistry *EventHookRegistryTransactor) AddEventHook(opts *bind.TransactOpts, _topic [32]byte, _origin common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.contract.Transact(opts, "addEventHook", _topic, _origin, _receiver)
}

// AddEventHook is a paid mutator transaction binding the contract method 0x7f5c4868.
//
// Solidity: function addEventHook(bytes32 _topic, address _origin, address _receiver) returns()
func (_EventHookRegistry *EventHookRegistrySession) AddEventHook(_topic [32]byte, _origin common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.AddEventHook(&_EventHookRegistry.TransactOpts, _topic, _origin, _receiver)
}

// AddEventHook is a paid mutator transaction binding the contract method 0x7f5c4868.
//
// Solidity: function addEventHook(bytes32 _topic, address _origin, address _receiver) returns()
func (_EventHookRegistry *EventHookRegistryTransactorSession) AddEventHook(_topic [32]byte, _origin common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.AddEventHook(&_EventHookRegistry.TransactOpts, _topic, _origin, _receiver)
}

// Goodbye is a paid mutator transaction binding the contract method 0x75fc8e3c.
//
// Solidity: function goodbye() returns(string)
func (_EventHookRegistry *EventHookRegistryTransactor) Goodbye(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventHookRegistry.contract.Transact(opts, "goodbye")
}

// Goodbye is a paid mutator transaction binding the contract method 0x75fc8e3c.
//
// Solidity: function goodbye() returns(string)
func (_EventHookRegistry *EventHookRegistrySession) Goodbye() (*types.Transaction, error) {
	return _EventHookRegistry.Contract.Goodbye(&_EventHookRegistry.TransactOpts)
}

// Goodbye is a paid mutator transaction binding the contract method 0x75fc8e3c.
//
// Solidity: function goodbye() returns(string)
func (_EventHookRegistry *EventHookRegistryTransactorSession) Goodbye() (*types.Transaction, error) {
	return _EventHookRegistry.Contract.Goodbye(&_EventHookRegistry.TransactOpts)
}
