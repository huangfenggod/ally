package service

import (
	"errors"
	"muback/nosql"
	"muback/sql"
)

func CreateUser(token string,username string,password string,bindAddress string) (bool,error) {
	existsToken := nosql.ExistsToken(token)
	if !existsToken {
		return false,errors.New("login invalid")
	}
	parseToken, err := ParseToken(token)
	if err!=nil {
		return false,err
	}
	if parseToken.Role!="admin" {
		return false,errors.New("dont have this authority")
	}

	_, err1 := sql.CreateUser(username, password, bindAddress)
	if err1!=nil {
		return false,err1
	}
	return true,nil
}

func DeleteUser(username string,token string) (bool,error) {
	existsToken := nosql.ExistsToken(token)
	if !existsToken {
		return false,errors.New("login invalid")
	}
	parseToken, err := ParseToken(token)
	if err!=nil {
		return false,err
	}
	if parseToken.Role!="admin"{
		return false,errors.New("dont have this authority")
	}
	sql.DeleteUser(username)
	return true,nil
}
