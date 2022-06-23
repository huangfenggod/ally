package service

import (
	"muback/sql"
)

type CPriceResponse struct {
	Status	bool `json:"status"`
	Msg string 	`json:"msg"`
}



func OpenAutoTrade() (bool,error){
	//price := make(map[string]string)
	//price["price"] ="0"
	//bytesData, _ := json.Marshal(price)
	//res, err := http.Post("http://rebelbett.com:8082/v1/api/cprice", "application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	//
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//}
	//
	//defer res.Body.Close()
	//
	//content, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//	return false,err
	//}
	//var cpr CPriceResponse
	//err1 := json.Unmarshal(content, &cpr)
	//if err1 !=nil{
	//return false,err1
	//}
	//if cpr.Status!=true {
	//	return false,errors.New("response error")
	//}
	sql.ChangePrice(0)
	_, err2 := OpenEthAutoTrade()
	if err2!=nil {
		return false,err2
	}
	return true,nil

}

func CloseAutoTrade() (bool,error) {
	//price := make(map[string]string)
	//price["price"] ="10"
	//bytesData, _ := json.Marshal(price)
	//res, err := http.Post("http://rebelbett.com:8082/v1/api/cprice", "application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//}
	//defer res.Body.Close()
	//content, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//	return false,err
	//}
	//var cpr CPriceResponse
	//err1 := json.Unmarshal(content, &cpr)
	//if err1 !=nil{
	//	return false,err1
	//}
	//if cpr.Status!=true {
	//	return false,errors.New("response error")
	//}
	sql.ChangePrice(10)
	_, err2 := CloseEthAutoTrade()
	if err2!=nil {
		return false,err2
	}
	return true,nil

}


