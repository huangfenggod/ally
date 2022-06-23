package service

import "muback/sql"

func GetUserListByaddress(address string,pageSize int,pageNumber int) (int,[]sql.BindAddress) {
	totalPage, users, err := sql.GetUserListByAddress(address, pageSize, pageNumber)
	if err!=nil{
		return totalPage,users
	}
	for i:=0;i<len(users);i++ {
		balance,allowance, _ := GetUsdtBalanceAllowance(users[i].Address)
		users[i].BalanceU = float32(balance)
		users[i].Allowance=int64(allowance)

	}
	return totalPage,users
}
