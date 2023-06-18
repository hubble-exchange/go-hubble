// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IOrderBookOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrder struct {
	AmmIndex          *big.Int
	Trader            common.Address
	BaseAssetQuantity *big.Int
	Price             *big.Int
	Salt              *big.Int
	ReduceOnly        bool
}

// OrderBookMetaData contains all meta data concerning the OrderBook contract.
var OrderBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearingHouse\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marginAccount\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"err\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toLiquidate\",\"type\":\"uint256\"}],\"name\":\"LiquidationError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openInterestNotional\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LiquidationOrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"err\",\"type\":\"string\"}],\"name\":\"OrderMatchingError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIOrderBook.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrderPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash0\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openInterestNotional\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bibliophile\",\"outputs\":[{\"internalType\":\"contractIHubbleBibliophile\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearingHouse\",\"outputs\":[{\"internalType\":\"contractIClearingHouse\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order[2]\",\"name\":\"orders\",\"type\":\"tuple[2]\"},{\"internalType\":\"int256\",\"name\":\"fillAmount\",\"type\":\"int256\"}],\"name\":\"executeMatchedOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastTradePrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"lastTradePrices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getRequiredMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredMargin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"minSize\",\"type\":\"int256\"}],\"name\":\"initializeMinSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateAndExecuteOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginAccount\",\"outputs\":[{\"internalType\":\"contractIMarginAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAllowableMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minSizes\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockPlaced\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"filledAmount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"reservedMargin\",\"type\":\"uint256\"},{\"internalType\":\"enumIOrderBook.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"err\",\"type\":\"string\"}],\"name\":\"parseMatchingError\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"placeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"placeOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reduceOnlyAmount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bibliophile\",\"type\":\"address\"}],\"name\":\"setBibliophile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__governance\",\"type\":\"address\"}],\"name\":\"setGovernace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"useNew\",\"type\":\"bool\"}],\"name\":\"setUseNewPricingAlgorithm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setValidatorStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"minSize\",\"type\":\"int256\"}],\"name\":\"updateMinSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAllowableMargin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_takerFee\",\"type\":\"uint256\"}],\"name\":\"updateParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useNewPricingAlgorithm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ammIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"baseAssetQuantity\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"reduceOnly\",\"type\":\"bool\"}],\"internalType\":\"structIOrderBook.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderBookMetaData.ABI instead.
var OrderBookABI = OrderBookMetaData.ABI

// OrderBook is an auto generated Go binding around an Ethereum contract.
type OrderBook struct {
	OrderBookCaller     // Read-only binding to the contract
	OrderBookTransactor // Write-only binding to the contract
	OrderBookFilterer   // Log filterer for contract events
}

// OrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderBookSession struct {
	Contract     *OrderBook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderBookCallerSession struct {
	Contract *OrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderBookTransactorSession struct {
	Contract     *OrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderBookRaw struct {
	Contract *OrderBook // Generic contract binding to access the raw methods on
}

// OrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderBookCallerRaw struct {
	Contract *OrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderBookTransactorRaw struct {
	Contract *OrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderBook creates a new instance of OrderBook, bound to a specific deployed contract.
func NewOrderBook(address common.Address, backend bind.ContractBackend) (*OrderBook, error) {
	contract, err := bindOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderBook{OrderBookCaller: OrderBookCaller{contract: contract}, OrderBookTransactor: OrderBookTransactor{contract: contract}, OrderBookFilterer: OrderBookFilterer{contract: contract}}, nil
}

// NewOrderBookCaller creates a new read-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookCaller(address common.Address, caller bind.ContractCaller) (*OrderBookCaller, error) {
	contract, err := bindOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookCaller{contract: contract}, nil
}

// NewOrderBookTransactor creates a new write-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderBookTransactor, error) {
	contract, err := bindOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookTransactor{contract: contract}, nil
}

// NewOrderBookFilterer creates a new log filterer instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderBookFilterer, error) {
	contract, err := bindOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderBookFilterer{contract: contract}, nil
}

