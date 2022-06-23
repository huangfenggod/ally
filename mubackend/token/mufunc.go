// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// MufuncMetaData contains all meta data concerning the Mufunc contract.
var MufuncMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ally\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pancakerouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autotrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceBNB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfTokenContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"changeAutoTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"changeFundAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tranferBNB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFromNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFromU\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount2\",\"type\":\"uint256\"}],\"name\":\"transferToCustomer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferTokenToOhters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUsdt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUsdtFromCustomer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"usdtApproveTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MufuncABI is the input ABI used to generate the binding from.
// Deprecated: Use MufuncMetaData.ABI instead.
var MufuncABI = MufuncMetaData.ABI

// Mufunc is an auto generated Go binding around an Ethereum contract.
type Mufunc struct {
	MufuncCaller     // Read-only binding to the contract
	MufuncTransactor // Write-only binding to the contract
	MufuncFilterer   // Log filterer for contract events
}

// MufuncCaller is an auto generated read-only Go binding around an Ethereum contract.
type MufuncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MufuncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MufuncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MufuncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MufuncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MufuncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MufuncSession struct {
	Contract     *Mufunc           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MufuncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MufuncCallerSession struct {
	Contract *MufuncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MufuncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MufuncTransactorSession struct {
	Contract     *MufuncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MufuncRaw is an auto generated low-level Go binding around an Ethereum contract.
type MufuncRaw struct {
	Contract *Mufunc // Generic contract binding to access the raw methods on
}

// MufuncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MufuncCallerRaw struct {
	Contract *MufuncCaller // Generic read-only contract binding to access the raw methods on
}

// MufuncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MufuncTransactorRaw struct {
	Contract *MufuncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMufunc creates a new instance of Mufunc, bound to a specific deployed contract.
func NewMufunc(address common.Address, backend bind.ContractBackend) (*Mufunc, error) {
	contract, err := bindMufunc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mufunc{MufuncCaller: MufuncCaller{contract: contract}, MufuncTransactor: MufuncTransactor{contract: contract}, MufuncFilterer: MufuncFilterer{contract: contract}}, nil
}

// NewMufuncCaller creates a new read-only instance of Mufunc, bound to a specific deployed contract.
func NewMufuncCaller(address common.Address, caller bind.ContractCaller) (*MufuncCaller, error) {
	contract, err := bindMufunc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MufuncCaller{contract: contract}, nil
}

// NewMufuncTransactor creates a new write-only instance of Mufunc, bound to a specific deployed contract.
func NewMufuncTransactor(address common.Address, transactor bind.ContractTransactor) (*MufuncTransactor, error) {
	contract, err := bindMufunc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MufuncTransactor{contract: contract}, nil
}

// NewMufuncFilterer creates a new log filterer instance of Mufunc, bound to a specific deployed contract.
func NewMufuncFilterer(address common.Address, filterer bind.ContractFilterer) (*MufuncFilterer, error) {
	contract, err := bindMufunc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MufuncFilterer{contract: contract}, nil
}

// bindMufunc binds a generic wrapper to an already deployed contract.
func bindMufunc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MufuncABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mufunc *MufuncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mufunc.Contract.MufuncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mufunc *MufuncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mufunc.Contract.MufuncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mufunc *MufuncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mufunc.Contract.MufuncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mufunc *MufuncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mufunc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mufunc *MufuncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mufunc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mufunc *MufuncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mufunc.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Mufunc *MufuncCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Mufunc *MufuncSession) Owner() (common.Address, error) {
	return _Mufunc.Contract.Owner(&_Mufunc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Mufunc *MufuncCallerSession) Owner() (common.Address, error) {
	return _Mufunc.Contract.Owner(&_Mufunc.CallOpts)
}

// Autotrade is a free data retrieval call binding the contract method 0x484ea02b.
//
// Solidity: function autotrade() view returns(bool)
func (_Mufunc *MufuncCaller) Autotrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "autotrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Autotrade is a free data retrieval call binding the contract method 0x484ea02b.
//
// Solidity: function autotrade() view returns(bool)
func (_Mufunc *MufuncSession) Autotrade() (bool, error) {
	return _Mufunc.Contract.Autotrade(&_Mufunc.CallOpts)
}

// Autotrade is a free data retrieval call binding the contract method 0x484ea02b.
//
// Solidity: function autotrade() view returns(bool)
func (_Mufunc *MufuncCallerSession) Autotrade() (bool, error) {
	return _Mufunc.Contract.Autotrade(&_Mufunc.CallOpts)
}

// BalanceBNB is a free data retrieval call binding the contract method 0x0b04fffe.
//
// Solidity: function balanceBNB() view returns(uint256)
func (_Mufunc *MufuncCaller) BalanceBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "balanceBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceBNB is a free data retrieval call binding the contract method 0x0b04fffe.
//
// Solidity: function balanceBNB() view returns(uint256)
func (_Mufunc *MufuncSession) BalanceBNB() (*big.Int, error) {
	return _Mufunc.Contract.BalanceBNB(&_Mufunc.CallOpts)
}

// BalanceBNB is a free data retrieval call binding the contract method 0x0b04fffe.
//
// Solidity: function balanceBNB() view returns(uint256)
func (_Mufunc *MufuncCallerSession) BalanceBNB() (*big.Int, error) {
	return _Mufunc.Contract.BalanceBNB(&_Mufunc.CallOpts)
}

// BalanceOfTokenContract is a free data retrieval call binding the contract method 0xae0dfc0e.
//
// Solidity: function balanceOfTokenContract() view returns(uint256)
func (_Mufunc *MufuncCaller) BalanceOfTokenContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "balanceOfTokenContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfTokenContract is a free data retrieval call binding the contract method 0xae0dfc0e.
//
// Solidity: function balanceOfTokenContract() view returns(uint256)
func (_Mufunc *MufuncSession) BalanceOfTokenContract() (*big.Int, error) {
	return _Mufunc.Contract.BalanceOfTokenContract(&_Mufunc.CallOpts)
}

// BalanceOfTokenContract is a free data retrieval call binding the contract method 0xae0dfc0e.
//
// Solidity: function balanceOfTokenContract() view returns(uint256)
func (_Mufunc *MufuncCallerSession) BalanceOfTokenContract() (*big.Int, error) {
	return _Mufunc.Contract.BalanceOfTokenContract(&_Mufunc.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256, uint256)
func (_Mufunc *MufuncCaller) GetRate(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256, uint256)
func (_Mufunc *MufuncSession) GetRate() (*big.Int, *big.Int, error) {
	return _Mufunc.Contract.GetRate(&_Mufunc.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256, uint256)
func (_Mufunc *MufuncCallerSession) GetRate() (*big.Int, *big.Int, error) {
	return _Mufunc.Contract.GetRate(&_Mufunc.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Mufunc *MufuncCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Mufunc *MufuncSession) UniswapV2Router() (common.Address, error) {
	return _Mufunc.Contract.UniswapV2Router(&_Mufunc.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Mufunc *MufuncCallerSession) UniswapV2Router() (common.Address, error) {
	return _Mufunc.Contract.UniswapV2Router(&_Mufunc.CallOpts)
}

// UsdtNumber is a free data retrieval call binding the contract method 0x3b0546ad.
//
// Solidity: function usdtNumber() view returns(uint256)
func (_Mufunc *MufuncCaller) UsdtNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mufunc.contract.Call(opts, &out, "usdtNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdtNumber is a free data retrieval call binding the contract method 0x3b0546ad.
//
// Solidity: function usdtNumber() view returns(uint256)
func (_Mufunc *MufuncSession) UsdtNumber() (*big.Int, error) {
	return _Mufunc.Contract.UsdtNumber(&_Mufunc.CallOpts)
}

// UsdtNumber is a free data retrieval call binding the contract method 0x3b0546ad.
//
// Solidity: function usdtNumber() view returns(uint256)
func (_Mufunc *MufuncCallerSession) UsdtNumber() (*big.Int, error) {
	return _Mufunc.Contract.UsdtNumber(&_Mufunc.CallOpts)
}

// ChangeAutoTrade is a paid mutator transaction binding the contract method 0x52bd61bc.
//
// Solidity: function changeAutoTrade() returns(bool)
func (_Mufunc *MufuncTransactor) ChangeAutoTrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "changeAutoTrade")
}

// ChangeAutoTrade is a paid mutator transaction binding the contract method 0x52bd61bc.
//
// Solidity: function changeAutoTrade() returns(bool)
func (_Mufunc *MufuncSession) ChangeAutoTrade() (*types.Transaction, error) {
	return _Mufunc.Contract.ChangeAutoTrade(&_Mufunc.TransactOpts)
}

// ChangeAutoTrade is a paid mutator transaction binding the contract method 0x52bd61bc.
//
// Solidity: function changeAutoTrade() returns(bool)
func (_Mufunc *MufuncTransactorSession) ChangeAutoTrade() (*types.Transaction, error) {
	return _Mufunc.Contract.ChangeAutoTrade(&_Mufunc.TransactOpts)
}

// ChangeFundAccount is a paid mutator transaction binding the contract method 0x96918c91.
//
// Solidity: function changeFundAccount(address account) returns(bool)
func (_Mufunc *MufuncTransactor) ChangeFundAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "changeFundAccount", account)
}

// ChangeFundAccount is a paid mutator transaction binding the contract method 0x96918c91.
//
// Solidity: function changeFundAccount(address account) returns(bool)
func (_Mufunc *MufuncSession) ChangeFundAccount(account common.Address) (*types.Transaction, error) {
	return _Mufunc.Contract.ChangeFundAccount(&_Mufunc.TransactOpts, account)
}

// ChangeFundAccount is a paid mutator transaction binding the contract method 0x96918c91.
//
// Solidity: function changeFundAccount(address account) returns(bool)
func (_Mufunc *MufuncTransactorSession) ChangeFundAccount(account common.Address) (*types.Transaction, error) {
	return _Mufunc.Contract.ChangeFundAccount(&_Mufunc.TransactOpts, account)
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(bool)
func (_Mufunc *MufuncTransactor) GetValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "getValue")
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(bool)
func (_Mufunc *MufuncSession) GetValue() (*types.Transaction, error) {
	return _Mufunc.Contract.GetValue(&_Mufunc.TransactOpts)
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(bool)
func (_Mufunc *MufuncTransactorSession) GetValue() (*types.Transaction, error) {
	return _Mufunc.Contract.GetValue(&_Mufunc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mufunc *MufuncTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mufunc *MufuncSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mufunc.Contract.RenounceOwnership(&_Mufunc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mufunc *MufuncTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mufunc.Contract.RenounceOwnership(&_Mufunc.TransactOpts)
}

// TranferBNB is a paid mutator transaction binding the contract method 0x1634d079.
//
// Solidity: function tranferBNB(address to, uint256 amount) payable returns(bool)
func (_Mufunc *MufuncTransactor) TranferBNB(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "tranferBNB", to, amount)
}

// TranferBNB is a paid mutator transaction binding the contract method 0x1634d079.
//
// Solidity: function tranferBNB(address to, uint256 amount) payable returns(bool)
func (_Mufunc *MufuncSession) TranferBNB(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TranferBNB(&_Mufunc.TransactOpts, to, amount)
}

// TranferBNB is a paid mutator transaction binding the contract method 0x1634d079.
//
// Solidity: function tranferBNB(address to, uint256 amount) payable returns(bool)
func (_Mufunc *MufuncTransactorSession) TranferBNB(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TranferBNB(&_Mufunc.TransactOpts, to, amount)
}

// TransferFromNode is a paid mutator transaction binding the contract method 0x8a8521db.
//
// Solidity: function transferFromNode(address customer, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactor) TransferFromNode(opts *bind.TransactOpts, customer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferFromNode", customer, amount)
}

// TransferFromNode is a paid mutator transaction binding the contract method 0x8a8521db.
//
// Solidity: function transferFromNode(address customer, uint256 amount) returns(bool)
func (_Mufunc *MufuncSession) TransferFromNode(customer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferFromNode(&_Mufunc.TransactOpts, customer, amount)
}

// TransferFromNode is a paid mutator transaction binding the contract method 0x8a8521db.
//
// Solidity: function transferFromNode(address customer, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactorSession) TransferFromNode(customer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferFromNode(&_Mufunc.TransactOpts, customer, amount)
}

// TransferFromU is a paid mutator transaction binding the contract method 0xebbb990c.
//
// Solidity: function transferFromU(address account, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactor) TransferFromU(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferFromU", account, amount)
}

// TransferFromU is a paid mutator transaction binding the contract method 0xebbb990c.
//
// Solidity: function transferFromU(address account, uint256 amount) returns(bool)
func (_Mufunc *MufuncSession) TransferFromU(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferFromU(&_Mufunc.TransactOpts, account, amount)
}

// TransferFromU is a paid mutator transaction binding the contract method 0xebbb990c.
//
// Solidity: function transferFromU(address account, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactorSession) TransferFromU(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferFromU(&_Mufunc.TransactOpts, account, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mufunc *MufuncTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mufunc *MufuncSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferOwnership(&_Mufunc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mufunc *MufuncTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferOwnership(&_Mufunc.TransactOpts, newOwner)
}

// TransferToCustomer is a paid mutator transaction binding the contract method 0x0c501331.
//
// Solidity: function transferToCustomer(address customer, uint256 amount1, uint256 amount2) returns(bool)
func (_Mufunc *MufuncTransactor) TransferToCustomer(opts *bind.TransactOpts, customer common.Address, amount1 *big.Int, amount2 *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferToCustomer", customer, amount1, amount2)
}

// TransferToCustomer is a paid mutator transaction binding the contract method 0x0c501331.
//
// Solidity: function transferToCustomer(address customer, uint256 amount1, uint256 amount2) returns(bool)
func (_Mufunc *MufuncSession) TransferToCustomer(customer common.Address, amount1 *big.Int, amount2 *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferToCustomer(&_Mufunc.TransactOpts, customer, amount1, amount2)
}

// TransferToCustomer is a paid mutator transaction binding the contract method 0x0c501331.
//
// Solidity: function transferToCustomer(address customer, uint256 amount1, uint256 amount2) returns(bool)
func (_Mufunc *MufuncTransactorSession) TransferToCustomer(customer common.Address, amount1 *big.Int, amount2 *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferToCustomer(&_Mufunc.TransactOpts, customer, amount1, amount2)
}

// TransferTokenToOhters is a paid mutator transaction binding the contract method 0x3b39321e.
//
// Solidity: function transferTokenToOhters(uint256 amount, address account) returns()
func (_Mufunc *MufuncTransactor) TransferTokenToOhters(opts *bind.TransactOpts, amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferTokenToOhters", amount, account)
}

// TransferTokenToOhters is a paid mutator transaction binding the contract method 0x3b39321e.
//
// Solidity: function transferTokenToOhters(uint256 amount, address account) returns()
func (_Mufunc *MufuncSession) TransferTokenToOhters(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferTokenToOhters(&_Mufunc.TransactOpts, amount, account)
}

// TransferTokenToOhters is a paid mutator transaction binding the contract method 0x3b39321e.
//
// Solidity: function transferTokenToOhters(uint256 amount, address account) returns()
func (_Mufunc *MufuncTransactorSession) TransferTokenToOhters(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferTokenToOhters(&_Mufunc.TransactOpts, amount, account)
}

// TransferUsdt is a paid mutator transaction binding the contract method 0xa30578a1.
//
// Solidity: function transferUsdt(address to, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactor) TransferUsdt(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferUsdt", to, amount)
}

