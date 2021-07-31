package routers

import (
	"github.com/gin-gonic/gin"
	"server/cmd/server/requestHandlers"
)

func RouteLive(router *gin.Engine) {
	router.GET("/live/:live_id/translations", requestHandlers.HandleGetTranslation)
}
