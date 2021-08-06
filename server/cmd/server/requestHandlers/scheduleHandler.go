package requestHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/cmd/server/constants"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/models/schedule"
)

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")

func HandleGetCurrentSchedule(c *gin.Context){
	scheduleData, err:= schedule.GetCurrentSchedule()

	if err != nil {
		c.String(http.StatusBadRequest, constants.ERROR_HANDLING)
	}

	c.JSON(http.StatusOK, scheduleData)
}



