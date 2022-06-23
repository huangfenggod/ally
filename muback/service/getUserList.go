package service

import "muback/sql"

func GetUserList(pageSize int,pageNumber int)(int,[]sql.BindAddress)  {
	number, addresses := sql.GetAllUserList(pageSize, pageNumber)
	for i:=0;i<len(addresses);i++ {
		balance,allowance, _ := GetUsdtBalanceAllowance(addresses[i].Address)

		addresses[i].BalanceU = float32(balance)
		addresses[i].Allowance=int64(allowance)
	}
return number,addresses
}

func GetAminList(pageSize int,pageNumber int)(int,[]sql.Admin)  {
	number := sql.GetAdminNumber()
	list := sql.GetUserList(pageSize,pageNumber)
	return number,list
}
