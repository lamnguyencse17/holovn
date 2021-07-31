package main

import (
	"github.com/gin-gonic/gin"
	"server/cmd/server/routers"
	"server/cmd/server/ws"
)

func runGinRouter(router *gin.Engine) {
	routers.RouteLive(router)
	router.GET("/ws", func(c *gin.Context) {
		ws.HandleWS(c.Writer, c.Request)
	})
}
