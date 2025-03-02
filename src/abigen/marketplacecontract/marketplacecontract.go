// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplacecontract

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

// MarketplacecontractMetaData contains all meta data concerning the Marketplacecontract contract.
var MarketplacecontractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addProduct\",\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"guid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProduct\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_escrowAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setEscrowAddress\",\"inputs\":[{\"name\":\"escrowAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePrice\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateStock\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ItemAdded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"guid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StockUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Marketplace_InvalidProduct\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Marketplace_Unauthorized\",\"inputs\":[]}]",
}

// MarketplacecontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplacecontractMetaData.ABI instead.
var MarketplacecontractABI = MarketplacecontractMetaData.ABI

// Marketplacecontract is an auto generated Go binding around an Ethereum contract.
type Marketplacecontract struct {
	MarketplacecontractCaller     // Read-only binding to the contract
	MarketplacecontractTransactor // Write-only binding to the contract
	MarketplacecontractFilterer   // Log filterer for contract events
}

// MarketplacecontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplacecontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplacecontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplacecontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplacecontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplacecontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplacecontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplacecontractSession struct {
	Contract     *Marketplacecontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MarketplacecontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplacecontractCallerSession struct {
	Contract *MarketplacecontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MarketplacecontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplacecontractTransactorSession struct {
	Contract     *MarketplacecontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MarketplacecontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplacecontractRaw struct {
	Contract *Marketplacecontract // Generic contract binding to access the raw methods on
}

// MarketplacecontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplacecontractCallerRaw struct {
	Contract *MarketplacecontractCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplacecontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplacecontractTransactorRaw struct {
	Contract *MarketplacecontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplacecontract creates a new instance of Marketplacecontract, bound to a specific deployed contract.
func NewMarketplacecontract(address common.Address, backend bind.ContractBackend) (*Marketplacecontract, error) {
	contract, err := bindMarketplacecontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplacecontract{MarketplacecontractCaller: MarketplacecontractCaller{contract: contract}, MarketplacecontractTransactor: MarketplacecontractTransactor{contract: contract}, MarketplacecontractFilterer: MarketplacecontractFilterer{contract: contract}}, nil
}

// NewMarketplacecontractCaller creates a new read-only instance of Marketplacecontract, bound to a specific deployed contract.
func NewMarketplacecontractCaller(address common.Address, caller bind.ContractCaller) (*MarketplacecontractCaller, error) {
	contract, err := bindMarketplacecontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplacecontractCaller{contract: contract}, nil
}

// NewMarketplacecontractTransactor creates a new write-only instance of Marketplacecontract, bound to a specific deployed contract.
func NewMarketplacecontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplacecontractTransactor, error) {
	contract, err := bindMarketplacecontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplacecontractTransactor{contract: contract}, nil
}

// NewMarketplacecontractFilterer creates a new log filterer instance of Marketplacecontract, bound to a specific deployed contract.
func NewMarketplacecontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplacecontractFilterer, error) {
	contract, err := bindMarketplacecontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplacecontractFilterer{contract: contract}, nil
}

// bindMarketplacecontract binds a generic wrapper to an already deployed contract.
func bindMarketplacecontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplacecontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplacecontract *MarketplacecontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplacecontract.Contract.MarketplacecontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplacecontract *MarketplacecontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.MarketplacecontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplacecontract *MarketplacecontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.MarketplacecontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplacecontract *MarketplacecontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplacecontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplacecontract *MarketplacecontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplacecontract *MarketplacecontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.contract.Transact(opts, method, params...)
}

// GetProduct is a free data retrieval call binding the contract method 0xb9db15b4.
//
// Solidity: function getProduct(uint256 id) view returns(address, uint256, uint256)
func (_Marketplacecontract *MarketplacecontractCaller) GetProduct(opts *bind.CallOpts, id *big.Int) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Marketplacecontract.contract.Call(opts, &out, "getProduct", id)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetProduct is a free data retrieval call binding the contract method 0xb9db15b4.
//
// Solidity: function getProduct(uint256 id) view returns(address, uint256, uint256)
func (_Marketplacecontract *MarketplacecontractSession) GetProduct(id *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _Marketplacecontract.Contract.GetProduct(&_Marketplacecontract.CallOpts, id)
}

