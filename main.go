package main

import (
	"chain-contract/config"
	"chain-contract/router"
	"github.com/gin-gonic/gin"
)


func main() {

	config.Init()

	gin.SetMode(gin.ReleaseMode)
	g:=gin.Default()
	router.Startup(g)
	g.Run()
}
