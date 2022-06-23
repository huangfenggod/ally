package service

import (
	"errors"
	"muback/sql"
)

func CloseNotice(token string) (bool,error) {
	parseToken, err := ParseToken(token)
	if err !=nil{
		return false,err
	}
	if parseToken.Role!="admin"{
		return false,errors.New("dont have authority")
	}
sql.CloseNotice()
return true,nil
}

func OpenNotice(token string,openNotice bool)(bool,error)  {
	parseToken, err := ParseToken(token)
	if err !=nil{
		return false,err
	}
	if parseToken.Role!="admin"{
		return false,errors.New("dont have authority")
	}
	sql.OpenNotice(openNotice)
	return true,nil
}
func ReleaseNotice(token string,chineseText string,englishText string) (bool,error) {
	parseToken, err := ParseToken(token)
	if err !=nil{
		return false,err
	}
	if parseToken.Role!="admin"{
		return false,errors.New("dont have authority")
	}
	sql.ReleaseNotice(chineseText,englishText)
	return true,nil
}

