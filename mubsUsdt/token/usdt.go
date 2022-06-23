// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

// UsdtABI is the input ABI used to generate the binding from.
const UsdtABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Usdt is an auto generated Go binding around an Ethereum contract.
type Usdt struct {
	UsdtCaller     // Read-only binding to the contract
	UsdtTransactor // Write-only binding to the contract
}

// UsdtCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdtSession struct {
	Contract     *Usdt             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdtCallerSession struct {
	Contract *UsdtCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsdtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdtTransactorSession struct {
	Contract     *UsdtTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdtRaw struct {
	Contract *Usdt // Generic contract binding to access the raw methods on
}

// UsdtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdtCallerRaw struct {
	Contract *UsdtCaller // Generic read-only contract binding to access the raw methods on
}

// UsdtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdtTransactorRaw struct {
	Contract *UsdtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdt creates a new instance of Usdt, bound to a specific deployed contract.
func NewUsdt(address common.Address, backend bind.ContractBackend) (*Usdt, error) {
	contract, err := bindUsdt(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdt{UsdtCaller: UsdtCaller{contract: contract}, UsdtTransactor: UsdtTransactor{contract: contract}}, nil
}

// NewUsdtCaller creates a new read-only instance of Usdt, bound to a specific deployed contract.
func NewUsdtCaller(address common.Address, caller bind.ContractCaller) (*UsdtCaller, error) {
	contract, err := bindUsdt(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtCaller{contract: contract}, nil
}

// NewUsdtTransactor creates a new write-only instance of Usdt, bound to a specific deployed contract.
func NewUsdtTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdtTransactor, error) {
	contract, err := bindUsdt(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UsdtTransactor{contract: contract}, nil
}

// bindUsdt binds a generic wrapper to an already deployed contract.
func bindUsdt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor,nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdt *UsdtRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Usdt.Contract.UsdtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdt *UsdtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.Contract.UsdtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdt *UsdtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdt.Contract.UsdtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdt *UsdtCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Usdt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdt *UsdtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdt *UsdtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdt.Contract.contract.Transact(opts, method, params...)
}

// _decimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() constant returns(uint8)
func (_Usdt *UsdtCaller) _decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "_decimals")
	return *ret0, err
}

// _decimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() constant returns(uint8)
func (_Usdt *UsdtSession) _decimals() (uint8, error) {
	return _Usdt.Contract._decimals(&_Usdt.CallOpts)
}

// _decimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() constant returns(uint8)
func (_Usdt *UsdtCallerSession) _decimals() (uint8, error) {
	return _Usdt.Contract._decimals(&_Usdt.CallOpts)
}

// _name is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() constant returns(string)
func (_Usdt *UsdtCaller) _name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "_name")
	return *ret0, err
}

// _name is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() constant returns(string)
func (_Usdt *UsdtSession) _name() (string, error) {
	return _Usdt.Contract._name(&_Usdt.CallOpts)
}

// _name is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() constant returns(string)
func (_Usdt *UsdtCallerSession) _name() (string, error) {
	return _Usdt.Contract._name(&_Usdt.CallOpts)
}

// _symbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() constant returns(string)
func (_Usdt *UsdtCaller) _symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "_symbol")
	return *ret0, err
}

// _symbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() constant returns(string)
func (_Usdt *UsdtSession) _symbol() (string, error) {
	return _Usdt.Contract._symbol(&_Usdt.CallOpts)
}

// _symbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() constant returns(string)
func (_Usdt *UsdtCallerSession) _symbol() (string, error) {
	return _Usdt.Contract._symbol(&_Usdt.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_Usdt *UsdtCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_Usdt *UsdtSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Usdt.Contract.Allowance(&_Usdt.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_Usdt *UsdtCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Usdt.Contract.Allowance(&_Usdt.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(account address) constant returns(uint256)
func (_Usdt *UsdtCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(account address) constant returns(uint256)
func (_Usdt *UsdtSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Usdt.Contract.BalanceOf(&_Usdt.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(account address) constant returns(uint256)
func (_Usdt *UsdtCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Usdt.Contract.BalanceOf(&_Usdt.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Usdt *UsdtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Usdt *UsdtSession) Decimals() (uint8, error) {
	return _Usdt.Contract.Decimals(&_Usdt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Usdt *UsdtCallerSession) Decimals() (uint8, error) {
	return _Usdt.Contract.Decimals(&_Usdt.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Usdt *UsdtCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Usdt *UsdtSession) GetOwner() (common.Address, error) {
	return _Usdt.Contract.GetOwner(&_Usdt.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Usdt *UsdtCallerSession) GetOwner() (common.Address, error) {
	return _Usdt.Contract.GetOwner(&_Usdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Usdt *UsdtCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Usdt *UsdtSession) Name() (string, error) {
	return _Usdt.Contract.Name(&_Usdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Usdt *UsdtCallerSession) Name() (string, error) {
	return _Usdt.Contract.Name(&_Usdt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Usdt *UsdtCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Usdt *UsdtSession) Owner() (common.Address, error) {
	return _Usdt.Contract.Owner(&_Usdt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Usdt *UsdtCallerSession) Owner() (common.Address, error) {
	return _Usdt.Contract.Owner(&_Usdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Usdt *UsdtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Usdt *UsdtSession) Symbol() (string, error) {
	return _Usdt.Contract.Symbol(&_Usdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Usdt *UsdtCallerSession) Symbol() (string, error) {
	return _Usdt.Contract.Symbol(&_Usdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Usdt *UsdtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Usdt.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Usdt *UsdtSession) TotalSupply() (*big.Int, error) {
	return _Usdt.Contract.TotalSupply(&_Usdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Usdt *UsdtCallerSession) TotalSupply() (*big.Int, error) {
	return _Usdt.Contract.TotalSupply(&_Usdt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_Usdt *UsdtTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_Usdt *UsdtSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Approve(&_Usdt.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Approve(&_Usdt.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns(bool)
func (_Usdt *UsdtTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns(bool)
func (_Usdt *UsdtSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Burn(&_Usdt.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Burn(&_Usdt.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_Usdt *UsdtTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_Usdt *UsdtSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.DecreaseAllowance(&_Usdt.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.DecreaseAllowance(&_Usdt.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_Usdt *UsdtTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_Usdt *UsdtSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.IncreaseAllowance(&_Usdt.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.IncreaseAllowance(&_Usdt.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns(bool)
func (_Usdt *UsdtTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns(bool)
func (_Usdt *UsdtSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Mint(&_Usdt.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Mint(&_Usdt.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdt *UsdtTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdt *UsdtSession) RenounceOwnership() (*types.Transaction, error) {
	return _Usdt.Contract.RenounceOwnership(&_Usdt.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdt *UsdtTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Usdt.Contract.RenounceOwnership(&_Usdt.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(recipient address, amount uint256) returns(bool)
func (_Usdt *UsdtTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(recipient address, amount uint256) returns(bool)
func (_Usdt *UsdtSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Transfer(&_Usdt.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(recipient address, amount uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Transfer(&_Usdt.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(sender address, recipient address, amount uint256) returns(bool)
func (_Usdt *UsdtTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(sender address, recipient address, amount uint256) returns(bool)
func (_Usdt *UsdtSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.TransferFrom(&_Usdt.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(sender address, recipient address, amount uint256) returns(bool)
func (_Usdt *UsdtTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.TransferFrom(&_Usdt.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Usdt *UsdtTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Usdt *UsdtSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdt.Contract.TransferOwnership(&_Usdt.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Usdt *UsdtTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdt.Contract.TransferOwnership(&_Usdt.TransactOpts, newOwner)
}
