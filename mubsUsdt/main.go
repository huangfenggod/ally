package main

import (
	"mubsUsdt/router"
	"mubsUsdt/service"
)

func main() {
	service.InitDial()
	router.InitGin()
	router.ApiGroup()
	router.GIN.Run(":8090")
}
