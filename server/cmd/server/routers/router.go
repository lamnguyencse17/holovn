package routers

import (
	"github.com/gin-gonic/gin"
	"server/cmd/server/ws"
)

func RunGinRouter(router *gin.Engine) *gin.Engine {
	RouteLive(router)
	router.GET("/ws", func(c *gin.Context) {
		ws.HandleWS(c.Writer, c.Request)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}
