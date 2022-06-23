package service

import "muback/nosql"

func CheckTokenStatus(token string)bool  {
	existsToken := nosql.ExistsToken(token)
	if !existsToken {
		return false
	}
	_, err := ParseToken(token)
	if err!=nil {
		return false
	}
	return true
}
