package requestHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/cmd/server/models/translation"
)

func HandleGetTranslation(c *gin.Context) {
	liveId := c.Param("live_id")
	result, err := translation.GetTranslation(liveId)
	if err != nil {
		c.String(http.StatusBadRequest, "Request cannot be handle")
		return
	}
	c.JSON(http.StatusOK, result)
}
