package schedule

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/structure/schedule"
	"time"
)

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")


func CreateSchedule(schedules []schedule.ScheduleData){
	updateOperation := make([]mongo.WriteModel , 0)

	for _,schedule := range schedules {
		scheduleOperation := mongo.NewUpdateOneModel()

		scheduleOperation.SetFilter(bson.M{"scheduleId": schedule.ScheduleId})
		scheduleOperation.SetUpdate(bson.D{
			{"$set",
				bson.D{
					{"lastUpdate", time.Now()},
					{"scheduleId", schedule.ScheduleId},
					{"title", schedule.Title},
					{"publishedAt", schedule.PublishedAt},
					{"availableAt", schedule.AvailableAt},
					{"duration", schedule.Duration},
					{"status", schedule.Status},
					{"channel", schedule.Channel},
				},
			},
		})
		scheduleOperation.SetUpsert(true)

		updateOperation = append(updateOperation, scheduleOperation)
	}

	result,err := scheduleCollection.BulkWrite(context.TODO(), updateOperation)

	if err!=nil {

		log.Println(err)
		return
	}
	log.Println(result)
}