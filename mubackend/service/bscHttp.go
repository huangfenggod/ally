package service

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	"mubackend/config"
	"mubackend/nosql"
	"net/http"
	"net/url"
	"strconv"
)

type BscResponse struct {
	Status	string `json:"status"`
	Message string 	`json:"message"`
	Result map[string]interface{}	`json:"result"`
}


//方法1
func BscHttp(transaction string)(error , BscResponse){
	var bsc BscResponse
	params := url.Values{}
	Url, _:= url.Parse("https://api.bscscan.com/api/")
	params.Set("module","transaction")
	params.Set("action","gettxreceiptstatus")
	params.Set("txhash",transaction)
	params.Set("apikey",config.Cfg.ApiKey)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	response, err := http.Get(urlPath)
	if err !=nil {
		log.Fatalf("get bscan transaction:%s fail",transaction)
		return err, bsc
	}
	defer response.Body.Close()
	all,_ := ioutil.ReadAll(response.Body)
	toStruct := StringToStruct(all)
	return err, toStruct
}

func StringToStruct(byteInfo []byte ) BscResponse  {
	var bsc BscResponse
	 err := json.Unmarshal(byteInfo, &bsc)
	if err !=nil {
		log.Fatal("err get from bsc ")
	}
	return bsc
}
func IdentifyTransactionHash(transactionHash string) bool{
	err, response := BscHttp(transactionHash)
	if err !=nil {
		log.Printf(" tradeHash:%s acquire treasure fail because check BNB fail: %s",transactionHash,err)
		return false
	}
	if response.Result["status"] != "1" {
		return false
	}
	has := nosql.IsHas(transactionHash)
	if has {
		return false
	}
	nosql.InsertHashToRedis(transactionHash)
	return true
}
//方法2
//使用bind去查询是否存在
func checkHash(tradHash string) (bool,error)  {
	_, _, err := Conn.TransactionByHash(context.Background(), common.HexToHash(tradHash))
	if err !=nil {
		return false,err
	}
	return true,nil
}

func IdentifyTransactionHash2(transactionHash string)bool  {
	has := nosql.IsHas(transactionHash)
	if has {
		return false
	}
	hash, _ := checkHash(transactionHash)
	log.Printf("txhash: %s ,bool: %s",transactionHash,strconv.FormatBool(hash))
	hash =true
	if  !hash{
		return false
	}else {
		nosql.InsertHashToRedis(transactionHash)
	return true
	}
}
