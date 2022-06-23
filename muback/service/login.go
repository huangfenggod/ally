package service

import (
	"errors"
	"muback/nosql"
	"muback/sql"
)
//登录返回token
func Login(username string,passwd string) (string,string,error){
	ad := sql.LoginUserPasswd(username)
	if ad.Username!=username {
		return "","",errors.New("user not exists")
	}
	if ad.Password !=passwd {
		return "","",errors.New("password wrong")
	}
	token, err := GenerateToken(username, ad.Role)
	if err !=nil {
		return "","",err
	}else {
		nosql.SetTokenInRedis(token)
		return token,ad.Role,nil
	}
}

func Logout(token string) bool {
	isSuccess := nosql.RemoveToken(token)
	return isSuccess
}
