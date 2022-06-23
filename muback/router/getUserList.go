package router

import (
	"github.com/gin-gonic/gin"
	"muback/service"
	"net/http"
	"strconv"
)

func getUserList(c *gin.Context)  {
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	pageSize:= c.Query("pageSize")
	pageNumber:= c.Query("pageNumber")
	size, err1 := strconv.ParseInt(pageSize, 10, 64)
	number, err2 := strconv.ParseInt(pageNumber, 10, 64)
	if err1!=nil||err2!=nil{
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err1.Error()})
		return
	}
	if size<=0||number<=0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "size,number invalid"})
		return
	}
	list, admins := service.GetAminList(int(size), int(number))
	freshToken, s := service.ReFreshToken(token)
	if freshToken{
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success",Data: gin.H{"pageTotal":list,"userList":admins}})
}
