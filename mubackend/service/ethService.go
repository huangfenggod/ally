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
	"mubackend/config"
	"mubackend/sql"
	"mubackend/token"
	"strconv"
)

var Conn *ethclient.Client
var mubs *token.Mubs
var mubfunc *token.Mufunc

//var PriceControll float32
//这是让bigint减少10位转换成int64
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

//这是把int64末尾添加10位0转成bigint
func intTobigInt(in int64) *big.Int {

	inString := strconv.FormatInt(in, 10)
	newString := inString+"0000000000"
	bigint, _ := new(big.Int).SetString(newString, 10)
	return bigint
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

func InitDial()  {
	dial, err := ethclient.Dial(config.Cfg.Ethereum.Network)
	if err!=nil {
		log.Printf("connect eth fail because:%s", err)
	}
	Conn = dial
	mub, err := token.NewMubs(common.HexToAddress(config.Cfg.Ethereum.MubsAddress), dial)
	if err !=nil {
		log.Printf("mubs wrong because: %s",err)
	}
	mubs = mub
	mubf, err := token.NewMufunc(common.HexToAddress(config.Cfg.Ethereum.ContractAddress), dial)
	if err !=nil {
		log.Printf("mubs wrong because: %s",err)
	}
	mubfunc = mubf
}



//获取池子占比
func GetRate() (float32,error)  {
	pool, total, err := mubfunc.GetRate(&bind.CallOpts{
		Pending:     false,
		From:        common.HexToAddress(config.Cfg.Ethereum.ContractAddress),
		BlockNumber: nil,
		Context:     nil,
	})

	if err !=nil{
		log.Printf("get rate fail : %s",err)
		return float32(0),err
	}
	poolInt, err := bigIntToInt(pool)
	if err !=nil {
		return float32(0),err
	}
	totalInt, err := bigIntToInt(total)
	if err !=nil {
		return float32(0),err
	}
	rate := float32(poolInt)/float32(totalInt)
	return rate,nil
}
//获取目前基金总共进量
func GetFundAmount() (int64,error) {
	amount := int64(0)
	amount1, err := mubs.GetFundAmount(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	})
	if err!=nil {
		return amount,err
	}
	toInt, err := bigIntToInt(amount1)
	if err !=nil {
		return amount,err
	}
	return toInt,err
}



func GetPrice() float32 {
	priceControll := sql.GetPriceOfContent()
	if priceControll>0{
	return float32(priceControll)
	}
	priceA, priceB, err := mubs.GetPrice(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	})
	if err !=nil{
		return float32(0)
	}
	price1, err := bigIntToInt(priceA)
	if err !=nil {
		return float32(0)
	}
	price2, err := bigIntToInt(priceB)
	if err !=nil{
		return float32(0)
	}
	if price1 ==1 {
		return float32(0)
	}
	return  float32(price2)/float32(price1)
}


//链转账功能
func Transfer(address string ,amount1 int64 ,amount2 int64) (string ,error){
	var txhash string
	tobigInt1 := intTobigInt(amount1)
	tobigInt2 := intTobigInt(amount2)
	auth, err2 := getTransactOpts()
	if err2 !=nil{
		return txhash,err2
	}
	transaction, err := mubfunc.TransferToCustomer(&bind.TransactOpts{
		From:      auth.From,
		//Nonce:     auth.Nonce,
		Signer:    auth.Signer,
		//Value:     big.NewInt(0),
		//GasPrice:  auth.GasPrice,
		//GasFeeCap: auth.GasFeeCap,
		//GasTipCap: auth.GasTipCap,
		//GasLimit:  auth.GasLimit,
		//Context:   auth.Context,
		//NoSend:    true,
	}, common.HexToAddress(address), tobigInt1, tobigInt2)
	if err!=nil {
		return txhash ,err
	}
	mined, err2 := bind.WaitMined(context.Background(), Conn, transaction)
	if err2 !=nil {
		log.Printf("waitMined transfer fail  %s",err2)
		return txhash,err2
	}
	txhash =mined.TxHash.String()
	log.Printf("transfer success address: %s ,TxHash: %s",address,mined.TxHash)
	return txhash ,nil
}