// GetProduct is a free data retrieval call binding the contract method 0xb9db15b4.
//
// Solidity: function getProduct(uint256 id) view returns(address, uint256, uint256)
func (_Marketplacecontract *MarketplacecontractCallerSession) GetProduct(id *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _Marketplacecontract.Contract.GetProduct(&_Marketplacecontract.CallOpts, id)
}

// SEscrowAddress is a free data retrieval call binding the contract method 0xbbf7fda1.
//
// Solidity: function s_escrowAddress() view returns(address)
func (_Marketplacecontract *MarketplacecontractCaller) SEscrowAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplacecontract.contract.Call(opts, &out, "s_escrowAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEscrowAddress is a free data retrieval call binding the contract method 0xbbf7fda1.
//
// Solidity: function s_escrowAddress() view returns(address)
func (_Marketplacecontract *MarketplacecontractSession) SEscrowAddress() (common.Address, error) {
	return _Marketplacecontract.Contract.SEscrowAddress(&_Marketplacecontract.CallOpts)
}

// SEscrowAddress is a free data retrieval call binding the contract method 0xbbf7fda1.
//
// Solidity: function s_escrowAddress() view returns(address)
func (_Marketplacecontract *MarketplacecontractCallerSession) SEscrowAddress() (common.Address, error) {
	return _Marketplacecontract.Contract.SEscrowAddress(&_Marketplacecontract.CallOpts)
}

// AddProduct is a paid mutator transaction binding the contract method 0x804ee2b5.
//
// Solidity: function addProduct(uint256 price, uint256 stock, string guid) returns(uint256)
func (_Marketplacecontract *MarketplacecontractTransactor) AddProduct(opts *bind.TransactOpts, price *big.Int, stock *big.Int, guid string) (*types.Transaction, error) {
	return _Marketplacecontract.contract.Transact(opts, "addProduct", price, stock, guid)
}

// AddProduct is a paid mutator transaction binding the contract method 0x804ee2b5.
//
// Solidity: function addProduct(uint256 price, uint256 stock, string guid) returns(uint256)
func (_Marketplacecontract *MarketplacecontractSession) AddProduct(price *big.Int, stock *big.Int, guid string) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.AddProduct(&_Marketplacecontract.TransactOpts, price, stock, guid)
}

// AddProduct is a paid mutator transaction binding the contract method 0x804ee2b5.
//
// Solidity: function addProduct(uint256 price, uint256 stock, string guid) returns(uint256)
func (_Marketplacecontract *MarketplacecontractTransactorSession) AddProduct(price *big.Int, stock *big.Int, guid string) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.AddProduct(&_Marketplacecontract.TransactOpts, price, stock, guid)
}

// SetEscrowAddress is a paid mutator transaction binding the contract method 0xddeb63b5.
//
// Solidity: function setEscrowAddress(address escrowAddress) returns()
func (_Marketplacecontract *MarketplacecontractTransactor) SetEscrowAddress(opts *bind.TransactOpts, escrowAddress common.Address) (*types.Transaction, error) {
	return _Marketplacecontract.contract.Transact(opts, "setEscrowAddress", escrowAddress)
}

// SetEscrowAddress is a paid mutator transaction binding the contract method 0xddeb63b5.
//
// Solidity: function setEscrowAddress(address escrowAddress) returns()
func (_Marketplacecontract *MarketplacecontractSession) SetEscrowAddress(escrowAddress common.Address) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.SetEscrowAddress(&_Marketplacecontract.TransactOpts, escrowAddress)
}

