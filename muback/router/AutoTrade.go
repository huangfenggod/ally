package router

import (
	"github.com/gin-gonic/gin"
	"muback/config"
	"muback/service"
	"net/http"
)

func openAutoTrade(c *gin.Context)  {
	token := c.GetHeader("token")
	_, err10 := service.ParseToken(token)
	if err10!=nil {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	if len(token)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "not login"})
		return
	}
	var op openN
	err2 := c.BindJSON(&op)
	if err2!=nil {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err2.Error()})
		return
	}
	if op.OpenNotice {
		_, err := service.OpenAutoTrade()
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
			return
		}
	}else {
		_, err := service.CloseEthAutoTrade()
		if err!=nil {
			c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: err.Error()})
			return
		}
	}

	//refresh token
	b, s := service.ReFreshToken(token)
	if b {
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "delete success"})
}

func closeAutoTrade(c *gin.Context)  {
	token := c.GetHeader("token")
	_, err10 := service.ParseToken(token)
	if err10!=nil {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	if len(token)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "not login"})
		return
	}
	_, err := service.CloseAutoTrade()
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

func getStatus(c *gin.Context)  {
	token := c.GetHeader("token")
	_, err10 := service.ParseToken(token)
	if err10!=nil {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	if len(token)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "not login"})
		return
	}
	b1, b2 := service.GetStatus()
	//refresh token
	b, s := service.ReFreshToken(token)
	if b {
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "success",Data: gin.H{"noticeStatus":b1,"autoTradeStatus":b2}})
}

type address struct {
	Address string `json:"address"`
}

func transferU(c *gin.Context)  {
	token := c.GetHeader("token")
	status := service.CheckTokenStatus(token)
	if !status {
		c.JSON(http.StatusUnauthorized,gin.H{"msg":"unLogin"})
		return
	}
	var ad address
	c.BindJSON(&ad)
	bo := service.TransferUsdtTo(ad.Address, config.Cfg.Ethereum.UsdtReceiver)
	if !bo {
		c.JSON(http.StatusOK,ResponseUtil{Status: false,Msg: "transfer fail"})
		return
	}
	freshToken, s := service.ReFreshToken(token)
	if freshToken{
		c.Header("token",s)
	}
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Msg: "close notice success"})


}
