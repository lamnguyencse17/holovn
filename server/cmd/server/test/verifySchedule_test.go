package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/h2non/gock.v1"
	"log"
	"server/cmd/server/constants"
	"server/cmd/server/httpClient"
	"server/cmd/server/mainFunctions"
	"server/cmd/server/models/scheduleModel"
	"server/cmd/server/structure/channel"
	"server/cmd/server/structure/scheduleStruct"
	"testing"
	"time"
)

func TestVerifySchedule(t *testing.T) {
	// SETUP
	// INITIAL DATA
	liveId := "live_status_test"
	timeValue := time.Now().String()
	channelData := channel.ChannelData{ChannelId: "testChannel", Name: "channel", Org: "Holotest", Type: "vtuber", Photo: "nah", EnglishName: ""}
	live := scheduleStruct.ScheduleData{ScheduleId: liveId, Title: "testTitle", Type: "live", PublishedAt: timeValue, AvailableAt: timeValue, StartScheduled: timeValue, Duration: 0, Status: constants.UPCOMING_STATUS, Channel: channelData}
	var liveSlice []scheduleStruct.ScheduleData
	liveSlice = append(liveSlice, live)
	scheduleModel.CreateSchedule(liveSlice)

	// MODIFIED FOR MOCKING
	live.Status = constants.PAST_STATUS
	var modifiedLiveSlice []scheduleStruct.ScheduleData
	modifiedLiveSlice = append(modifiedLiveSlice, live)
	mockResponse := verifyResponse{Items: modifiedLiveSlice, Total: "1"}
	gock.New("https://holodex.net/api/v2/videos?lang=all&sort=available_at&order=desc&limit=100&offset=0&paginated=%3Cempty%3E&id=" + liveId).Reply(200).JSON(mockResponse)
	defer gock.Off()
	client := httpClient.CreateHttpClient()
	gock.InterceptClient(client)
	// TEST
	mainFunctions.VerifySchedules()
	findFilter := bson.M{"scheduleId": liveId}

	var result scheduleStruct.ResponseScheduleData
	err := scheduleCollection.FindOne(context.TODO(), findFilter).Decode(&result)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, "past", result.Status)
	assert.Equal(t, liveId, result.ScheduleId)
	// CLEANUP
	deleteFilter := bson.M{"scheduleId": liveId}
	_, err = scheduleCollection.DeleteOne(context.TODO(), deleteFilter)
	if err != nil {
		log.Println(err)
	}
}

type verifyResponse struct {
	Total string                        `json:"total"`
	Items []scheduleStruct.ScheduleData `json:"items"`
}