// SetEscrowAddress is a paid mutator transaction binding the contract method 0xddeb63b5.
//
// Solidity: function setEscrowAddress(address escrowAddress) returns()
func (_Marketplacecontract *MarketplacecontractTransactorSession) SetEscrowAddress(escrowAddress common.Address) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.SetEscrowAddress(&_Marketplacecontract.TransactOpts, escrowAddress)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x82367b2d.
//
// Solidity: function updatePrice(uint256 id, uint256 price) returns()
func (_Marketplacecontract *MarketplacecontractTransactor) UpdatePrice(opts *bind.TransactOpts, id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Marketplacecontract.contract.Transact(opts, "updatePrice", id, price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x82367b2d.
//
// Solidity: function updatePrice(uint256 id, uint256 price) returns()
func (_Marketplacecontract *MarketplacecontractSession) UpdatePrice(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.UpdatePrice(&_Marketplacecontract.TransactOpts, id, price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x82367b2d.
//
// Solidity: function updatePrice(uint256 id, uint256 price) returns()
func (_Marketplacecontract *MarketplacecontractTransactorSession) UpdatePrice(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.UpdatePrice(&_Marketplacecontract.TransactOpts, id, price)
}

// UpdateStock is a paid mutator transaction binding the contract method 0x9dfc3dd1.
//
// Solidity: function updateStock(uint256 id, uint256 stock) returns()
func (_Marketplacecontract *MarketplacecontractTransactor) UpdateStock(opts *bind.TransactOpts, id *big.Int, stock *big.Int) (*types.Transaction, error) {
	return _Marketplacecontract.contract.Transact(opts, "updateStock", id, stock)
}

// UpdateStock is a paid mutator transaction binding the contract method 0x9dfc3dd1.
//
// Solidity: function updateStock(uint256 id, uint256 stock) returns()
func (_Marketplacecontract *MarketplacecontractSession) UpdateStock(id *big.Int, stock *big.Int) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.UpdateStock(&_Marketplacecontract.TransactOpts, id, stock)
}

// UpdateStock is a paid mutator transaction binding the contract method 0x9dfc3dd1.
//
// Solidity: function updateStock(uint256 id, uint256 stock) returns()
func (_Marketplacecontract *MarketplacecontractTransactorSession) UpdateStock(id *big.Int, stock *big.Int) (*types.Transaction, error) {
	return _Marketplacecontract.Contract.UpdateStock(&_Marketplacecontract.TransactOpts, id, stock)
}

// MarketplacecontractItemAddedIterator is returned from FilterItemAdded and is used to iterate over the raw logs and unpacked data for ItemAdded events raised by the Marketplacecontract contract.
type MarketplacecontractItemAddedIterator struct {
	Event *MarketplacecontractItemAdded // Event containing the contract specifics and raw log

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
func (it *MarketplacecontractItemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacecontractItemAdded)
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
		it.Event = new(MarketplacecontractItemAdded)
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
func (it *MarketplacecontractItemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacecontractItemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacecontractItemAdded represents a ItemAdded event raised by the Marketplacecontract contract.
type MarketplacecontractItemAdded struct {
	Id     *big.Int
	Seller common.Address
	Guid   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterItemAdded is a free log retrieval operation binding the contract event 0x846d9a44fcae2ae5c88f54e5d384c87f8dcd917e1141c30b4f4aa41d0543224a.
//
// Solidity: event ItemAdded(uint256 indexed id, address indexed seller, string guid)
func (_Marketplacecontract *MarketplacecontractFilterer) FilterItemAdded(opts *bind.FilterOpts, id []*big.Int, seller []common.Address) (*MarketplacecontractItemAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplacecontract.contract.FilterLogs(opts, "ItemAdded", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacecontractItemAddedIterator{contract: _Marketplacecontract.contract, event: "ItemAdded", logs: logs, sub: sub}, nil
}

// WatchItemAdded is a free log subscription operation binding the contract event 0x846d9a44fcae2ae5c88f54e5d384c87f8dcd917e1141c30b4f4aa41d0543224a.
//
// Solidity: event ItemAdded(uint256 indexed id, address indexed seller, string guid)
func (_Marketplacecontract *MarketplacecontractFilterer) WatchItemAdded(opts *bind.WatchOpts, sink chan<- *MarketplacecontractItemAdded, id []*big.Int, seller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplacecontract.contract.WatchLogs(opts, "ItemAdded", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacecontractItemAdded)
				if err := _Marketplacecontract.contract.UnpackLog(event, "ItemAdded", log); err != nil {
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

// ParseItemAdded is a log parse operation binding the contract event 0x846d9a44fcae2ae5c88f54e5d384c87f8dcd917e1141c30b4f4aa41d0543224a.
//
// Solidity: event ItemAdded(uint256 indexed id, address indexed seller, string guid)
func (_Marketplacecontract *MarketplacecontractFilterer) ParseItemAdded(log types.Log) (*MarketplacecontractItemAdded, error) {
	event := new(MarketplacecontractItemAdded)
	if err := _Marketplacecontract.contract.UnpackLog(event, "ItemAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacecontractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Marketplacecontract contract.
type MarketplacecontractPriceUpdatedIterator struct {
	Event *MarketplacecontractPriceUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplacecontractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacecontractPriceUpdated)
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
		it.Event = new(MarketplacecontractPriceUpdated)
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
func (it *MarketplacecontractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacecontractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacecontractPriceUpdated represents a PriceUpdated event raised by the Marketplacecontract contract.
type MarketplacecontractPriceUpdated struct {
	Id    *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(uint256 indexed id, uint256 indexed price)
func (_Marketplacecontract *MarketplacecontractFilterer) FilterPriceUpdated(opts *bind.FilterOpts, id []*big.Int, price []*big.Int) (*MarketplacecontractPriceUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _Marketplacecontract.contract.FilterLogs(opts, "PriceUpdated", idRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacecontractPriceUpdatedIterator{contract: _Marketplacecontract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(uint256 indexed id, uint256 indexed price)
func (_Marketplacecontract *MarketplacecontractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *MarketplacecontractPriceUpdated, id []*big.Int, price []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _Marketplacecontract.contract.WatchLogs(opts, "PriceUpdated", idRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacecontractPriceUpdated)
				if err := _Marketplacecontract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(uint256 indexed id, uint256 indexed price)
func (_Marketplacecontract *MarketplacecontractFilterer) ParsePriceUpdated(log types.Log) (*MarketplacecontractPriceUpdated, error) {
	event := new(MarketplacecontractPriceUpdated)
	if err := _Marketplacecontract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacecontractStockUpdatedIterator is returned from FilterStockUpdated and is used to iterate over the raw logs and unpacked data for StockUpdated events raised by the Marketplacecontract contract.
type MarketplacecontractStockUpdatedIterator struct {
	Event *MarketplacecontractStockUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplacecontractStockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacecontractStockUpdated)
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
		it.Event = new(MarketplacecontractStockUpdated)
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
func (it *MarketplacecontractStockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacecontractStockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacecontractStockUpdated represents a StockUpdated event raised by the Marketplacecontract contract.
type MarketplacecontractStockUpdated struct {
	Id       *big.Int
	NewStock *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStockUpdated is a free log retrieval operation binding the contract event 0xebcfd4ce21b6aefd071d3e0f4bfa57e491b5308a1ffb5a3a8e4805e23f00b8d7.
//
// Solidity: event StockUpdated(uint256 indexed id, uint256 indexed newStock)
func (_Marketplacecontract *MarketplacecontractFilterer) FilterStockUpdated(opts *bind.FilterOpts, id []*big.Int, newStock []*big.Int) (*MarketplacecontractStockUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var newStockRule []interface{}
	for _, newStockItem := range newStock {
		newStockRule = append(newStockRule, newStockItem)
	}

	logs, sub, err := _Marketplacecontract.contract.FilterLogs(opts, "StockUpdated", idRule, newStockRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacecontractStockUpdatedIterator{contract: _Marketplacecontract.contract, event: "StockUpdated", logs: logs, sub: sub}, nil
}

// WatchStockUpdated is a free log subscription operation binding the contract event 0xebcfd4ce21b6aefd071d3e0f4bfa57e491b5308a1ffb5a3a8e4805e23f00b8d7.
//
// Solidity: event StockUpdated(uint256 indexed id, uint256 indexed newStock)
func (_Marketplacecontract *MarketplacecontractFilterer) WatchStockUpdated(opts *bind.WatchOpts, sink chan<- *MarketplacecontractStockUpdated, id []*big.Int, newStock []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var newStockRule []interface{}
	for _, newStockItem := range newStock {
		newStockRule = append(newStockRule, newStockItem)
	}

	logs, sub, err := _Marketplacecontract.contract.WatchLogs(opts, "StockUpdated", idRule, newStockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacecontractStockUpdated)
				if err := _Marketplacecontract.contract.UnpackLog(event, "StockUpdated", log); err != nil {
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

// ParseStockUpdated is a log parse operation binding the contract event 0xebcfd4ce21b6aefd071d3e0f4bfa57e491b5308a1ffb5a3a8e4805e23f00b8d7.
//
// Solidity: event StockUpdated(uint256 indexed id, uint256 indexed newStock)
func (_Marketplacecontract *MarketplacecontractFilterer) ParseStockUpdated(log types.Log) (*MarketplacecontractStockUpdated, error) {
	event := new(MarketplacecontractStockUpdated)
	if err := _Marketplacecontract.contract.UnpackLog(event, "StockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
