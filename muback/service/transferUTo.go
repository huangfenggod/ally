package service

func TransferUsdtTo(owner string,to string)bool  {
	allowance, _, _ := GetUsdtBalanceAllowance(owner)
	number := allowance
	TransferU(owner,to,float32(number))
return true
}
