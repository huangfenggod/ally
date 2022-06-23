package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"mubsUsdt/token"
	"strconv"
)
var usdt *token.Usdt
func InitDial()  {
	dial, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err !=nil {
		log.Println(err)
	}

	usdt1,err := token.NewUsdt(common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), dial)
	if err!=nil {
		fmt.Println(err)
	}
	usdt=usdt1
}

func GetBalance(address string)float32  {
	balance, err1 := usdt.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil}, common.HexToAddress(address))
	if err1!=nil {
		fmt.Println(err1)
	}
	int18, err3 := bigIntToInt18(balance)
	if err3 !=nil {
		fmt.Println(int18)
	}
	return float32(int18)

}
func bigIntToInt18(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	if len(s)<=18 {
		return int64(0),nil
	}
	s1 := s[0 : len(s)-18]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err!=nil {
		if len(s)>18 {
			return 10000000000,nil
		}else {
			return in,err
		}
	}
	return atoi,nil
}

func GetAllowance(address string)float32  {
	allowance, err1 := usdt.Allowance(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil}, common.HexToAddress(address),common.HexToAddress("0x42a49394574bd6f66f148cf8B856B6aDBB43FdBe"))
	if err1 !=nil{
		fmt.Println(err1)
	}
	int18, err2:= bigIntToInt18(allowance)
	if err2 !=nil {
		fmt.Println(err2)
	}
	return float32(int18)
}
