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

// EventHookItem is an auto generated low-level Go binding around an user-defined struct.
type EventHookItem struct {
	Topic    [32]byte
	Origin   common.Address
	Receiver common.Address
}

// EventHookRegistryMetaData contains all meta data concerning the EventHookRegistry contract.
var EventHookRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"EventHookAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"EventHookRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"addEventHook\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEventHooks\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structEventHookItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"removeEventHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610ba9380380610ba983398101604081905261002f91610097565b61003833610047565b61004181610047565b506100c7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100a957600080fd5b81516001600160a01b03811681146100c057600080fd5b9392505050565b610ad3806100d66000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80638da5cb5b116100505780638da5cb5b146100ba578063dc85b499146100e2578063f2fde38b146100f757600080fd5b8063715018a6146100775780637bae7b72146100815780637f5c486814610094575b600080fd5b61007f61010a565b005b61007f61008f3660046108af565b61011e565b6100a76100a23660046108f1565b61021f565b6040519081526020015b60405180910390f35b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b1565b6100ea61040e565b6040516100b1919061092d565b61007f6101053660046109a0565b61054b565b610112610602565b61011c6000610683565b565b610126610602565b6101316002826106f8565b61019c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f21686f6f6b00000000000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b60008181526001602081905260408220918255810180547fffffffffffffffffffffffff00000000000000000000000000000000000000009081169091556002918201805490911690556101f09082610715565b5060405181907f06ec7e790ccfc5856f4aa76b96fe94c880e0da3848b7160500b8aaf346fd77f190600090a250565b6000610229610602565b60008484846040516020016102679392919092835273ffffffffffffffffffffffffffffffffffffffff918216602084015216604082015260600190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012090506102aa6002826106f8565b15610311576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f686f6f6b206578697374730000000000000000000000000000000000000000006044820152606401610193565b6040805160608101825286815273ffffffffffffffffffffffffffffffffffffffff808716602080840191825287831684860190815260008781526001928390529590952084518155915190820180549184167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055935160029182018054919093169416939093179055906103ab9083610721565b506040805187815273ffffffffffffffffffffffffffffffffffffffff8781166020830152861681830152905183917fdf6bdf83d24c78bb26eb3d7b34123cc0be2370aa49944d3ca1b9b8801becd41c919081900360600190a250949350505050565b6060600061041c600261072d565b67ffffffffffffffff811115610434576104346109bb565b60405190808252806020026020018201604052801561049d57816020015b60408051606081018252600080825260208083018290529282015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816104525790505b50905060005b6104ad600261072d565b8110156105455760006104c1600283610737565b6000818152600160208181526040928390208351606081018552815481529281015473ffffffffffffffffffffffffffffffffffffffff90811692840192909252600201541691810191909152845191925090849084908110610526576105266109ea565b602002602001018190525050808061053d90610a48565b9150506104a3565b50919050565b610553610602565b73ffffffffffffffffffffffffffffffffffffffff81166105f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610193565b6105ff81610683565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461011c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610193565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600081815260018301602052604081205415155b90505b92915050565b600061070c8383610743565b600061070c8383610836565b600061070f825490565b600061070c8383610885565b6000818152600183016020526040812054801561082c576000610767600183610a80565b855490915060009061077b90600190610a80565b90508181146107e057600086600001828154811061079b5761079b6109ea565b90600052602060002001549050808760000184815481106107be576107be6109ea565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806107f1576107f1610a97565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061070f565b600091505061070f565b600081815260018301602052604081205461087d5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561070f565b50600061070f565b600082600001828154811061089c5761089c6109ea565b9060005260206000200154905092915050565b6000602082840312156108c157600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146108ec57600080fd5b919050565b60008060006060848603121561090657600080fd5b83359250610916602085016108c8565b9150610924604085016108c8565b90509250925092565b602080825282518282018190526000919060409081850190868401855b82811015610993578151805185528681015173ffffffffffffffffffffffffffffffffffffffff908116888701529086015116858501526060909301929085019060010161094a565b5091979650505050505050565b6000602082840312156109b257600080fd5b61070c826108c8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610a7957610a79610a19565b5060010190565b600082821015610a9257610a92610a19565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000a",
}

// EventHookRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EventHookRegistryMetaData.ABI instead.
var EventHookRegistryABI = EventHookRegistryMetaData.ABI

// EventHookRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventHookRegistryMetaData.Bin instead.
var EventHookRegistryBin = EventHookRegistryMetaData.Bin

// DeployEventHookRegistry deploys a new Ethereum contract, binding an instance of EventHookRegistry to it.
func DeployEventHookRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *EventHookRegistry, error) {
	parsed, err := EventHookRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventHookRegistryBin), backend, _owner)
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

// GetEventHooks is a free data retrieval call binding the contract method 0xdc85b499.
//
// Solidity: function getEventHooks() view returns((bytes32,address,address)[])
func (_EventHookRegistry *EventHookRegistryCaller) GetEventHooks(opts *bind.CallOpts) ([]EventHookItem, error) {
	var out []interface{}
	err := _EventHookRegistry.contract.Call(opts, &out, "getEventHooks")

	if err != nil {
		return *new([]EventHookItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventHookItem)).(*[]EventHookItem)

	return out0, err

}

// GetEventHooks is a free data retrieval call binding the contract method 0xdc85b499.
//
// Solidity: function getEventHooks() view returns((bytes32,address,address)[])
func (_EventHookRegistry *EventHookRegistrySession) GetEventHooks() ([]EventHookItem, error) {
	return _EventHookRegistry.Contract.GetEventHooks(&_EventHookRegistry.CallOpts)
}

