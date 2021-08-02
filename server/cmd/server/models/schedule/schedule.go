package schedule

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/structure/schedule"
)

type ISchedule struct {
	ID primitive.ObjectID 				`bson:"_id,omitempty" json:"_id"`
	Schedules []schedule.ScheduleData			`bson:"schedules,omitempty" json:"schedules"`
	LastUpdated primitive.DateTime 		`bson:"last_updated,omitempty" json:"last_updated"`
}

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")

func CreateSchedule(schedules []schedule.ScheduleData){

	insertFilter := make([]interface{}, 0)

	for _, v := range schedules {
		insertFilter = append(insertFilter,v)
	}

	opts := options.InsertMany().SetOrdered(false)
	result, err := scheduleCollection.InsertMany(context.TODO(),  insertFilter,opts)

	if err!=nil {
		log.Println(err)
		return
	}
	log.Println(result)
}