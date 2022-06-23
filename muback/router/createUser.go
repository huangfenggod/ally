package router

import (
	"github.com/gin-gonic/gin"
	"muback/service"
	"net/http"
)
type cUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Bind string `json:"bind"`
}
func createUser(c *gin.Context)  {
	var cuser cUser
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	c.BindJSON(&cuser)
	if len(token)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "not login"})
		return
	}
	if len(cuser.Username)==0||len(cuser.Password)==0||len(cuser.Bind)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "para wrong"})
		return
	}
	_, err := service.CreateUser(token, cuser.Username, cuser.Password, cuser.Bind)
	if err!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
		return
	}
	//refresh token
	b, s := service.ReFreshToken(token)
	if b {
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "create user success"})
}

func deleteUser(c *gin.Context)  {
	var cuser cUser
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	c.BindJSON(&cuser)
	if len(token)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "not login"})
		return
	}
	if len(cuser.Username)==0{
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "para wrong"})
		return
	}
	_, err := service.DeleteUser(cuser.Username, token)
	if err!=nil {
	c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
		return
	}
	//refresh token
	b, s := service.ReFreshToken(token)
	if b {
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "delete success"})
}
