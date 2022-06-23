package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
Status bool `json:"status"`
Msg string `json:"msg"`
Data map[string]interface{} `json:"data"`
}



func GetRequest(address string) (Response,error) {
	var bsc Response
	params := url.Values{}
	Url, _:= url.Parse("http://localhost:8090/v1/api/get")
	params.Set("address",address)

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	response, err := http.Get(urlPath)
	if err !=nil {

		return bsc, err
	}
	defer response.Body.Close()
	all,_ := ioutil.ReadAll(response.Body)
	toStruct := StringToStruct(all)
	return toStruct,err
}
func StringToStruct(byteInfo []byte ) Response  {
	var bsc Response
	err := json.Unmarshal(byteInfo, &bsc)
	if err !=nil {
		fmt.Println(err)
	}
	return bsc
}
