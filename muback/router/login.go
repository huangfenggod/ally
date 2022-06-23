package router

import (
	"github.com/gin-gonic/gin"
	"muback/service"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func login(c *gin.Context)  {
	var u user
	err := c.BindJSON(&u)
	if err!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "params wrong" })
		return
	}
	if len(u.Username)==0||len(u.Password)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "login fail"})
		return
	}
	token,role,err1 := service.Login(u.Username, u.Password)
	if err1 !=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "username or password wrong"})
		return
	}
	c.Header("token",token)
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success",Data: gin.H{"role":role}})
}

func logout(c *gin.Context)  {
	token := c.GetHeader("token")
	//status := service.CheckTokenStatus(token)
	//if !status {
	//	c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
	//	return
	//}
	b := service.Logout(token)
	if b{
		c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "logout success"})
		return
	} else{
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "logout fail"})
	}
}
