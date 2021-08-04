package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"server/cmd/server/env"
	"server/cmd/server/models"
	schedule2 "server/cmd/server/models/schedule"
	"server/cmd/server/structure/channel"
	"server/cmd/server/structure/schedule"
	"testing"
)

var database = env.ReadEnv("DatabaseName")
var scheduleCollection = models.GetMongoClient().Database(database).Collection("schedules")

func TestScheduleHandler (t *testing.T) {
	scheduleId := "test_scheduleId"
	prepTestScheduleData(scheduleId)
	scheduleResult := getSavedSchedule(scheduleId)

	assert.Equal(t, scheduleResult.ScheduleId, scheduleId)
	assert.Equal(t, scheduleResult.Status, "live")
	assert.Equal(t, scheduleResult.Title,"TestScheduleTitle" )
	assert.Equal(t, scheduleResult.Channel.ChannelId, "ChannelHolo")

	deleteSavedSchedule(scheduleId)
}

func prepTestScheduleData(scheduleId string){
	scheduleData := make([]schedule.ScheduleData, 0)

	schedule := schedule.ScheduleData{
		ScheduleId: scheduleId,
		Title: "TestScheduleTitle",
		PublishedAt: "",
		AvailableAt: "",
		Duration: 0,
		Status: "live",
		Channel: channel.ChannelData{
			ChannelId: "ChannelHolo",
			Name: "ChannelName",
			Org: "Holovn",Type: "",
			Photo: "",
			EnglishName: "",
		}}

	scheduleData = append(scheduleData, schedule)

	schedule2.CreateSchedule(scheduleData)
}

func getSavedSchedule(scheduleId string) schedule.ScheduleData {

	findFilter := bson.M{"scheduleId": scheduleId }

	var result schedule.ScheduleData

	err := scheduleCollection.FindOne(context.TODO(), findFilter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func deleteSavedSchedule (scheduleId string) {
	deleteFilter := bson.M{"scheduleId": scheduleId }
	_,err := scheduleCollection.DeleteOne(context.TODO(), deleteFilter)
	if err != nil {
		log.Fatal(err)
	}
}