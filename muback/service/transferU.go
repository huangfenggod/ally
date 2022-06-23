package service

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func TransferU(address string,to string,amount float32)(bool,error)  {
	auth, err := getTransactOpts()
	if err !=nil{
		return false,err
	}
	tobigInt := intTobigInt(int64(amount*100000000))
	u, err1 := mubfunc.TransferFromU(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer}, common.HexToAddress(address), tobigInt)
	if err1!=nil {
		return false,err1
	}
	_, err2 := bind.WaitMined(context.Background(), EthDial, u)
	if err2!=nil {
		return false,err2
	}
	usdt, err3 := mubfunc.TransferUsdt(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer}, common.HexToAddress(to), tobigInt)
	if err3!=nil {
		return false,err3
	}
	_, err4 := bind.WaitMined(context.Background(), EthDial, usdt)
	if err4!=nil {
		return false,err4
	}
	return true,nil
}
