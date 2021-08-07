package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"server/cmd/server/constants"
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

func TestGetCurrentScheduleHandler(t *testing.T){
	assert.Condition(t, checkCurrentStatus, "Found wrong status")
}

func checkCurrentStatus() bool {
	scheduleData, _ := schedule2.GetCurrentSchedule()

	for _, schedule:=range scheduleData {
		if schedule.Status != constants.LIVE_STATUS && schedule.Status != constants.UPCOMING_STATUS {
			return false
		}
	}

	return true
}

func createTestSchedule(newSchedule schedule.ScheduleData) {
	scheduleData := make([]schedule.ScheduleData, 0)

	scheduleData = append(scheduleData, newSchedule)

	schedule2.CreateSchedule(scheduleData)
}

func  getSavedSchedule(scheduleId string) schedule.ResponseScheduleData {

	findFilter := bson.M{"scheduleId": scheduleId }

	var result schedule.ResponseScheduleData

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
		PublishedAt:"2021-08-02T07:49:21.000Z",
		AvailableAt: "2021-08-02T07:49:21.000Z",
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
		PublishedAt: "2021-08-02T07:49:21.000Z",
		AvailableAt: "2021-08-02T07:49:21.000Z",
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

