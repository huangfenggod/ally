package sql

import "time"

func CashRecorder(txhash string,address string,amount int64,price float32)  {
	DB.Exec("insert into withdraw(txhash,address,amount,price,create_time) values(?,?,?,?,?)",txhash,address,amount,price,time.Now())
}


