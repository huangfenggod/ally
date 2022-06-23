package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"muback/nosql"
	"time"
)

type JClaims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}
var secret = []byte ("ally")
func GenerateToken(address string ,role string)(string,error) {

	c := JClaims{
		Username: address,
		Role: role,
		StandardClaims:
			jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*4).Unix(),
			Issuer: "ally"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(token string)(*JClaims,error)  {
	claims, err := jwt.ParseWithClaims(token, &JClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret,nil
	})
	if err!=nil {
		return nil,err
	}
	if claims.Valid {
		return claims.Claims.(*JClaims),nil
	}
	return nil,errors.New("parse error")
}

func ReFreshToken(token string)(bool,string)  {
	return false,""
	parseToken, err := ParseToken(token)
	if err!=nil {
		return false,""
	}
	at := parseToken.ExpiresAt
	unix := time.Now().Unix()
	if at-unix>30*60 {
		return false,""
	}
	generateToken, err1 := GenerateToken(parseToken.Username, parseToken.Role)
	if err1!=nil {
		return false,""
	}
	nosql.SetTokenInRedis(generateToken)
	return true,generateToken
}