// bindOrderBook binds a generic wrapper to an already deployed contract.
func bindOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.OrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transact(opts, method, params...)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_OrderBook *OrderBookCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_OrderBook *OrderBookSession) ORDERTYPEHASH() ([32]byte, error) {
	return _OrderBook.Contract.ORDERTYPEHASH(&_OrderBook.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_OrderBook *OrderBookCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _OrderBook.Contract.ORDERTYPEHASH(&_OrderBook.CallOpts)
}

// Bibliophile is a free data retrieval call binding the contract method 0xb0ef6e7c.
//
// Solidity: function bibliophile() view returns(address)
func (_OrderBook *OrderBookCaller) Bibliophile(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "bibliophile")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bibliophile is a free data retrieval call binding the contract method 0xb0ef6e7c.
//
// Solidity: function bibliophile() view returns(address)
func (_OrderBook *OrderBookSession) Bibliophile() (common.Address, error) {
	return _OrderBook.Contract.Bibliophile(&_OrderBook.CallOpts)
}

// Bibliophile is a free data retrieval call binding the contract method 0xb0ef6e7c.
//
// Solidity: function bibliophile() view returns(address)
func (_OrderBook *OrderBookCallerSession) Bibliophile() (common.Address, error) {
	return _OrderBook.Contract.Bibliophile(&_OrderBook.CallOpts)
}

// ClearingHouse is a free data retrieval call binding the contract method 0x0af96800.
//
// Solidity: function clearingHouse() view returns(address)
func (_OrderBook *OrderBookCaller) ClearingHouse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "clearingHouse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClearingHouse is a free data retrieval call binding the contract method 0x0af96800.
//
// Solidity: function clearingHouse() view returns(address)
func (_OrderBook *OrderBookSession) ClearingHouse() (common.Address, error) {
	return _OrderBook.Contract.ClearingHouse(&_OrderBook.CallOpts)
}

// ClearingHouse is a free data retrieval call binding the contract method 0x0af96800.
//
// Solidity: function clearingHouse() view returns(address)
func (_OrderBook *OrderBookCallerSession) ClearingHouse() (common.Address, error) {
	return _OrderBook.Contract.ClearingHouse(&_OrderBook.CallOpts)
}

// GetLastTradePrices is a free data retrieval call binding the contract method 0x7114f7f8.
//
// Solidity: function getLastTradePrices() view returns(uint256[] lastTradePrices)
func (_OrderBook *OrderBookCaller) GetLastTradePrices(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getLastTradePrices")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLastTradePrices is a free data retrieval call binding the contract method 0x7114f7f8.
//
// Solidity: function getLastTradePrices() view returns(uint256[] lastTradePrices)
func (_OrderBook *OrderBookSession) GetLastTradePrices() ([]*big.Int, error) {
	return _OrderBook.Contract.GetLastTradePrices(&_OrderBook.CallOpts)
}

// GetLastTradePrices is a free data retrieval call binding the contract method 0x7114f7f8.
//
// Solidity: function getLastTradePrices() view returns(uint256[] lastTradePrices)
func (_OrderBook *OrderBookCallerSession) GetLastTradePrices() ([]*big.Int, error) {
	return _OrderBook.Contract.GetLastTradePrices(&_OrderBook.CallOpts)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xdbe64846.
//
// Solidity: function getOrderHash((uint256,address,int256,uint256,uint256,bool) order) view returns(bytes32)
func (_OrderBook *OrderBookCaller) GetOrderHash(opts *bind.CallOpts, order IOrderBookOrder) ([32]byte, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0xdbe64846.
//
// Solidity: function getOrderHash((uint256,address,int256,uint256,uint256,bool) order) view returns(bytes32)
func (_OrderBook *OrderBookSession) GetOrderHash(order IOrderBookOrder) ([32]byte, error) {
	return _OrderBook.Contract.GetOrderHash(&_OrderBook.CallOpts, order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xdbe64846.
//
// Solidity: function getOrderHash((uint256,address,int256,uint256,uint256,bool) order) view returns(bytes32)
func (_OrderBook *OrderBookCallerSession) GetOrderHash(order IOrderBookOrder) ([32]byte, error) {
	return _OrderBook.Contract.GetOrderHash(&_OrderBook.CallOpts, order)
}

// GetRequiredMargin is a free data retrieval call binding the contract method 0xb7c2e7ae.
//
// Solidity: function getRequiredMargin(int256 baseAssetQuantity, uint256 price) view returns(uint256 requiredMargin)
func (_OrderBook *OrderBookCaller) GetRequiredMargin(opts *bind.CallOpts, baseAssetQuantity *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getRequiredMargin", baseAssetQuantity, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredMargin is a free data retrieval call binding the contract method 0xb7c2e7ae.
//
// Solidity: function getRequiredMargin(int256 baseAssetQuantity, uint256 price) view returns(uint256 requiredMargin)
func (_OrderBook *OrderBookSession) GetRequiredMargin(baseAssetQuantity *big.Int, price *big.Int) (*big.Int, error) {
	return _OrderBook.Contract.GetRequiredMargin(&_OrderBook.CallOpts, baseAssetQuantity, price)
}

// GetRequiredMargin is a free data retrieval call binding the contract method 0xb7c2e7ae.
//
// Solidity: function getRequiredMargin(int256 baseAssetQuantity, uint256 price) view returns(uint256 requiredMargin)
func (_OrderBook *OrderBookCallerSession) GetRequiredMargin(baseAssetQuantity *big.Int, price *big.Int) (*big.Int, error) {
	return _OrderBook.Contract.GetRequiredMargin(&_OrderBook.CallOpts, baseAssetQuantity, price)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_OrderBook *OrderBookCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_OrderBook *OrderBookSession) Governance() (common.Address, error) {
	return _OrderBook.Contract.Governance(&_OrderBook.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_OrderBook *OrderBookCallerSession) Governance() (common.Address, error) {
	return _OrderBook.Contract.Governance(&_OrderBook.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_OrderBook *OrderBookCaller) IsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "isValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_OrderBook *OrderBookSession) IsValidator(arg0 common.Address) (bool, error) {
	return _OrderBook.Contract.IsValidator(&_OrderBook.CallOpts, arg0)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address ) view returns(bool)
func (_OrderBook *OrderBookCallerSession) IsValidator(arg0 common.Address) (bool, error) {
	return _OrderBook.Contract.IsValidator(&_OrderBook.CallOpts, arg0)
}

// MarginAccount is a free data retrieval call binding the contract method 0xf742269d.
//
// Solidity: function marginAccount() view returns(address)
func (_OrderBook *OrderBookCaller) MarginAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "marginAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarginAccount is a free data retrieval call binding the contract method 0xf742269d.
//
// Solidity: function marginAccount() view returns(address)
func (_OrderBook *OrderBookSession) MarginAccount() (common.Address, error) {
	return _OrderBook.Contract.MarginAccount(&_OrderBook.CallOpts)
}

// MarginAccount is a free data retrieval call binding the contract method 0xf742269d.
//
// Solidity: function marginAccount() view returns(address)
func (_OrderBook *OrderBookCallerSession) MarginAccount() (common.Address, error) {
	return _OrderBook.Contract.MarginAccount(&_OrderBook.CallOpts)
}

// MinAllowableMargin is a free data retrieval call binding the contract method 0x91735a39.
//
// Solidity: function minAllowableMargin() view returns(uint256)
func (_OrderBook *OrderBookCaller) MinAllowableMargin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "minAllowableMargin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAllowableMargin is a free data retrieval call binding the contract method 0x91735a39.
//
// Solidity: function minAllowableMargin() view returns(uint256)
func (_OrderBook *OrderBookSession) MinAllowableMargin() (*big.Int, error) {
	return _OrderBook.Contract.MinAllowableMargin(&_OrderBook.CallOpts)
}

// MinAllowableMargin is a free data retrieval call binding the contract method 0x91735a39.
//
// Solidity: function minAllowableMargin() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) MinAllowableMargin() (*big.Int, error) {
	return _OrderBook.Contract.MinAllowableMargin(&_OrderBook.CallOpts)
}

// MinSizes is a free data retrieval call binding the contract method 0x8274da46.
//
// Solidity: function minSizes(uint256 ) view returns(int256)
func (_OrderBook *OrderBookCaller) MinSizes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "minSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSizes is a free data retrieval call binding the contract method 0x8274da46.
//
// Solidity: function minSizes(uint256 ) view returns(int256)
func (_OrderBook *OrderBookSession) MinSizes(arg0 *big.Int) (*big.Int, error) {
	return _OrderBook.Contract.MinSizes(&_OrderBook.CallOpts, arg0)
}

// MinSizes is a free data retrieval call binding the contract method 0x8274da46.
//
// Solidity: function minSizes(uint256 ) view returns(int256)
func (_OrderBook *OrderBookCallerSession) MinSizes(arg0 *big.Int) (*big.Int, error) {
	return _OrderBook.Contract.MinSizes(&_OrderBook.CallOpts, arg0)
}

// OrderInfo is a free data retrieval call binding the contract method 0x238e203f.
//
// Solidity: function orderInfo(bytes32 ) view returns(uint256 blockPlaced, int256 filledAmount, uint256 reservedMargin, uint8 status)
func (_OrderBook *OrderBookCaller) OrderInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	BlockPlaced    *big.Int
	FilledAmount   *big.Int
	ReservedMargin *big.Int
	Status         uint8
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "orderInfo", arg0)

	outstruct := new(struct {
		BlockPlaced    *big.Int
		FilledAmount   *big.Int
		ReservedMargin *big.Int
		Status         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockPlaced = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FilledAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReservedMargin = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// OrderInfo is a free data retrieval call binding the contract method 0x238e203f.
//
// Solidity: function orderInfo(bytes32 ) view returns(uint256 blockPlaced, int256 filledAmount, uint256 reservedMargin, uint8 status)
func (_OrderBook *OrderBookSession) OrderInfo(arg0 [32]byte) (struct {
	BlockPlaced    *big.Int
	FilledAmount   *big.Int
	ReservedMargin *big.Int
	Status         uint8
}, error) {
	return _OrderBook.Contract.OrderInfo(&_OrderBook.CallOpts, arg0)
}

// OrderInfo is a free data retrieval call binding the contract method 0x238e203f.
//
// Solidity: function orderInfo(bytes32 ) view returns(uint256 blockPlaced, int256 filledAmount, uint256 reservedMargin, uint8 status)
func (_OrderBook *OrderBookCallerSession) OrderInfo(arg0 [32]byte) (struct {
	BlockPlaced    *big.Int
	FilledAmount   *big.Int
	ReservedMargin *big.Int
	Status         uint8
}, error) {
	return _OrderBook.Contract.OrderInfo(&_OrderBook.CallOpts, arg0)
}

// ParseMatchingError is a free data retrieval call binding the contract method 0x656f1645.
//
// Solidity: function parseMatchingError(string err) pure returns(bytes32 orderHash, string reason)
func (_OrderBook *OrderBookCaller) ParseMatchingError(opts *bind.CallOpts, errStr string) (struct {
	OrderHash [32]byte
	Reason    string
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "parseMatchingError", errStr)

	outstruct := new(struct {
		OrderHash [32]byte
		Reason    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Reason = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ParseMatchingError is a free data retrieval call binding the contract method 0x656f1645.
//
// Solidity: function parseMatchingError(string err) pure returns(bytes32 orderHash, string reason)
func (_OrderBook *OrderBookSession) ParseMatchingError(err string) (struct {
	OrderHash [32]byte
	Reason    string
}, error) {
	return _OrderBook.Contract.ParseMatchingError(&_OrderBook.CallOpts, err)
}

// ParseMatchingError is a free data retrieval call binding the contract method 0x656f1645.
//
// Solidity: function parseMatchingError(string err) pure returns(bytes32 orderHash, string reason)
func (_OrderBook *OrderBookCallerSession) ParseMatchingError(err string) (struct {
	OrderHash [32]byte
	Reason    string
}, error) {
	return _OrderBook.Contract.ParseMatchingError(&_OrderBook.CallOpts, err)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OrderBook *OrderBookCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OrderBook *OrderBookSession) Paused() (bool, error) {
	return _OrderBook.Contract.Paused(&_OrderBook.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OrderBook *OrderBookCallerSession) Paused() (bool, error) {
	return _OrderBook.Contract.Paused(&_OrderBook.CallOpts)
}

// ReduceOnlyAmount is a free data retrieval call binding the contract method 0x81d5702b.
//
// Solidity: function reduceOnlyAmount(address , uint256 ) view returns(int256)
func (_OrderBook *OrderBookCaller) ReduceOnlyAmount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "reduceOnlyAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReduceOnlyAmount is a free data retrieval call binding the contract method 0x81d5702b.
//
// Solidity: function reduceOnlyAmount(address , uint256 ) view returns(int256)
func (_OrderBook *OrderBookSession) ReduceOnlyAmount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _OrderBook.Contract.ReduceOnlyAmount(&_OrderBook.CallOpts, arg0, arg1)
}

// ReduceOnlyAmount is a free data retrieval call binding the contract method 0x81d5702b.
//
// Solidity: function reduceOnlyAmount(address , uint256 ) view returns(int256)
func (_OrderBook *OrderBookCallerSession) ReduceOnlyAmount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _OrderBook.Contract.ReduceOnlyAmount(&_OrderBook.CallOpts, arg0, arg1)
}

// TakerFee is a free data retrieval call binding the contract method 0x43f0179b.
//
// Solidity: function takerFee() view returns(uint256)
func (_OrderBook *OrderBookCaller) TakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "takerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TakerFee is a free data retrieval call binding the contract method 0x43f0179b.
//
// Solidity: function takerFee() view returns(uint256)
func (_OrderBook *OrderBookSession) TakerFee() (*big.Int, error) {
	return _OrderBook.Contract.TakerFee(&_OrderBook.CallOpts)
}

// TakerFee is a free data retrieval call binding the contract method 0x43f0179b.
//
// Solidity: function takerFee() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) TakerFee() (*big.Int, error) {
	return _OrderBook.Contract.TakerFee(&_OrderBook.CallOpts)
}

// UseNewPricingAlgorithm is a free data retrieval call binding the contract method 0xf804baa4.
//
// Solidity: function useNewPricingAlgorithm() view returns(uint256)
func (_OrderBook *OrderBookCaller) UseNewPricingAlgorithm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "useNewPricingAlgorithm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UseNewPricingAlgorithm is a free data retrieval call binding the contract method 0xf804baa4.
//
// Solidity: function useNewPricingAlgorithm() view returns(uint256)
func (_OrderBook *OrderBookSession) UseNewPricingAlgorithm() (*big.Int, error) {
	return _OrderBook.Contract.UseNewPricingAlgorithm(&_OrderBook.CallOpts)
}

// UseNewPricingAlgorithm is a free data retrieval call binding the contract method 0xf804baa4.
//
// Solidity: function useNewPricingAlgorithm() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) UseNewPricingAlgorithm() (*big.Int, error) {
	return _OrderBook.Contract.UseNewPricingAlgorithm(&_OrderBook.CallOpts)
}

// VerifySigner is a free data retrieval call binding the contract method 0x22dae637.
//
// Solidity: function verifySigner((uint256,address,int256,uint256,uint256,bool) order, bytes signature) view returns(address, bytes32)
func (_OrderBook *OrderBookCaller) VerifySigner(opts *bind.CallOpts, order IOrderBookOrder, signature []byte) (common.Address, [32]byte, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "verifySigner", order, signature)

	if err != nil {
		return *new(common.Address), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// VerifySigner is a free data retrieval call binding the contract method 0x22dae637.
//
// Solidity: function verifySigner((uint256,address,int256,uint256,uint256,bool) order, bytes signature) view returns(address, bytes32)
func (_OrderBook *OrderBookSession) VerifySigner(order IOrderBookOrder, signature []byte) (common.Address, [32]byte, error) {
	return _OrderBook.Contract.VerifySigner(&_OrderBook.CallOpts, order, signature)
}

// VerifySigner is a free data retrieval call binding the contract method 0x22dae637.
//
// Solidity: function verifySigner((uint256,address,int256,uint256,uint256,bool) order, bytes signature) view returns(address, bytes32)
func (_OrderBook *OrderBookCallerSession) VerifySigner(order IOrderBookOrder, signature []byte) (common.Address, [32]byte, error) {
	return _OrderBook.Contract.VerifySigner(&_OrderBook.CallOpts, order, signature)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2c82ce17.
//
// Solidity: function cancelOrder((uint256,address,int256,uint256,uint256,bool) order) returns()
func (_OrderBook *OrderBookTransactor) CancelOrder(opts *bind.TransactOpts, order IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2c82ce17.
//
// Solidity: function cancelOrder((uint256,address,int256,uint256,uint256,bool) order) returns()
func (_OrderBook *OrderBookSession) CancelOrder(order IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelOrder(&_OrderBook.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2c82ce17.
//
// Solidity: function cancelOrder((uint256,address,int256,uint256,uint256,bool) order) returns()
func (_OrderBook *OrderBookTransactorSession) CancelOrder(order IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelOrder(&_OrderBook.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xf6eec9ce.
//
// Solidity: function cancelOrders((uint256,address,int256,uint256,uint256,bool)[] orders) returns()
func (_OrderBook *OrderBookTransactor) CancelOrders(opts *bind.TransactOpts, orders []IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xf6eec9ce.
//
// Solidity: function cancelOrders((uint256,address,int256,uint256,uint256,bool)[] orders) returns()
func (_OrderBook *OrderBookSession) CancelOrders(orders []IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelOrders(&_OrderBook.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xf6eec9ce.
//
// Solidity: function cancelOrders((uint256,address,int256,uint256,uint256,bool)[] orders) returns()
func (_OrderBook *OrderBookTransactorSession) CancelOrders(orders []IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelOrders(&_OrderBook.TransactOpts, orders)
}

// ExecuteMatchedOrders is a paid mutator transaction binding the contract method 0x49608b58.
//
// Solidity: function executeMatchedOrders((uint256,address,int256,uint256,uint256,bool)[2] orders, int256 fillAmount) returns()
func (_OrderBook *OrderBookTransactor) ExecuteMatchedOrders(opts *bind.TransactOpts, orders [2]IOrderBookOrder, fillAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "executeMatchedOrders", orders, fillAmount)
}

// ExecuteMatchedOrders is a paid mutator transaction binding the contract method 0x49608b58.
//
// Solidity: function executeMatchedOrders((uint256,address,int256,uint256,uint256,bool)[2] orders, int256 fillAmount) returns()
func (_OrderBook *OrderBookSession) ExecuteMatchedOrders(orders [2]IOrderBookOrder, fillAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteMatchedOrders(&_OrderBook.TransactOpts, orders, fillAmount)
}

// ExecuteMatchedOrders is a paid mutator transaction binding the contract method 0x49608b58.
//
// Solidity: function executeMatchedOrders((uint256,address,int256,uint256,uint256,bool)[2] orders, int256 fillAmount) returns()
func (_OrderBook *OrderBookTransactorSession) ExecuteMatchedOrders(orders [2]IOrderBookOrder, fillAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteMatchedOrders(&_OrderBook.TransactOpts, orders, fillAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _version, address _governance) returns()
func (_OrderBook *OrderBookTransactor) Initialize(opts *bind.TransactOpts, _name string, _version string, _governance common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "initialize", _name, _version, _governance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _version, address _governance) returns()
func (_OrderBook *OrderBookSession) Initialize(_name string, _version string, _governance common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.Initialize(&_OrderBook.TransactOpts, _name, _version, _governance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _version, address _governance) returns()
func (_OrderBook *OrderBookTransactorSession) Initialize(_name string, _version string, _governance common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.Initialize(&_OrderBook.TransactOpts, _name, _version, _governance)
}

// InitializeMinSize is a paid mutator transaction binding the contract method 0x0e4bfc51.
//
// Solidity: function initializeMinSize(int256 minSize) returns()
func (_OrderBook *OrderBookTransactor) InitializeMinSize(opts *bind.TransactOpts, minSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "initializeMinSize", minSize)
}

// InitializeMinSize is a paid mutator transaction binding the contract method 0x0e4bfc51.
//
// Solidity: function initializeMinSize(int256 minSize) returns()
func (_OrderBook *OrderBookSession) InitializeMinSize(minSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.InitializeMinSize(&_OrderBook.TransactOpts, minSize)
}

// InitializeMinSize is a paid mutator transaction binding the contract method 0x0e4bfc51.
//
// Solidity: function initializeMinSize(int256 minSize) returns()
func (_OrderBook *OrderBookTransactorSession) InitializeMinSize(minSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.InitializeMinSize(&_OrderBook.TransactOpts, minSize)
}

// LiquidateAndExecuteOrder is a paid mutator transaction binding the contract method 0xf09748df.
//
// Solidity: function liquidateAndExecuteOrder(address trader, (uint256,address,int256,uint256,uint256,bool) order, uint256 liquidationAmount) returns()
func (_OrderBook *OrderBookTransactor) LiquidateAndExecuteOrder(opts *bind.TransactOpts, trader common.Address, order IOrderBookOrder, liquidationAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "liquidateAndExecuteOrder", trader, order, liquidationAmount)
}

// LiquidateAndExecuteOrder is a paid mutator transaction binding the contract method 0xf09748df.
//
// Solidity: function liquidateAndExecuteOrder(address trader, (uint256,address,int256,uint256,uint256,bool) order, uint256 liquidationAmount) returns()
func (_OrderBook *OrderBookSession) LiquidateAndExecuteOrder(trader common.Address, order IOrderBookOrder, liquidationAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.LiquidateAndExecuteOrder(&_OrderBook.TransactOpts, trader, order, liquidationAmount)
}

// LiquidateAndExecuteOrder is a paid mutator transaction binding the contract method 0xf09748df.
//
// Solidity: function liquidateAndExecuteOrder(address trader, (uint256,address,int256,uint256,uint256,bool) order, uint256 liquidationAmount) returns()
func (_OrderBook *OrderBookTransactorSession) LiquidateAndExecuteOrder(trader common.Address, order IOrderBookOrder, liquidationAmount *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.LiquidateAndExecuteOrder(&_OrderBook.TransactOpts, trader, order, liquidationAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OrderBook *OrderBookTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OrderBook *OrderBookSession) Pause() (*types.Transaction, error) {
	return _OrderBook.Contract.Pause(&_OrderBook.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OrderBook *OrderBookTransactorSession) Pause() (*types.Transaction, error) {
	return _OrderBook.Contract.Pause(&_OrderBook.TransactOpts)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x4d39aa9f.
//
// Solidity: function placeOrder((uint256,address,int256,uint256,uint256,bool) order) returns()
func (_OrderBook *OrderBookTransactor) PlaceOrder(opts *bind.TransactOpts, order IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "placeOrder", order)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x4d39aa9f.
//
// Solidity: function placeOrder((uint256,address,int256,uint256,uint256,bool) order) returns()
func (_OrderBook *OrderBookSession) PlaceOrder(order IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.PlaceOrder(&_OrderBook.TransactOpts, order)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x4d39aa9f.
//
// Solidity: function placeOrder((uint256,address,int256,uint256,uint256,bool) order) returns()
func (_OrderBook *OrderBookTransactorSession) PlaceOrder(order IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.PlaceOrder(&_OrderBook.TransactOpts, order)
}

// PlaceOrders is a paid mutator transaction binding the contract method 0x82a42eb9.
//
// Solidity: function placeOrders((uint256,address,int256,uint256,uint256,bool)[] orders) returns()
func (_OrderBook *OrderBookTransactor) PlaceOrders(opts *bind.TransactOpts, orders []IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "placeOrders", orders)
}

// PlaceOrders is a paid mutator transaction binding the contract method 0x82a42eb9.
//
// Solidity: function placeOrders((uint256,address,int256,uint256,uint256,bool)[] orders) returns()
func (_OrderBook *OrderBookSession) PlaceOrders(orders []IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.PlaceOrders(&_OrderBook.TransactOpts, orders)
}

// PlaceOrders is a paid mutator transaction binding the contract method 0x82a42eb9.
//
// Solidity: function placeOrders((uint256,address,int256,uint256,uint256,bool)[] orders) returns()
func (_OrderBook *OrderBookTransactorSession) PlaceOrders(orders []IOrderBookOrder) (*types.Transaction, error) {
	return _OrderBook.Contract.PlaceOrders(&_OrderBook.TransactOpts, orders)
}

// SetBibliophile is a paid mutator transaction binding the contract method 0x450bc34c.
//
// Solidity: function setBibliophile(address _bibliophile) returns()
func (_OrderBook *OrderBookTransactor) SetBibliophile(opts *bind.TransactOpts, _bibliophile common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setBibliophile", _bibliophile)
}

// SetBibliophile is a paid mutator transaction binding the contract method 0x450bc34c.
//
// Solidity: function setBibliophile(address _bibliophile) returns()
func (_OrderBook *OrderBookSession) SetBibliophile(_bibliophile common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.SetBibliophile(&_OrderBook.TransactOpts, _bibliophile)
}

// SetBibliophile is a paid mutator transaction binding the contract method 0x450bc34c.
//
// Solidity: function setBibliophile(address _bibliophile) returns()
func (_OrderBook *OrderBookTransactorSession) SetBibliophile(_bibliophile common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.SetBibliophile(&_OrderBook.TransactOpts, _bibliophile)
}

// SetGovernace is a paid mutator transaction binding the contract method 0x851e3ca5.
//
// Solidity: function setGovernace(address __governance) returns()
func (_OrderBook *OrderBookTransactor) SetGovernace(opts *bind.TransactOpts, __governance common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setGovernace", __governance)
}

// SetGovernace is a paid mutator transaction binding the contract method 0x851e3ca5.
//
// Solidity: function setGovernace(address __governance) returns()
func (_OrderBook *OrderBookSession) SetGovernace(__governance common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.SetGovernace(&_OrderBook.TransactOpts, __governance)
}

// SetGovernace is a paid mutator transaction binding the contract method 0x851e3ca5.
//
// Solidity: function setGovernace(address __governance) returns()
func (_OrderBook *OrderBookTransactorSession) SetGovernace(__governance common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.SetGovernace(&_OrderBook.TransactOpts, __governance)
}

// SetUseNewPricingAlgorithm is a paid mutator transaction binding the contract method 0xb189830e.
//
// Solidity: function setUseNewPricingAlgorithm(bool useNew) returns()
func (_OrderBook *OrderBookTransactor) SetUseNewPricingAlgorithm(opts *bind.TransactOpts, useNew bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setUseNewPricingAlgorithm", useNew)
}

// SetUseNewPricingAlgorithm is a paid mutator transaction binding the contract method 0xb189830e.
//
// Solidity: function setUseNewPricingAlgorithm(bool useNew) returns()
func (_OrderBook *OrderBookSession) SetUseNewPricingAlgorithm(useNew bool) (*types.Transaction, error) {
	return _OrderBook.Contract.SetUseNewPricingAlgorithm(&_OrderBook.TransactOpts, useNew)
}

// SetUseNewPricingAlgorithm is a paid mutator transaction binding the contract method 0xb189830e.
//
// Solidity: function setUseNewPricingAlgorithm(bool useNew) returns()
func (_OrderBook *OrderBookTransactorSession) SetUseNewPricingAlgorithm(useNew bool) (*types.Transaction, error) {
	return _OrderBook.Contract.SetUseNewPricingAlgorithm(&_OrderBook.TransactOpts, useNew)
}

// SetValidatorStatus is a paid mutator transaction binding the contract method 0xaca2490b.
//
// Solidity: function setValidatorStatus(address validator, bool status) returns()
func (_OrderBook *OrderBookTransactor) SetValidatorStatus(opts *bind.TransactOpts, validator common.Address, status bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setValidatorStatus", validator, status)
}

// SetValidatorStatus is a paid mutator transaction binding the contract method 0xaca2490b.
//
// Solidity: function setValidatorStatus(address validator, bool status) returns()
func (_OrderBook *OrderBookSession) SetValidatorStatus(validator common.Address, status bool) (*types.Transaction, error) {
	return _OrderBook.Contract.SetValidatorStatus(&_OrderBook.TransactOpts, validator, status)
}

// SetValidatorStatus is a paid mutator transaction binding the contract method 0xaca2490b.
//
// Solidity: function setValidatorStatus(address validator, bool status) returns()
func (_OrderBook *OrderBookTransactorSession) SetValidatorStatus(validator common.Address, status bool) (*types.Transaction, error) {
	return _OrderBook.Contract.SetValidatorStatus(&_OrderBook.TransactOpts, validator, status)
}

// SettleFunding is a paid mutator transaction binding the contract method 0xed83d79c.
//
// Solidity: function settleFunding() returns()
func (_OrderBook *OrderBookTransactor) SettleFunding(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "settleFunding")
}

// SettleFunding is a paid mutator transaction binding the contract method 0xed83d79c.
//
// Solidity: function settleFunding() returns()
func (_OrderBook *OrderBookSession) SettleFunding() (*types.Transaction, error) {
	return _OrderBook.Contract.SettleFunding(&_OrderBook.TransactOpts)
}

// SettleFunding is a paid mutator transaction binding the contract method 0xed83d79c.
//
// Solidity: function settleFunding() returns()
func (_OrderBook *OrderBookTransactorSession) SettleFunding() (*types.Transaction, error) {
	return _OrderBook.Contract.SettleFunding(&_OrderBook.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OrderBook *OrderBookTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OrderBook *OrderBookSession) Unpause() (*types.Transaction, error) {
	return _OrderBook.Contract.Unpause(&_OrderBook.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OrderBook *OrderBookTransactorSession) Unpause() (*types.Transaction, error) {
	return _OrderBook.Contract.Unpause(&_OrderBook.TransactOpts)
}

// UpdateMinSize is a paid mutator transaction binding the contract method 0x5548e0e9.
//
// Solidity: function updateMinSize(uint256 ammIndex, int256 minSize) returns()
func (_OrderBook *OrderBookTransactor) UpdateMinSize(opts *bind.TransactOpts, ammIndex *big.Int, minSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "updateMinSize", ammIndex, minSize)
}

// UpdateMinSize is a paid mutator transaction binding the contract method 0x5548e0e9.
//
// Solidity: function updateMinSize(uint256 ammIndex, int256 minSize) returns()
func (_OrderBook *OrderBookSession) UpdateMinSize(ammIndex *big.Int, minSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateMinSize(&_OrderBook.TransactOpts, ammIndex, minSize)
}

// UpdateMinSize is a paid mutator transaction binding the contract method 0x5548e0e9.
//
// Solidity: function updateMinSize(uint256 ammIndex, int256 minSize) returns()
func (_OrderBook *OrderBookTransactorSession) UpdateMinSize(ammIndex *big.Int, minSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateMinSize(&_OrderBook.TransactOpts, ammIndex, minSize)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xb649d125.
//
// Solidity: function updateParams(uint256 _minAllowableMargin, uint256 _takerFee) returns()
func (_OrderBook *OrderBookTransactor) UpdateParams(opts *bind.TransactOpts, _minAllowableMargin *big.Int, _takerFee *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "updateParams", _minAllowableMargin, _takerFee)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xb649d125.
//
// Solidity: function updateParams(uint256 _minAllowableMargin, uint256 _takerFee) returns()
func (_OrderBook *OrderBookSession) UpdateParams(_minAllowableMargin *big.Int, _takerFee *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateParams(&_OrderBook.TransactOpts, _minAllowableMargin, _takerFee)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xb649d125.
//
// Solidity: function updateParams(uint256 _minAllowableMargin, uint256 _takerFee) returns()
func (_OrderBook *OrderBookTransactorSession) UpdateParams(_minAllowableMargin *big.Int, _takerFee *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateParams(&_OrderBook.TransactOpts, _minAllowableMargin, _takerFee)
}

// OrderBookInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OrderBook contract.
type OrderBookInitializedIterator struct {
	Event *OrderBookInitialized // Event containing the contract specifics and raw log

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
func (it *OrderBookInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookInitialized)
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
		it.Event = new(OrderBookInitialized)
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
func (it *OrderBookInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookInitialized represents a Initialized event raised by the OrderBook contract.
type OrderBookInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OrderBook *OrderBookFilterer) FilterInitialized(opts *bind.FilterOpts) (*OrderBookInitializedIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OrderBookInitializedIterator{contract: _OrderBook.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OrderBook *OrderBookFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OrderBookInitialized) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookInitialized)
				if err := _OrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OrderBook *OrderBookFilterer) ParseInitialized(log types.Log) (*OrderBookInitialized, error) {
	event := new(OrderBookInitialized)
	if err := _OrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookLiquidationErrorIterator is returned from FilterLiquidationError and is used to iterate over the raw logs and unpacked data for LiquidationError events raised by the OrderBook contract.
type OrderBookLiquidationErrorIterator struct {
	Event *OrderBookLiquidationError // Event containing the contract specifics and raw log

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
func (it *OrderBookLiquidationErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookLiquidationError)
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
		it.Event = new(OrderBookLiquidationError)
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
func (it *OrderBookLiquidationErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookLiquidationErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookLiquidationError represents a LiquidationError event raised by the OrderBook contract.
type OrderBookLiquidationError struct {
	Trader      common.Address
	OrderHash   [32]byte
	Err         string
	ToLiquidate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidationError is a free log retrieval operation binding the contract event 0x647843a5f461d447abbbb5434773da0b7bc49eda6fd5e50022efb4ebedfddf88.
//
// Solidity: event LiquidationError(address indexed trader, bytes32 indexed orderHash, string err, uint256 toLiquidate)
func (_OrderBook *OrderBookFilterer) FilterLiquidationError(opts *bind.FilterOpts, trader []common.Address, orderHash [][32]byte) (*OrderBookLiquidationErrorIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "LiquidationError", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookLiquidationErrorIterator{contract: _OrderBook.contract, event: "LiquidationError", logs: logs, sub: sub}, nil
}

// WatchLiquidationError is a free log subscription operation binding the contract event 0x647843a5f461d447abbbb5434773da0b7bc49eda6fd5e50022efb4ebedfddf88.
//
// Solidity: event LiquidationError(address indexed trader, bytes32 indexed orderHash, string err, uint256 toLiquidate)
func (_OrderBook *OrderBookFilterer) WatchLiquidationError(opts *bind.WatchOpts, sink chan<- *OrderBookLiquidationError, trader []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "LiquidationError", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookLiquidationError)
				if err := _OrderBook.contract.UnpackLog(event, "LiquidationError", log); err != nil {
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

// ParseLiquidationError is a log parse operation binding the contract event 0x647843a5f461d447abbbb5434773da0b7bc49eda6fd5e50022efb4ebedfddf88.
//
// Solidity: event LiquidationError(address indexed trader, bytes32 indexed orderHash, string err, uint256 toLiquidate)
func (_OrderBook *OrderBookFilterer) ParseLiquidationError(log types.Log) (*OrderBookLiquidationError, error) {
	event := new(OrderBookLiquidationError)
	if err := _OrderBook.contract.UnpackLog(event, "LiquidationError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookLiquidationOrderMatchedIterator is returned from FilterLiquidationOrderMatched and is used to iterate over the raw logs and unpacked data for LiquidationOrderMatched events raised by the OrderBook contract.
type OrderBookLiquidationOrderMatchedIterator struct {
	Event *OrderBookLiquidationOrderMatched // Event containing the contract specifics and raw log

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
func (it *OrderBookLiquidationOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookLiquidationOrderMatched)
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
		it.Event = new(OrderBookLiquidationOrderMatched)
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
func (it *OrderBookLiquidationOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookLiquidationOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookLiquidationOrderMatched represents a LiquidationOrderMatched event raised by the OrderBook contract.
type OrderBookLiquidationOrderMatched struct {
	Trader               common.Address
	OrderHash            [32]byte
	FillAmount           *big.Int
	Price                *big.Int
	OpenInterestNotional *big.Int
	Relayer              common.Address
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLiquidationOrderMatched is a free log retrieval operation binding the contract event 0x40fc27a15b65d8d0e06fc922be163e3826f702b164077a45b5d19fd9ce86bd90.
//
// Solidity: event LiquidationOrderMatched(address indexed trader, bytes32 indexed orderHash, uint256 fillAmount, uint256 price, uint256 openInterestNotional, address relayer, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) FilterLiquidationOrderMatched(opts *bind.FilterOpts, trader []common.Address, orderHash [][32]byte) (*OrderBookLiquidationOrderMatchedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "LiquidationOrderMatched", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookLiquidationOrderMatchedIterator{contract: _OrderBook.contract, event: "LiquidationOrderMatched", logs: logs, sub: sub}, nil
}

// WatchLiquidationOrderMatched is a free log subscription operation binding the contract event 0x40fc27a15b65d8d0e06fc922be163e3826f702b164077a45b5d19fd9ce86bd90.
//
// Solidity: event LiquidationOrderMatched(address indexed trader, bytes32 indexed orderHash, uint256 fillAmount, uint256 price, uint256 openInterestNotional, address relayer, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) WatchLiquidationOrderMatched(opts *bind.WatchOpts, sink chan<- *OrderBookLiquidationOrderMatched, trader []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "LiquidationOrderMatched", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookLiquidationOrderMatched)
				if err := _OrderBook.contract.UnpackLog(event, "LiquidationOrderMatched", log); err != nil {
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

// ParseLiquidationOrderMatched is a log parse operation binding the contract event 0x40fc27a15b65d8d0e06fc922be163e3826f702b164077a45b5d19fd9ce86bd90.
//
// Solidity: event LiquidationOrderMatched(address indexed trader, bytes32 indexed orderHash, uint256 fillAmount, uint256 price, uint256 openInterestNotional, address relayer, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) ParseLiquidationOrderMatched(log types.Log) (*OrderBookLiquidationOrderMatched, error) {
	event := new(OrderBookLiquidationOrderMatched)
	if err := _OrderBook.contract.UnpackLog(event, "LiquidationOrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the OrderBook contract.
type OrderBookOrderCancelledIterator struct {
	Event *OrderBookOrderCancelled // Event containing the contract specifics and raw log

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
func (it *OrderBookOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookOrderCancelled)
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
		it.Event = new(OrderBookOrderCancelled)
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
func (it *OrderBookOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookOrderCancelled represents a OrderCancelled event raised by the OrderBook contract.
type OrderBookOrderCancelled struct {
	Trader    common.Address
	OrderHash [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x26b214029d2b6a3a3bb2ae7cc0a5d4c9329a86381429e16dc45b3633cf83d369.
//
// Solidity: event OrderCancelled(address indexed trader, bytes32 indexed orderHash, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) FilterOrderCancelled(opts *bind.FilterOpts, trader []common.Address, orderHash [][32]byte) (*OrderBookOrderCancelledIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "OrderCancelled", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookOrderCancelledIterator{contract: _OrderBook.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x26b214029d2b6a3a3bb2ae7cc0a5d4c9329a86381429e16dc45b3633cf83d369.
//
// Solidity: event OrderCancelled(address indexed trader, bytes32 indexed orderHash, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *OrderBookOrderCancelled, trader []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "OrderCancelled", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookOrderCancelled)
				if err := _OrderBook.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x26b214029d2b6a3a3bb2ae7cc0a5d4c9329a86381429e16dc45b3633cf83d369.
//
// Solidity: event OrderCancelled(address indexed trader, bytes32 indexed orderHash, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) ParseOrderCancelled(log types.Log) (*OrderBookOrderCancelled, error) {
	event := new(OrderBookOrderCancelled)
	if err := _OrderBook.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookOrderMatchingErrorIterator is returned from FilterOrderMatchingError and is used to iterate over the raw logs and unpacked data for OrderMatchingError events raised by the OrderBook contract.
type OrderBookOrderMatchingErrorIterator struct {
	Event *OrderBookOrderMatchingError // Event containing the contract specifics and raw log

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
func (it *OrderBookOrderMatchingErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookOrderMatchingError)
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
		it.Event = new(OrderBookOrderMatchingError)
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
func (it *OrderBookOrderMatchingErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookOrderMatchingErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookOrderMatchingError represents a OrderMatchingError event raised by the OrderBook contract.
type OrderBookOrderMatchingError struct {
	OrderHash [32]byte
	Err       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderMatchingError is a free log retrieval operation binding the contract event 0xc8d4808dc7702e31368c7d7c8bc95e94f332759de3c6d76f8676668e55de25b0.
//
// Solidity: event OrderMatchingError(bytes32 indexed orderHash, string err)
func (_OrderBook *OrderBookFilterer) FilterOrderMatchingError(opts *bind.FilterOpts, orderHash [][32]byte) (*OrderBookOrderMatchingErrorIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "OrderMatchingError", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookOrderMatchingErrorIterator{contract: _OrderBook.contract, event: "OrderMatchingError", logs: logs, sub: sub}, nil
}

// WatchOrderMatchingError is a free log subscription operation binding the contract event 0xc8d4808dc7702e31368c7d7c8bc95e94f332759de3c6d76f8676668e55de25b0.
//
// Solidity: event OrderMatchingError(bytes32 indexed orderHash, string err)
func (_OrderBook *OrderBookFilterer) WatchOrderMatchingError(opts *bind.WatchOpts, sink chan<- *OrderBookOrderMatchingError, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "OrderMatchingError", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookOrderMatchingError)
				if err := _OrderBook.contract.UnpackLog(event, "OrderMatchingError", log); err != nil {
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

// ParseOrderMatchingError is a log parse operation binding the contract event 0xc8d4808dc7702e31368c7d7c8bc95e94f332759de3c6d76f8676668e55de25b0.
//
// Solidity: event OrderMatchingError(bytes32 indexed orderHash, string err)
func (_OrderBook *OrderBookFilterer) ParseOrderMatchingError(log types.Log) (*OrderBookOrderMatchingError, error) {
	event := new(OrderBookOrderMatchingError)
	if err := _OrderBook.contract.UnpackLog(event, "OrderMatchingError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the OrderBook contract.
type OrderBookOrderPlacedIterator struct {
	Event *OrderBookOrderPlaced // Event containing the contract specifics and raw log

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
func (it *OrderBookOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookOrderPlaced)
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
		it.Event = new(OrderBookOrderPlaced)
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
func (it *OrderBookOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookOrderPlaced represents a OrderPlaced event raised by the OrderBook contract.
type OrderBookOrderPlaced struct {
	Trader    common.Address
	OrderHash [32]byte
	Order     IOrderBookOrder
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0xfd027921ef87d77081c96b2b26a62c1512ee2652f0c049891faed86661570fbe.
//
// Solidity: event OrderPlaced(address indexed trader, bytes32 indexed orderHash, (uint256,address,int256,uint256,uint256,bool) order, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) FilterOrderPlaced(opts *bind.FilterOpts, trader []common.Address, orderHash [][32]byte) (*OrderBookOrderPlacedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "OrderPlaced", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookOrderPlacedIterator{contract: _OrderBook.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0xfd027921ef87d77081c96b2b26a62c1512ee2652f0c049891faed86661570fbe.
//
// Solidity: event OrderPlaced(address indexed trader, bytes32 indexed orderHash, (uint256,address,int256,uint256,uint256,bool) order, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *OrderBookOrderPlaced, trader []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "OrderPlaced", traderRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookOrderPlaced)
				if err := _OrderBook.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0xfd027921ef87d77081c96b2b26a62c1512ee2652f0c049891faed86661570fbe.
//
// Solidity: event OrderPlaced(address indexed trader, bytes32 indexed orderHash, (uint256,address,int256,uint256,uint256,bool) order, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) ParseOrderPlaced(log types.Log) (*OrderBookOrderPlaced, error) {
	event := new(OrderBookOrderPlaced)
	if err := _OrderBook.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the OrderBook contract.
type OrderBookOrdersMatchedIterator struct {
	Event *OrderBookOrdersMatched // Event containing the contract specifics and raw log

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
func (it *OrderBookOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookOrdersMatched)
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
		it.Event = new(OrderBookOrdersMatched)
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
func (it *OrderBookOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookOrdersMatched represents a OrdersMatched event raised by the OrderBook contract.
type OrderBookOrdersMatched struct {
	OrderHash0           [32]byte
	OrderHash1           [32]byte
	FillAmount           *big.Int
	Price                *big.Int
	OpenInterestNotional *big.Int
	Relayer              common.Address
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0xaf4b403d9952e032974b549a4abad80faca307b0acc6e34d7e0b8c274d504590.
//
// Solidity: event OrdersMatched(bytes32 indexed orderHash0, bytes32 indexed orderHash1, uint256 fillAmount, uint256 price, uint256 openInterestNotional, address relayer, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) FilterOrdersMatched(opts *bind.FilterOpts, orderHash0 [][32]byte, orderHash1 [][32]byte) (*OrderBookOrdersMatchedIterator, error) {

	var orderHash0Rule []interface{}
	for _, orderHash0Item := range orderHash0 {
		orderHash0Rule = append(orderHash0Rule, orderHash0Item)
	}
	var orderHash1Rule []interface{}
	for _, orderHash1Item := range orderHash1 {
		orderHash1Rule = append(orderHash1Rule, orderHash1Item)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "OrdersMatched", orderHash0Rule, orderHash1Rule)
	if err != nil {
		return nil, err
	}
	return &OrderBookOrdersMatchedIterator{contract: _OrderBook.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0xaf4b403d9952e032974b549a4abad80faca307b0acc6e34d7e0b8c274d504590.
//
// Solidity: event OrdersMatched(bytes32 indexed orderHash0, bytes32 indexed orderHash1, uint256 fillAmount, uint256 price, uint256 openInterestNotional, address relayer, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *OrderBookOrdersMatched, orderHash0 [][32]byte, orderHash1 [][32]byte) (event.Subscription, error) {

	var orderHash0Rule []interface{}
	for _, orderHash0Item := range orderHash0 {
		orderHash0Rule = append(orderHash0Rule, orderHash0Item)
	}
	var orderHash1Rule []interface{}
	for _, orderHash1Item := range orderHash1 {
		orderHash1Rule = append(orderHash1Rule, orderHash1Item)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "OrdersMatched", orderHash0Rule, orderHash1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookOrdersMatched)
				if err := _OrderBook.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0xaf4b403d9952e032974b549a4abad80faca307b0acc6e34d7e0b8c274d504590.
//
// Solidity: event OrdersMatched(bytes32 indexed orderHash0, bytes32 indexed orderHash1, uint256 fillAmount, uint256 price, uint256 openInterestNotional, address relayer, uint256 timestamp)
func (_OrderBook *OrderBookFilterer) ParseOrdersMatched(log types.Log) (*OrderBookOrdersMatched, error) {
	event := new(OrderBookOrdersMatched)
	if err := _OrderBook.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OrderBook contract.
type OrderBookPausedIterator struct {
	Event *OrderBookPaused // Event containing the contract specifics and raw log

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
func (it *OrderBookPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookPaused)
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
		it.Event = new(OrderBookPaused)
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
func (it *OrderBookPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookPaused represents a Paused event raised by the OrderBook contract.
type OrderBookPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OrderBook *OrderBookFilterer) FilterPaused(opts *bind.FilterOpts) (*OrderBookPausedIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OrderBookPausedIterator{contract: _OrderBook.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OrderBook *OrderBookFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OrderBookPaused) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookPaused)
				if err := _OrderBook.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OrderBook *OrderBookFilterer) ParsePaused(log types.Log) (*OrderBookPaused, error) {
	event := new(OrderBookPaused)
	if err := _OrderBook.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OrderBook contract.
type OrderBookUnpausedIterator struct {
	Event *OrderBookUnpaused // Event containing the contract specifics and raw log

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
func (it *OrderBookUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUnpaused)
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
		it.Event = new(OrderBookUnpaused)
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
func (it *OrderBookUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUnpaused represents a Unpaused event raised by the OrderBook contract.
type OrderBookUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OrderBook *OrderBookFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OrderBookUnpausedIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OrderBookUnpausedIterator{contract: _OrderBook.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OrderBook *OrderBookFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OrderBookUnpaused) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUnpaused)
				if err := _OrderBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OrderBook *OrderBookFilterer) ParseUnpaused(log types.Log) (*OrderBookUnpaused, error) {
	event := new(OrderBookUnpaused)
	if err := _OrderBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
