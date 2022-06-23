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

// MubsMetaData contains all meta data concerning the Mubs contract.
var MubsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pancake\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"lock\",\"type\":\"bool\"}],\"name\":\"OpenLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fundAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_operateAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBrunList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBrunSenderList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"changeExchangeLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operateAccount\",\"type\":\"address\"}],\"name\":\"changeOperate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeA2U\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBrunList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBrunSenderList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalcaiwu\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MubsABI is the input ABI used to generate the binding from.
// Deprecated: Use MubsMetaData.ABI instead.
var MubsABI = MubsMetaData.ABI

// Mubs is an auto generated Go binding around an Ethereum contract.
type Mubs struct {
	MubsCaller     // Read-only binding to the contract
	MubsTransactor // Write-only binding to the contract
	MubsFilterer   // Log filterer for contract events
}

// MubsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MubsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MubsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MubsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MubsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MubsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MubsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MubsSession struct {
	Contract     *Mubs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MubsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MubsCallerSession struct {
	Contract *MubsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MubsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MubsTransactorSession struct {
	Contract     *MubsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MubsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MubsRaw struct {
	Contract *Mubs // Generic contract binding to access the raw methods on
}

// MubsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MubsCallerRaw struct {
	Contract *MubsCaller // Generic read-only contract binding to access the raw methods on
}

// MubsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MubsTransactorRaw struct {
	Contract *MubsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMubs creates a new instance of Mubs, bound to a specific deployed contract.
func NewMubs(address common.Address, backend bind.ContractBackend) (*Mubs, error) {
	contract, err := bindMubs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mubs{MubsCaller: MubsCaller{contract: contract}, MubsTransactor: MubsTransactor{contract: contract}, MubsFilterer: MubsFilterer{contract: contract}}, nil
}

// NewMubsCaller creates a new read-only instance of Mubs, bound to a specific deployed contract.
func NewMubsCaller(address common.Address, caller bind.ContractCaller) (*MubsCaller, error) {
	contract, err := bindMubs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MubsCaller{contract: contract}, nil
}

// NewMubsTransactor creates a new write-only instance of Mubs, bound to a specific deployed contract.
func NewMubsTransactor(address common.Address, transactor bind.ContractTransactor) (*MubsTransactor, error) {
	contract, err := bindMubs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MubsTransactor{contract: contract}, nil
}

// NewMubsFilterer creates a new log filterer instance of Mubs, bound to a specific deployed contract.
func NewMubsFilterer(address common.Address, filterer bind.ContractFilterer) (*MubsFilterer, error) {
	contract, err := bindMubs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MubsFilterer{contract: contract}, nil
}

// bindMubs binds a generic wrapper to an already deployed contract.
func bindMubs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MubsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mubs *MubsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mubs.Contract.MubsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mubs *MubsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mubs.Contract.MubsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mubs *MubsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mubs.Contract.MubsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mubs *MubsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mubs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mubs *MubsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mubs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mubs *MubsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mubs.Contract.contract.Transact(opts, method, params...)
}

// FundAccount is a free data retrieval call binding the contract method 0xafecdee8.
//
// Solidity: function _fundAccount() view returns(address)
func (_Mubs *MubsCaller) FundAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "_fundAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundAccount is a free data retrieval call binding the contract method 0xafecdee8.
//
// Solidity: function _fundAccount() view returns(address)
func (_Mubs *MubsSession) FundAccount() (common.Address, error) {
	return _Mubs.Contract.FundAccount(&_Mubs.CallOpts)
}

// FundAccount is a free data retrieval call binding the contract method 0xafecdee8.
//
// Solidity: function _fundAccount() view returns(address)
func (_Mubs *MubsCallerSession) FundAccount() (common.Address, error) {
	return _Mubs.Contract.FundAccount(&_Mubs.CallOpts)
}

// OperateAccount is a free data retrieval call binding the contract method 0xe0166ef5.
//
// Solidity: function _operateAccount() view returns(address)
func (_Mubs *MubsCaller) OperateAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "_operateAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperateAccount is a free data retrieval call binding the contract method 0xe0166ef5.
//
// Solidity: function _operateAccount() view returns(address)
func (_Mubs *MubsSession) OperateAccount() (common.Address, error) {
	return _Mubs.Contract.OperateAccount(&_Mubs.CallOpts)
}

