// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleStorage

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimpleStorageABI is the input ABI used to generate the binding from.
const SimpleStorageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldData\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newData\",\"type\":\"uint256\"}],\"name\":\"DataSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"x\",\"type\":\"uint8\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleStorageBin is the compiled bytecode used for deploying new contracts.
var SimpleStorageBin = "0x608060405234801561001057600080fd5b5061012e806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806324b8ba5f1460375780636d4ce63c146065575b600080fd5b606360048036036020811015604b57600080fd5b81019080803560ff1690602001909291905050506084565b005b606b60e2565b604051808260ff16815260200191505060405180910390f35b8060ff1660008054906101000a900460ff1660ff167fc911a63e29945a04493ec58a89a96aa00a33c2609f1c96d4e506a7fb094e4c9460405160405180910390a3806000806101000a81548160ff021916908360ff16021790555050565b60008060009054906101000a900460ff1690509056fea2646970667358221220c9d8f03031269e6ec69a818d9ce0c3f81987ffeb407421297f538d98c925fe8064736f6c63430007050033"

// DeploySimpleStorage deploys a new Ethereum contract, binding an instance of SimpleStorage to it.
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleStorageCallerSession struct {
	Contract *SimpleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleStorageTransactorSession struct {
	Contract     *SimpleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleStorageRaw struct {
	Contract *SimpleStorage // Generic contract binding to access the raw methods on
}

// SimpleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleStorageCallerRaw struct {
	Contract *SimpleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleStorageTransactorRaw struct {
	Contract *SimpleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// NewSimpleStorageCaller creates a new read-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageCaller(address common.Address, caller bind.ContractCaller) (*SimpleStorageCaller, error) {
	contract, err := bindSimpleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageCaller{contract: contract}, nil
}

// NewSimpleStorageTransactor creates a new write-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleStorageTransactor, error) {
	contract, err := bindSimpleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageTransactor{contract: contract}, nil
}

// NewSimpleStorageFilterer creates a new log filterer instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleStorageFilterer, error) {
	contract, err := bindSimpleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageFilterer{contract: contract}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.SimpleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8)
func (_SimpleStorage *SimpleStorageCaller) Get(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8)
func (_SimpleStorage *SimpleStorageSession) Get() (uint8, error) {
	return _SimpleStorage.Contract.Get(&_SimpleStorage.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8)
func (_SimpleStorage *SimpleStorageCallerSession) Get() (uint8, error) {
	return _SimpleStorage.Contract.Get(&_SimpleStorage.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x24b8ba5f.
//
// Solidity: function set(uint8 x) returns()
func (_SimpleStorage *SimpleStorageTransactor) Set(opts *bind.TransactOpts, x uint8) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x24b8ba5f.
//
// Solidity: function set(uint8 x) returns()
func (_SimpleStorage *SimpleStorageSession) Set(x uint8) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Set(&_SimpleStorage.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x24b8ba5f.
//
// Solidity: function set(uint8 x) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) Set(x uint8) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Set(&_SimpleStorage.TransactOpts, x)
}

// SimpleStorageDataSetIterator is returned from FilterDataSet and is used to iterate over the raw logs and unpacked data for DataSet events raised by the SimpleStorage contract.
type SimpleStorageDataSetIterator struct {
	Event *SimpleStorageDataSet // Event containing the contract specifics and raw log

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
func (it *SimpleStorageDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleStorageDataSet)
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
		it.Event = new(SimpleStorageDataSet)
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
func (it *SimpleStorageDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleStorageDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleStorageDataSet represents a DataSet event raised by the SimpleStorage contract.
type SimpleStorageDataSet struct {
	OldData *big.Int
	NewData *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDataSet is a free log retrieval operation binding the contract event 0xc911a63e29945a04493ec58a89a96aa00a33c2609f1c96d4e506a7fb094e4c94.
//
// Solidity: event DataSet(uint256 indexed oldData, uint256 indexed newData)
func (_SimpleStorage *SimpleStorageFilterer) FilterDataSet(opts *bind.FilterOpts, oldData []*big.Int, newData []*big.Int) (*SimpleStorageDataSetIterator, error) {

	var oldDataRule []interface{}
	for _, oldDataItem := range oldData {
		oldDataRule = append(oldDataRule, oldDataItem)
	}
	var newDataRule []interface{}
	for _, newDataItem := range newData {
		newDataRule = append(newDataRule, newDataItem)
	}

	logs, sub, err := _SimpleStorage.contract.FilterLogs(opts, "DataSet", oldDataRule, newDataRule)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageDataSetIterator{contract: _SimpleStorage.contract, event: "DataSet", logs: logs, sub: sub}, nil
}

// WatchDataSet is a free log subscription operation binding the contract event 0xc911a63e29945a04493ec58a89a96aa00a33c2609f1c96d4e506a7fb094e4c94.
//
// Solidity: event DataSet(uint256 indexed oldData, uint256 indexed newData)
func (_SimpleStorage *SimpleStorageFilterer) WatchDataSet(opts *bind.WatchOpts, sink chan<- *SimpleStorageDataSet, oldData []*big.Int, newData []*big.Int) (event.Subscription, error) {

	var oldDataRule []interface{}
	for _, oldDataItem := range oldData {
		oldDataRule = append(oldDataRule, oldDataItem)
	}
	var newDataRule []interface{}
	for _, newDataItem := range newData {
		newDataRule = append(newDataRule, newDataItem)
	}

	logs, sub, err := _SimpleStorage.contract.WatchLogs(opts, "DataSet", oldDataRule, newDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleStorageDataSet)
				if err := _SimpleStorage.contract.UnpackLog(event, "DataSet", log); err != nil {
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

// ParseDataSet is a log parse operation binding the contract event 0xc911a63e29945a04493ec58a89a96aa00a33c2609f1c96d4e506a7fb094e4c94.
//
// Solidity: event DataSet(uint256 indexed oldData, uint256 indexed newData)
func (_SimpleStorage *SimpleStorageFilterer) ParseDataSet(log types.Log) (*SimpleStorageDataSet, error) {
	event := new(SimpleStorageDataSet)
	if err := _SimpleStorage.contract.UnpackLog(event, "DataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