// GetEventHooks is a free data retrieval call binding the contract method 0xdc85b499.
//
// Solidity: function getEventHooks() view returns((bytes32,address,address)[])
func (_EventHookRegistry *EventHookRegistryCallerSession) GetEventHooks() ([]EventHookItem, error) {
	return _EventHookRegistry.Contract.GetEventHooks(&_EventHookRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventHookRegistry *EventHookRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventHookRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventHookRegistry *EventHookRegistrySession) Owner() (common.Address, error) {
	return _EventHookRegistry.Contract.Owner(&_EventHookRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventHookRegistry *EventHookRegistryCallerSession) Owner() (common.Address, error) {
	return _EventHookRegistry.Contract.Owner(&_EventHookRegistry.CallOpts)
}

// AddEventHook is a paid mutator transaction binding the contract method 0x7f5c4868.
//
// Solidity: function addEventHook(bytes32 _topic, address _origin, address _receiver) returns(bytes32)
func (_EventHookRegistry *EventHookRegistryTransactor) AddEventHook(opts *bind.TransactOpts, _topic [32]byte, _origin common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.contract.Transact(opts, "addEventHook", _topic, _origin, _receiver)
}

// AddEventHook is a paid mutator transaction binding the contract method 0x7f5c4868.
//
// Solidity: function addEventHook(bytes32 _topic, address _origin, address _receiver) returns(bytes32)
func (_EventHookRegistry *EventHookRegistrySession) AddEventHook(_topic [32]byte, _origin common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.AddEventHook(&_EventHookRegistry.TransactOpts, _topic, _origin, _receiver)
}

// AddEventHook is a paid mutator transaction binding the contract method 0x7f5c4868.
//
// Solidity: function addEventHook(bytes32 _topic, address _origin, address _receiver) returns(bytes32)
func (_EventHookRegistry *EventHookRegistryTransactorSession) AddEventHook(_topic [32]byte, _origin common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.AddEventHook(&_EventHookRegistry.TransactOpts, _topic, _origin, _receiver)
}

// RemoveEventHook is a paid mutator transaction binding the contract method 0x7bae7b72.
//
// Solidity: function removeEventHook(bytes32 id) returns()
func (_EventHookRegistry *EventHookRegistryTransactor) RemoveEventHook(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _EventHookRegistry.contract.Transact(opts, "removeEventHook", id)
}

// RemoveEventHook is a paid mutator transaction binding the contract method 0x7bae7b72.
//
// Solidity: function removeEventHook(bytes32 id) returns()
func (_EventHookRegistry *EventHookRegistrySession) RemoveEventHook(id [32]byte) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.RemoveEventHook(&_EventHookRegistry.TransactOpts, id)
}

// RemoveEventHook is a paid mutator transaction binding the contract method 0x7bae7b72.
//
// Solidity: function removeEventHook(bytes32 id) returns()
func (_EventHookRegistry *EventHookRegistryTransactorSession) RemoveEventHook(id [32]byte) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.RemoveEventHook(&_EventHookRegistry.TransactOpts, id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EventHookRegistry *EventHookRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventHookRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EventHookRegistry *EventHookRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _EventHookRegistry.Contract.RenounceOwnership(&_EventHookRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EventHookRegistry *EventHookRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EventHookRegistry.Contract.RenounceOwnership(&_EventHookRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EventHookRegistry *EventHookRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EventHookRegistry *EventHookRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.TransferOwnership(&_EventHookRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EventHookRegistry *EventHookRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EventHookRegistry.Contract.TransferOwnership(&_EventHookRegistry.TransactOpts, newOwner)
}

// EventHookRegistryEventHookAddedIterator is returned from FilterEventHookAdded and is used to iterate over the raw logs and unpacked data for EventHookAdded events raised by the EventHookRegistry contract.
type EventHookRegistryEventHookAddedIterator struct {
	Event *EventHookRegistryEventHookAdded // Event containing the contract specifics and raw log

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
func (it *EventHookRegistryEventHookAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventHookRegistryEventHookAdded)
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
		it.Event = new(EventHookRegistryEventHookAdded)
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
func (it *EventHookRegistryEventHookAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventHookRegistryEventHookAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventHookRegistryEventHookAdded represents a EventHookAdded event raised by the EventHookRegistry contract.
type EventHookRegistryEventHookAdded struct {
	Id       [32]byte
	Topic    [32]byte
	Origin   common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventHookAdded is a free log retrieval operation binding the contract event 0xdf6bdf83d24c78bb26eb3d7b34123cc0be2370aa49944d3ca1b9b8801becd41c.
//
// Solidity: event EventHookAdded(bytes32 indexed id, bytes32 topic, address origin, address receiver)
func (_EventHookRegistry *EventHookRegistryFilterer) FilterEventHookAdded(opts *bind.FilterOpts, id [][32]byte) (*EventHookRegistryEventHookAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventHookRegistry.contract.FilterLogs(opts, "EventHookAdded", idRule)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistryEventHookAddedIterator{contract: _EventHookRegistry.contract, event: "EventHookAdded", logs: logs, sub: sub}, nil
}

// WatchEventHookAdded is a free log subscription operation binding the contract event 0xdf6bdf83d24c78bb26eb3d7b34123cc0be2370aa49944d3ca1b9b8801becd41c.
//
// Solidity: event EventHookAdded(bytes32 indexed id, bytes32 topic, address origin, address receiver)
func (_EventHookRegistry *EventHookRegistryFilterer) WatchEventHookAdded(opts *bind.WatchOpts, sink chan<- *EventHookRegistryEventHookAdded, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventHookRegistry.contract.WatchLogs(opts, "EventHookAdded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventHookRegistryEventHookAdded)
				if err := _EventHookRegistry.contract.UnpackLog(event, "EventHookAdded", log); err != nil {
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

// ParseEventHookAdded is a log parse operation binding the contract event 0xdf6bdf83d24c78bb26eb3d7b34123cc0be2370aa49944d3ca1b9b8801becd41c.
//
// Solidity: event EventHookAdded(bytes32 indexed id, bytes32 topic, address origin, address receiver)
func (_EventHookRegistry *EventHookRegistryFilterer) ParseEventHookAdded(log types.Log) (*EventHookRegistryEventHookAdded, error) {
	event := new(EventHookRegistryEventHookAdded)
	if err := _EventHookRegistry.contract.UnpackLog(event, "EventHookAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventHookRegistryEventHookRemovedIterator is returned from FilterEventHookRemoved and is used to iterate over the raw logs and unpacked data for EventHookRemoved events raised by the EventHookRegistry contract.
type EventHookRegistryEventHookRemovedIterator struct {
	Event *EventHookRegistryEventHookRemoved // Event containing the contract specifics and raw log

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
func (it *EventHookRegistryEventHookRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventHookRegistryEventHookRemoved)
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
		it.Event = new(EventHookRegistryEventHookRemoved)
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
func (it *EventHookRegistryEventHookRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventHookRegistryEventHookRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventHookRegistryEventHookRemoved represents a EventHookRemoved event raised by the EventHookRegistry contract.
type EventHookRegistryEventHookRemoved struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEventHookRemoved is a free log retrieval operation binding the contract event 0x06ec7e790ccfc5856f4aa76b96fe94c880e0da3848b7160500b8aaf346fd77f1.
//
// Solidity: event EventHookRemoved(bytes32 indexed id)
func (_EventHookRegistry *EventHookRegistryFilterer) FilterEventHookRemoved(opts *bind.FilterOpts, id [][32]byte) (*EventHookRegistryEventHookRemovedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventHookRegistry.contract.FilterLogs(opts, "EventHookRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistryEventHookRemovedIterator{contract: _EventHookRegistry.contract, event: "EventHookRemoved", logs: logs, sub: sub}, nil
}

// WatchEventHookRemoved is a free log subscription operation binding the contract event 0x06ec7e790ccfc5856f4aa76b96fe94c880e0da3848b7160500b8aaf346fd77f1.
//
// Solidity: event EventHookRemoved(bytes32 indexed id)
func (_EventHookRegistry *EventHookRegistryFilterer) WatchEventHookRemoved(opts *bind.WatchOpts, sink chan<- *EventHookRegistryEventHookRemoved, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventHookRegistry.contract.WatchLogs(opts, "EventHookRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventHookRegistryEventHookRemoved)
				if err := _EventHookRegistry.contract.UnpackLog(event, "EventHookRemoved", log); err != nil {
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

// ParseEventHookRemoved is a log parse operation binding the contract event 0x06ec7e790ccfc5856f4aa76b96fe94c880e0da3848b7160500b8aaf346fd77f1.
//
// Solidity: event EventHookRemoved(bytes32 indexed id)
func (_EventHookRegistry *EventHookRegistryFilterer) ParseEventHookRemoved(log types.Log) (*EventHookRegistryEventHookRemoved, error) {
	event := new(EventHookRegistryEventHookRemoved)
	if err := _EventHookRegistry.contract.UnpackLog(event, "EventHookRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventHookRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EventHookRegistry contract.
type EventHookRegistryOwnershipTransferredIterator struct {
	Event *EventHookRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EventHookRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventHookRegistryOwnershipTransferred)
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
		it.Event = new(EventHookRegistryOwnershipTransferred)
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
func (it *EventHookRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventHookRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventHookRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the EventHookRegistry contract.
type EventHookRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EventHookRegistry *EventHookRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EventHookRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EventHookRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EventHookRegistryOwnershipTransferredIterator{contract: _EventHookRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EventHookRegistry *EventHookRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EventHookRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EventHookRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventHookRegistryOwnershipTransferred)
				if err := _EventHookRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EventHookRegistry *EventHookRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*EventHookRegistryOwnershipTransferred, error) {
	event := new(EventHookRegistryOwnershipTransferred)
	if err := _EventHookRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
