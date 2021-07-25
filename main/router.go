package main

import (
	"github.com/gin-gonic/gin"
	"main/ws"
	"net/http"
)

func runGinRouter(router *gin.Engine){
	router.GET("/live/:live_id", func(c *gin.Context) {
		liveId := c.Param("live_id")
		result, err := getTl(liveId)
		if err != nil {
			c.String(http.StatusBadRequest, "Request cannot be handle")
			return
		}
		c.JSON(http.StatusOK, result)
	})
	router.GET("/ws", func (c* gin.Context){
		ws.HandleWS(c.Writer, c.Request)
	})
}