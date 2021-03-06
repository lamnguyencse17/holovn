package requestHandlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/cmd/server/constants"
	"server/cmd/server/models/translationStore"
	"server/cmd/server/redis"
)

func HandleGetTranslation(c *gin.Context) {
	liveId := c.Param("live_id")
	unparsedTranslation := redis.GetValue(liveId)
	if unparsedTranslation != "" {
		var translations translationStore.ITranslationStore
		err := json.Unmarshal([]byte(unparsedTranslation), &translations)
		if err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, constants.ERROR_HANDLING)
			return
		}
		c.JSON(http.StatusOK, translations)
		return
	}
	result, err := translationStore.GetTranslation(liveId, 0)
	if err != nil {
		c.JSON(http.StatusOK, result)
		return
	}
	stringifiedResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, constants.ERROR_HANDLING)
		return
	}
	redis.SetKeyValueWithExpiration(liveId, string(stringifiedResult), 15)
	if err != nil {
		c.String(http.StatusBadRequest, constants.ERROR_HANDLING)
		return
	}
	c.JSON(http.StatusOK, result)
}