// TransferUsdt is a paid mutator transaction binding the contract method 0xa30578a1.
//
// Solidity: function transferUsdt(address to, uint256 amount) returns(bool)
func (_Mufunc *MufuncSession) TransferUsdt(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferUsdt(&_Mufunc.TransactOpts, to, amount)
}

// TransferUsdt is a paid mutator transaction binding the contract method 0xa30578a1.
//
// Solidity: function transferUsdt(address to, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactorSession) TransferUsdt(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferUsdt(&_Mufunc.TransactOpts, to, amount)
}

// TransferUsdtFromCustomer is a paid mutator transaction binding the contract method 0xf49b034f.
//
// Solidity: function transferUsdtFromCustomer(address customer, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactor) TransferUsdtFromCustomer(opts *bind.TransactOpts, customer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "transferUsdtFromCustomer", customer, amount)
}

// TransferUsdtFromCustomer is a paid mutator transaction binding the contract method 0xf49b034f.
//
// Solidity: function transferUsdtFromCustomer(address customer, uint256 amount) returns(bool)
func (_Mufunc *MufuncSession) TransferUsdtFromCustomer(customer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferUsdtFromCustomer(&_Mufunc.TransactOpts, customer, amount)
}

// TransferUsdtFromCustomer is a paid mutator transaction binding the contract method 0xf49b034f.
//
// Solidity: function transferUsdtFromCustomer(address customer, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactorSession) TransferUsdtFromCustomer(customer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.TransferUsdtFromCustomer(&_Mufunc.TransactOpts, customer, amount)
}

// UsdtApproveTo is a paid mutator transaction binding the contract method 0x2c434cac.
//
// Solidity: function usdtApproveTo(address to, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactor) UsdtApproveTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.contract.Transact(opts, "usdtApproveTo", to, amount)
}

// UsdtApproveTo is a paid mutator transaction binding the contract method 0x2c434cac.
//
// Solidity: function usdtApproveTo(address to, uint256 amount) returns(bool)
func (_Mufunc *MufuncSession) UsdtApproveTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.UsdtApproveTo(&_Mufunc.TransactOpts, to, amount)
}

// UsdtApproveTo is a paid mutator transaction binding the contract method 0x2c434cac.
//
// Solidity: function usdtApproveTo(address to, uint256 amount) returns(bool)
func (_Mufunc *MufuncTransactorSession) UsdtApproveTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mufunc.Contract.UsdtApproveTo(&_Mufunc.TransactOpts, to, amount)
}
