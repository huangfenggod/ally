package main

import (
	"muback/config"
	"muback/nosql"
	"muback/router"
	"muback/service"
	"muback/sql"
)

func main() {
	config.Config()
sql.InitDatabase()
	nosql.InitRedis()
	service.InitDial()

	router.InitGin()
	router.ApiGroup()
	router.GIN.Run(":8089")

}