// OperateAccount is a free data retrieval call binding the contract method 0xe0166ef5.
//
// Solidity: function _operateAccount() view returns(address)
func (_Mubs *MubsCallerSession) OperateAccount() (common.Address, error) {
	return _Mubs.Contract.OperateAccount(&_Mubs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Mubs *MubsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Mubs *MubsSession) Owner() (common.Address, error) {
	return _Mubs.Contract.Owner(&_Mubs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Mubs *MubsCallerSession) Owner() (common.Address, error) {
	return _Mubs.Contract.Owner(&_Mubs.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mubs *MubsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mubs *MubsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mubs.Contract.Allowance(&_Mubs.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mubs *MubsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mubs.Contract.Allowance(&_Mubs.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mubs *MubsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mubs *MubsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mubs.Contract.BalanceOf(&_Mubs.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mubs *MubsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mubs.Contract.BalanceOf(&_Mubs.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mubs *MubsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mubs *MubsSession) Decimals() (uint8, error) {
	return _Mubs.Contract.Decimals(&_Mubs.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mubs *MubsCallerSession) Decimals() (uint8, error) {
	return _Mubs.Contract.Decimals(&_Mubs.CallOpts)
}

// ExchangeLock is a free data retrieval call binding the contract method 0xfed34be1.
//
// Solidity: function exchangeLock() view returns(bool)
func (_Mubs *MubsCaller) ExchangeLock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "exchangeLock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExchangeLock is a free data retrieval call binding the contract method 0xfed34be1.
//
// Solidity: function exchangeLock() view returns(bool)
func (_Mubs *MubsSession) ExchangeLock() (bool, error) {
	return _Mubs.Contract.ExchangeLock(&_Mubs.CallOpts)
}

// ExchangeLock is a free data retrieval call binding the contract method 0xfed34be1.
//
// Solidity: function exchangeLock() view returns(bool)
func (_Mubs *MubsCallerSession) ExchangeLock() (bool, error) {
	return _Mubs.Contract.ExchangeLock(&_Mubs.CallOpts)
}

// GetFundAmount is a free data retrieval call binding the contract method 0xde7d282a.
//
// Solidity: function getFundAmount() view returns(uint256)
func (_Mubs *MubsCaller) GetFundAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "getFundAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundAmount is a free data retrieval call binding the contract method 0xde7d282a.
//
// Solidity: function getFundAmount() view returns(uint256)
func (_Mubs *MubsSession) GetFundAmount() (*big.Int, error) {
	return _Mubs.Contract.GetFundAmount(&_Mubs.CallOpts)
}

// GetFundAmount is a free data retrieval call binding the contract method 0xde7d282a.
//
// Solidity: function getFundAmount() view returns(uint256)
func (_Mubs *MubsCallerSession) GetFundAmount() (*big.Int, error) {
	return _Mubs.Contract.GetFundAmount(&_Mubs.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint112, uint112)
func (_Mubs *MubsCaller) GetPrice(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint112, uint112)
func (_Mubs *MubsSession) GetPrice() (*big.Int, *big.Int, error) {
	return _Mubs.Contract.GetPrice(&_Mubs.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint112, uint112)
func (_Mubs *MubsCallerSession) GetPrice() (*big.Int, *big.Int, error) {
	return _Mubs.Contract.GetPrice(&_Mubs.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mubs *MubsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mubs *MubsSession) Name() (string, error) {
	return _Mubs.Contract.Name(&_Mubs.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mubs *MubsCallerSession) Name() (string, error) {
	return _Mubs.Contract.Name(&_Mubs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mubs *MubsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mubs *MubsSession) Symbol() (string, error) {
	return _Mubs.Contract.Symbol(&_Mubs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mubs *MubsCallerSession) Symbol() (string, error) {
	return _Mubs.Contract.Symbol(&_Mubs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mubs *MubsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mubs *MubsSession) TotalSupply() (*big.Int, error) {
	return _Mubs.Contract.TotalSupply(&_Mubs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mubs *MubsCallerSession) TotalSupply() (*big.Int, error) {
	return _Mubs.Contract.TotalSupply(&_Mubs.CallOpts)
}

// Totalcaiwu is a free data retrieval call binding the contract method 0xbcbbf59a.
//
// Solidity: function totalcaiwu() view returns(address)
func (_Mubs *MubsCaller) Totalcaiwu(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "totalcaiwu")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Totalcaiwu is a free data retrieval call binding the contract method 0xbcbbf59a.
//
// Solidity: function totalcaiwu() view returns(address)
func (_Mubs *MubsSession) Totalcaiwu() (common.Address, error) {
	return _Mubs.Contract.Totalcaiwu(&_Mubs.CallOpts)
}

// Totalcaiwu is a free data retrieval call binding the contract method 0xbcbbf59a.
//
// Solidity: function totalcaiwu() view returns(address)
func (_Mubs *MubsCallerSession) Totalcaiwu() (common.Address, error) {
	return _Mubs.Contract.Totalcaiwu(&_Mubs.CallOpts)
}

// USender is a free data retrieval call binding the contract method 0xef3c12c5.
//
// Solidity: function uSender() view returns(address)
func (_Mubs *MubsCaller) USender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "uSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USender is a free data retrieval call binding the contract method 0xef3c12c5.
//
// Solidity: function uSender() view returns(address)
func (_Mubs *MubsSession) USender() (common.Address, error) {
	return _Mubs.Contract.USender(&_Mubs.CallOpts)
}

// USender is a free data retrieval call binding the contract method 0xef3c12c5.
//
// Solidity: function uSender() view returns(address)
func (_Mubs *MubsCallerSession) USender() (common.Address, error) {
	return _Mubs.Contract.USender(&_Mubs.CallOpts)
}

// Unlock is a free data retrieval call binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() view returns(bool)
func (_Mubs *MubsCaller) Unlock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mubs.contract.Call(opts, &out, "unlock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Unlock is a free data retrieval call binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() view returns(bool)
func (_Mubs *MubsSession) Unlock() (bool, error) {
	return _Mubs.Contract.Unlock(&_Mubs.CallOpts)
}

// Unlock is a free data retrieval call binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() view returns(bool)
func (_Mubs *MubsCallerSession) Unlock() (bool, error) {
	return _Mubs.Contract.Unlock(&_Mubs.CallOpts)
}

// OpenLock is a paid mutator transaction binding the contract method 0xa57824b5.
//
// Solidity: function OpenLock(bool lock) returns(bool)
func (_Mubs *MubsTransactor) OpenLock(opts *bind.TransactOpts, lock bool) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "OpenLock", lock)
}

// OpenLock is a paid mutator transaction binding the contract method 0xa57824b5.
//
// Solidity: function OpenLock(bool lock) returns(bool)
func (_Mubs *MubsSession) OpenLock(lock bool) (*types.Transaction, error) {
	return _Mubs.Contract.OpenLock(&_Mubs.TransactOpts, lock)
}

// OpenLock is a paid mutator transaction binding the contract method 0xa57824b5.
//
// Solidity: function OpenLock(bool lock) returns(bool)
func (_Mubs *MubsTransactorSession) OpenLock(lock bool) (*types.Transaction, error) {
	return _Mubs.Contract.OpenLock(&_Mubs.TransactOpts, lock)
}

// AddBrunList is a paid mutator transaction binding the contract method 0xb30ad11c.
//
// Solidity: function addBrunList(address account) returns(bool)
func (_Mubs *MubsTransactor) AddBrunList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "addBrunList", account)
}

// AddBrunList is a paid mutator transaction binding the contract method 0xb30ad11c.
//
// Solidity: function addBrunList(address account) returns(bool)
func (_Mubs *MubsSession) AddBrunList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.AddBrunList(&_Mubs.TransactOpts, account)
}

// AddBrunList is a paid mutator transaction binding the contract method 0xb30ad11c.
//
// Solidity: function addBrunList(address account) returns(bool)
func (_Mubs *MubsTransactorSession) AddBrunList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.AddBrunList(&_Mubs.TransactOpts, account)
}

// AddBrunSenderList is a paid mutator transaction binding the contract method 0xbc96092d.
//
// Solidity: function addBrunSenderList(address account) returns(bool)
func (_Mubs *MubsTransactor) AddBrunSenderList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "addBrunSenderList", account)
}

// AddBrunSenderList is a paid mutator transaction binding the contract method 0xbc96092d.
//
// Solidity: function addBrunSenderList(address account) returns(bool)
func (_Mubs *MubsSession) AddBrunSenderList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.AddBrunSenderList(&_Mubs.TransactOpts, account)
}

// AddBrunSenderList is a paid mutator transaction binding the contract method 0xbc96092d.
//
// Solidity: function addBrunSenderList(address account) returns(bool)
func (_Mubs *MubsTransactorSession) AddBrunSenderList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.AddBrunSenderList(&_Mubs.TransactOpts, account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address account) returns(bool)
func (_Mubs *MubsTransactor) AddWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "addWhiteList", account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address account) returns(bool)
func (_Mubs *MubsSession) AddWhiteList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.AddWhiteList(&_Mubs.TransactOpts, account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address account) returns(bool)
func (_Mubs *MubsTransactorSession) AddWhiteList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.AddWhiteList(&_Mubs.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mubs *MubsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mubs *MubsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.Approve(&_Mubs.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mubs *MubsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.Approve(&_Mubs.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_Mubs *MubsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_Mubs *MubsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.Burn(&_Mubs.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_Mubs *MubsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.Burn(&_Mubs.TransactOpts, amount)
}

// ChangeExchangeLock is a paid mutator transaction binding the contract method 0x9098fa79.
//
// Solidity: function changeExchangeLock() returns(bool)
func (_Mubs *MubsTransactor) ChangeExchangeLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "changeExchangeLock")
}

// ChangeExchangeLock is a paid mutator transaction binding the contract method 0x9098fa79.
//
// Solidity: function changeExchangeLock() returns(bool)
func (_Mubs *MubsSession) ChangeExchangeLock() (*types.Transaction, error) {
	return _Mubs.Contract.ChangeExchangeLock(&_Mubs.TransactOpts)
}

// ChangeExchangeLock is a paid mutator transaction binding the contract method 0x9098fa79.
//
// Solidity: function changeExchangeLock() returns(bool)
func (_Mubs *MubsTransactorSession) ChangeExchangeLock() (*types.Transaction, error) {
	return _Mubs.Contract.ChangeExchangeLock(&_Mubs.TransactOpts)
}

// ChangeOperate is a paid mutator transaction binding the contract method 0xbe80c917.
//
// Solidity: function changeOperate(address operateAccount) returns(bool)
func (_Mubs *MubsTransactor) ChangeOperate(opts *bind.TransactOpts, operateAccount common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "changeOperate", operateAccount)
}

// ChangeOperate is a paid mutator transaction binding the contract method 0xbe80c917.
//
// Solidity: function changeOperate(address operateAccount) returns(bool)
func (_Mubs *MubsSession) ChangeOperate(operateAccount common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.ChangeOperate(&_Mubs.TransactOpts, operateAccount)
}

// ChangeOperate is a paid mutator transaction binding the contract method 0xbe80c917.
//
// Solidity: function changeOperate(address operateAccount) returns(bool)
func (_Mubs *MubsTransactorSession) ChangeOperate(operateAccount common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.ChangeOperate(&_Mubs.TransactOpts, operateAccount)
}

// ExchangeA2U is a paid mutator transaction binding the contract method 0x32c839f0.
//
// Solidity: function exchangeA2U(uint256 amount) returns(bool)
func (_Mubs *MubsTransactor) ExchangeA2U(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "exchangeA2U", amount)
}

// ExchangeA2U is a paid mutator transaction binding the contract method 0x32c839f0.
//
// Solidity: function exchangeA2U(uint256 amount) returns(bool)
func (_Mubs *MubsSession) ExchangeA2U(amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.ExchangeA2U(&_Mubs.TransactOpts, amount)
}

// ExchangeA2U is a paid mutator transaction binding the contract method 0x32c839f0.
//
// Solidity: function exchangeA2U(uint256 amount) returns(bool)
func (_Mubs *MubsTransactorSession) ExchangeA2U(amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.ExchangeA2U(&_Mubs.TransactOpts, amount)
}

// RemoveBrunList is a paid mutator transaction binding the contract method 0xd4af2722.
//
// Solidity: function removeBrunList(address account) returns(bool)
func (_Mubs *MubsTransactor) RemoveBrunList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "removeBrunList", account)
}

// RemoveBrunList is a paid mutator transaction binding the contract method 0xd4af2722.
//
// Solidity: function removeBrunList(address account) returns(bool)
func (_Mubs *MubsSession) RemoveBrunList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.RemoveBrunList(&_Mubs.TransactOpts, account)
}

// RemoveBrunList is a paid mutator transaction binding the contract method 0xd4af2722.
//
// Solidity: function removeBrunList(address account) returns(bool)
func (_Mubs *MubsTransactorSession) RemoveBrunList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.RemoveBrunList(&_Mubs.TransactOpts, account)
}

// RemoveBrunSenderList is a paid mutator transaction binding the contract method 0x810d1475.
//
// Solidity: function removeBrunSenderList(address account) returns(bool)
func (_Mubs *MubsTransactor) RemoveBrunSenderList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "removeBrunSenderList", account)
}

