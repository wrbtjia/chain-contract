package router

import (
	"chain-contract/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Startup(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// Middlewares.
	g.Use(gin.Recovery())

	g.Use(mw...)


	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})




	g.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	g.GET("/q",handler.Process)

	return g
}