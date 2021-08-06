package routers

import (
	"github.com/gin-gonic/gin"
	"server/cmd/server/requestHandlers"
)

func RouteCurrentSchedule(router *gin.Engine){
	router.GET("/schedules/current", requestHandlers.HandleGetCurrentSchedule)
}