// RemoveBrunSenderList is a paid mutator transaction binding the contract method 0x810d1475.
//
// Solidity: function removeBrunSenderList(address account) returns(bool)
func (_Mubs *MubsSession) RemoveBrunSenderList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.RemoveBrunSenderList(&_Mubs.TransactOpts, account)
}

// RemoveBrunSenderList is a paid mutator transaction binding the contract method 0x810d1475.
//
// Solidity: function removeBrunSenderList(address account) returns(bool)
func (_Mubs *MubsTransactorSession) RemoveBrunSenderList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.RemoveBrunSenderList(&_Mubs.TransactOpts, account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(address account) returns(bool)
func (_Mubs *MubsTransactor) RemoveWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "removeWhiteList", account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(address account) returns(bool)
func (_Mubs *MubsSession) RemoveWhiteList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.RemoveWhiteList(&_Mubs.TransactOpts, account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(address account) returns(bool)
func (_Mubs *MubsTransactorSession) RemoveWhiteList(account common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.RemoveWhiteList(&_Mubs.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mubs *MubsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mubs *MubsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mubs.Contract.RenounceOwnership(&_Mubs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mubs *MubsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mubs.Contract.RenounceOwnership(&_Mubs.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mubs *MubsTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mubs *MubsSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.Transfer(&_Mubs.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mubs *MubsTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.Transfer(&_Mubs.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mubs *MubsTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mubs *MubsSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.TransferFrom(&_Mubs.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mubs *MubsTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mubs.Contract.TransferFrom(&_Mubs.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mubs *MubsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mubs.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mubs *MubsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.TransferOwnership(&_Mubs.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mubs *MubsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mubs.Contract.TransferOwnership(&_Mubs.TransactOpts, newOwner)
}

// MubsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mubs contract.
type MubsApprovalIterator struct {
	Event *MubsApproval // Event containing the contract specifics and raw log

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
func (it *MubsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MubsApproval)
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
		it.Event = new(MubsApproval)
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
func (it *MubsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MubsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MubsApproval represents a Approval event raised by the Mubs contract.
type MubsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mubs *MubsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MubsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mubs.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MubsApprovalIterator{contract: _Mubs.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mubs *MubsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MubsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mubs.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MubsApproval)
				if err := _Mubs.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mubs *MubsFilterer) ParseApproval(log types.Log) (*MubsApproval, error) {
	event := new(MubsApproval)
	if err := _Mubs.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MubsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mubs contract.
type MubsTransferIterator struct {
	Event *MubsTransfer // Event containing the contract specifics and raw log

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
func (it *MubsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MubsTransfer)
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
		it.Event = new(MubsTransfer)
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
func (it *MubsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MubsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MubsTransfer represents a Transfer event raised by the Mubs contract.
type MubsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mubs *MubsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MubsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mubs.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MubsTransferIterator{contract: _Mubs.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mubs *MubsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MubsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mubs.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MubsTransfer)
				if err := _Mubs.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mubs *MubsFilterer) ParseTransfer(log types.Log) (*MubsTransfer, error) {
	event := new(MubsTransfer)
	if err := _Mubs.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
