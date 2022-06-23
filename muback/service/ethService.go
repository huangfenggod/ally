package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"muback/config"
	"muback/token"
	"strconv"
)

var EthDial  *ethclient.Client
var mubs *token.Mubs
var mubfunc *token.Mufunc
//var usdts *token.Usdts
func InitDial()  {
	dial, err := ethclient.Dial(config.Cfg.Ethereum.Network)
	if err !=nil {
		log.Println(err)
	}
	EthDial = dial
	mub, err1 := token.NewMubs(common.HexToAddress(config.Cfg.Ethereum.MubsAddress), dial)
	if err !=nil {
		log.Printf("mubs wrong because: %s",err1)
	}
	mubs = mub
	mubf, err2 := token.NewMufunc(common.HexToAddress(config.Cfg.Ethereum.ContractAddress), dial)
	if err !=nil {
		log.Printf("mubs wrong because: %s",err2)
	}
	mubfunc = mubf
	//us, err3 := token.NewUsdts(common.HexToAddress(config.Cfg.Ethereum.UsdtsAddress), dial)
	//if err !=nil {
	//	log.Printf("mubs wrong because: %s",err3)
	//}
	//usdts = us
}
func getTransactOpts() (*bind.TransactOpts ,error) {
	privateKey, err := crypto.HexToECDSA(config.Cfg.Ethereum.KeyStore)
	if err !=nil {
		log.Fatalln(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return nil,errors.New("has got treasure")
	}
	crypto.PubkeyToAddress(*publicKeyECDSA)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(config.Cfg.Ethereum.ChainId)))
	if err !=nil {
		log.Println(err)
		return nil,errors.New("privatekey wrong")
	}
	return auth ,nil
}

func GetStatusAutoTrade() (bool,error) {

	autotrade, err := mubfunc.Autotrade(&bind.CallOpts{
		Pending:     false,
		From:        common.HexToAddress(config.Cfg.Ethereum.ContractAddress),
		BlockNumber: nil,
		Context:     nil,
	})

	return autotrade,err
}

func OpenEthAutoTrade() (bool,error) {
	autotrade, err := mubfunc.Autotrade(&bind.CallOpts{
		Pending:     false,
		From:        common.HexToAddress(config.Cfg.Ethereum.ContractAddress),
		BlockNumber: nil,
		Context:     nil,
	})
	if err!=nil {
		return false,err
	}
	if autotrade {
		return true,nil
	}
	auth, err2 := getTransactOpts()
	if err2 !=nil{
		return false,err2
	}
	transcation, err1 := mubfunc.ChangeAutoTrade(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer})
	if err!=nil {
		return false, err1
	}
	_,err3 := bind.WaitMined(context.Background(), EthDial, transcation)
	if err3 !=nil {
		log.Printf("waitMined transfer fail  %s",err2)
		return false,err2
	}
	return true,nil
}
func CloseEthAutoTrade() (bool,error) {
	autotrade, err := mubfunc.Autotrade(&bind.CallOpts{
		Pending:     false,
		From:        common.HexToAddress(config.Cfg.Ethereum.ContractAddress),
		BlockNumber: nil,
		Context:     nil,
	})
	if err!=nil {
		return false,err
	}
	if !autotrade {
		return true,nil
	}
	auth, err2 := getTransactOpts()
	if err2 !=nil{
		return false,err2
	}
	transcation, err1 := mubfunc.ChangeAutoTrade(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer})
	if err!=nil {
		return false, err1
	}
	_,err3 := bind.WaitMined(context.Background(), EthDial, transcation)
	if err3 !=nil {
		log.Printf("waitMined transfer fail  %s",err2)
		return false,err2
	}
	return true,nil
}

func GetUsdtNumber(address string)float32 {
	return 0
//	uNumber, _ := usdts.BalanceOfU(&bind.CallOpts{
//		Pending:     false,
//		From:        common.Address{},
//		BlockNumber: nil,
//		Context:     nil}, common.HexToAddress(address))
//	toInt, _ := bigIntToInt(uNumber)
//	return float32(toInt)/100000000
//}
//func getUsatAllowance(address string)(int64,error){
//	return 0,nil
//	allowance, _ := usdts.Allowance(&bind.CallOpts{
//		Pending:     false,
//		From:        common.Address{},
//		BlockNumber: nil,
//		Context:     nil}, common.HexToAddress(address), common.HexToAddress(config.Cfg.Ethereum.ContractAddress))
//	toInt, _ := bigIntToInt18(allowance)
//	return toInt,nil
}
func getUsatAllowance(address string) (int64,error) {
	return 0,nil
}
func bigIntToInt(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	if len(s)<=10 {
		return int64(0),nil
	}
	s1 := s[0 : len(s)-10]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err !=nil {
		return in,err
	}
	return atoi,nil
}

func bigIntToInt18(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	if len(s)<=10 {
		return int64(0),nil
	}
	s1 := s[0 : len(s)-18]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err !=nil {
		return in,err
	}
	return atoi,nil
}
//这是把int64末尾添加10位0转成bigint
func intTobigInt(in int64) *big.Int {

	inString := strconv.FormatInt(in, 10)
	newString := inString+"0000000000"
	bigint, _ := new(big.Int).SetString(newString, 10)
	return bigint
}
