package requestHandlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"server/cmd/server/constants"
	"server/cmd/server/env"
	"server/cmd/server/models"
	schedule2 "server/cmd/server/structure/schedule"
)

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")

func HandleGetCurrentSchedule(c *gin.Context){
	schedule, err:= getCurrentSchedule()

	if err != nil {
		c.String(http.StatusBadRequest, constants.ERROR_HANDLING)
	}

	c.JSON(http.StatusOK, schedule)
}



func getCurrentSchedule() ([]schedule2.ScheduleData, error) {
	filter := bson.D{{"status", bson.D {{"$in", bson.A{constants.LIVE_STATUS, constants.UPCOMING_STATUS}}}}}

	result,err := scheduleCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Println(err)
		return nil,err
	}

	var schedule []schedule2.ScheduleData

	err = result.All(context.TODO(), &schedule)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return schedule,nil
}