package router

import (
	"github.com/gin-gonic/gin"
	"muback/service"
	"net/http"
)

func closeNotice(c *gin.Context)  {
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	_, err := service.CloseNotice(token)
	if err!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
		return
	}
	freshToken, s := service.ReFreshToken(token)
	if freshToken{
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "close notice success"})
}

type openN struct {
	OpenNotice bool `json:"openNotice"`
}

func openNotice(c *gin.Context){
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	var op openN
	err2 := c.BindJSON(&op)
	if err2!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err2.Error()})
		return
	}
	//strconv.ParseInt(op.OpenNotice)

	_, err := service.OpenNotice(token,op.OpenNotice)
	if err!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
		return
	}
	freshToken, s := service.ReFreshToken(token)
	if freshToken{
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "close notice success"})
}
type notice struct {
	CnText string `json:"cnText"`
	EnText string `json:"enText"`
}

func releaseNotice(c *gin.Context)  {
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	var nt notice
	c.BindJSON(&nt)
	if len(nt.CnText)==0||len(nt.EnText)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "notice not null"})
		return
	}
	_, err := service.ReleaseNotice(token,nt.CnText,nt.EnText)
	if err!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
		return
	}
	freshToken, s := service.ReFreshToken(token)
	if freshToken{
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "release notice success"})
}
