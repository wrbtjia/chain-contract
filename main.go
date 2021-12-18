package main

import (
	"chain-contract/router"
	"github.com/gin-gonic/gin"
)


func main() {

	gin.SetMode(gin.ReleaseMode)
	g:=gin.Default()
	router.Startup(g)
	g.Run()
}
