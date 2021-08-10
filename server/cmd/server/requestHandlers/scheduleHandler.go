package requestHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/cmd/server/constants"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/models/scheduleModel"
)

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")

func HandleGetCurrentSchedule(c *gin.Context) {
	scheduleData, err := scheduleModel.GetCurrentSchedule()

	if err != nil {
		c.String(http.StatusBadRequest, constants.ERROR_HANDLING)
	}

	c.JSON(http.StatusOK, scheduleData)
}
