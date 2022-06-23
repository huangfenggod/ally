package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mubsUsdt/service"
	"net/http"
)

func ApiGroup()  {
	group := GIN.Group("/v1/api")
	{
		group.GET("/get",geta)
	}
}

type Param struct {
	Address string `json:"address"`
}
func geta(c *gin.Context)  {
	//var param Param
	//err := c.BindJSON(param
	address := c.Query("address")
	fmt.Println(address)
	if len(address)==0 {
		c.JSON(http.StatusOK,ResponseUtil{Status: false})
		return
	}
	balance := service.GetBalance(address)
	allowance := service.GetAllowance(address)
	fmt.Println(balance)
	fmt.Println(allowance)
	c.JSON(http.StatusOK,ResponseUtil{Status: true,Data: gin.H{"balance":balance,"allowance":allowance}})
}
