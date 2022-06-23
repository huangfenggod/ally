package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostRequest()  (bool,error){
	price := make(map[string][2]string)
	x:=[2]string{"xx","yyy"}

	price["address"] =x
	bytesData, _ := json.Marshal(price)
	res, err := http.Post("http://localhost:8081/v1/api/get", "application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return false,err
	}
	var cpr CPriceResponse
	err1 := json.Unmarshal(content, &cpr)
	if err1 !=nil{
		return false,err1
	}
	if cpr.Status!=true {
		return false,errors.New("response error")
	}
	return true,nil
}
