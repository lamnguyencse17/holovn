package requestHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/cmd/server/models/translationStore"
)

func HandleGetTranslation(c *gin.Context) {
	liveId := c.Param("live_id")
	result, err := translationStore.GetTranslation(liveId)
	if err != nil {
		c.String(http.StatusBadRequest, "Request cannot be handle")
		return
	}
	c.JSON(http.StatusOK, result)
}
