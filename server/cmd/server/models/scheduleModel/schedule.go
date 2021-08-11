package scheduleModel

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"server/cmd/server/constants"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/structure/scheduleStruct"
	"server/cmd/server/util"
	"time"
)

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")

func CreateSchedule(schedules []scheduleStruct.ScheduleData) {
	updateOperation := make([]mongo.WriteModel, 0)

	for _, schedule := range schedules {
		scheduleOperation := mongo.NewUpdateOneModel()

		scheduleOperation.SetFilter(bson.M{"scheduleId": schedule.ScheduleId})

		startScheduled, _ := util.ConvertTimeStringToDate(schedule.StartScheduled)
		availableAt, _ := util.ConvertTimeStringToDate(schedule.AvailableAt)

		publishedAt, _ := util.ConvertTimeStringToDate(schedule.PublishedAt)
		scheduleOperation.SetUpdate(bson.D{
			{"$set",
				bson.D{
					{"lastUpdated", time.Now()},
					{"scheduleId", schedule.ScheduleId},
					{"title", schedule.Title},
					{"publishedAt", publishedAt},
					{"availableAt", availableAt},
					{"startScheduled", startScheduled},
					{"duration", schedule.Duration},
					{"status", schedule.Status},
					{"channel", schedule.Channel},
				},
			},
		})
		
		scheduleOperation.SetUpsert(true)

		updateOperation = append(updateOperation, scheduleOperation)
	}

	_, err := scheduleCollection.BulkWrite(context.TODO(), updateOperation)

	if err != nil {

		log.Println(err)
		return
	}
}

func GetCurrentSchedule() ([]scheduleStruct.ResponseScheduleData, error) {
	filter := bson.D{{"status", bson.D{{"$in", bson.A{constants.LIVE_STATUS, constants.UPCOMING_STATUS}}}}}

	result, err := scheduleCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var schedule []scheduleStruct.ResponseScheduleData

	err = result.All(context.TODO(), &schedule)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return schedule, nil
}
