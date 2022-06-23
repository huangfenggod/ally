package main

import (
	"mubackend/config"
	"mubackend/log"
	"mubackend/nosql"
	"mubackend/router"
	"mubackend/service"
	"mubackend/sql"
)

func main()  {
	//service.PriceControll =10
	config.Config()
	log.InitLog()
	nosql.InitRedis()
	sql.InitDatabase()
	service.InitDial()
	router.InitGin()

	service.InitCron()
	router.ApiGroup()
	router.GIN.Run(":8082")

	defer sql.DB.Close()
	defer nosql.RedisConn.Close()
	defer service.Conn.Close()
}
