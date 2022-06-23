package service

import (
	"muback/tool"
)

func GetUsdtBalanceAllowance(address string)(int,int,error)  {
	request, err := GetRequest(address)
	if err!=nil {
		return 0,0,err
	}
	return tool.GetInterfaceToInt(request.Data["balance"]),tool.GetInterfaceToInt(request.Data["allowance"]),nil

}
