package router

func ApiGroup()  {
	group := GIN.Group("/v1/api")
	{
	group.POST("/login",login)
	group.POST("/logout",logout)
	group.POST("/createUser",createUser)
	group.POST("/deleteUser",deleteUser)
	group.POST("/closeNotice",closeNotice)
	group.POST("/openNotice",openNotice)
	group.POST("/releaseNotice",releaseNotice)
	group.POST("/openAutoTrade",openAutoTrade)
	group.POST("/closeAutoTrade",closeAutoTrade)
	group.GET("/getStatus",getStatus)
	group.GET("/getUserList",getUserList)
	group.GET("/getAddressList",getAddressList)
	group.POST("/transferU",transferU)


	}
}
