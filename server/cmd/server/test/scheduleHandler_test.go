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

func TestInsertScheduleHandler (t *testing.T) {
	scheduleId := "test_scheduleId"
	prepTestInsertScheduleData(scheduleId)
	scheduleResult := getSavedSchedule(scheduleId)

	assert.Equal(t, scheduleResult.ScheduleId, scheduleId)
	assert.Equal(t, scheduleResult.Status, "live")
	assert.Equal(t, scheduleResult.Title,"TestScheduleInsertTitle" )
	assert.Equal(t, scheduleResult.Channel.ChannelId, "ChannelHolo")
}

func TestUpdateMatchScheduleHandler (t *testing.T) {
	scheduleId := "test_scheduleId"
	prepTestUpdateScheduleData(scheduleId)

	scheduleResult := getSavedSchedule(scheduleId)

	assert.Equal(t, scheduleResult.ScheduleId, scheduleId)
	assert.Equal(t, scheduleResult.Status, "stream")
	assert.Equal(t, scheduleResult.Title,"TestScheduleUpdateTitle" )
	assert.Equal(t, scheduleResult.Channel.ChannelId, "ChannelHolo")

	deleteSavedSchedule(scheduleId)
}


func createTestSchedule(newSchedule schedule.ScheduleData) {
	scheduleData := make([]schedule.ScheduleData, 0)

	scheduleData = append(scheduleData, newSchedule)

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

func prepTestInsertScheduleData(scheduleId string){
	newSchedule := schedule.ScheduleData{
		ScheduleId: scheduleId,
		Title: "TestScheduleInsertTitle",
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
	createTestSchedule(newSchedule)
}

func prepTestUpdateScheduleData(scheduleId string){
	newSchedule := schedule.ScheduleData{
		ScheduleId: scheduleId,
		Title: "TestScheduleUpdateTitle",
		PublishedAt: "",
		AvailableAt: "",
		Duration: 0,
		Status: "stream",
		Channel: channel.ChannelData{
			ChannelId: "ChannelHolo",
			Name: "ChannelName",
			Org: "Holovn",Type: "",
			Photo: "",
			EnglishName: "",
		}}
	createTestSchedule(newSchedule)
}