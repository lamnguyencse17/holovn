package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	schedule2 "server/cmd/server/structure/schedule"
	"testing"
	"time"
)

func TestRequestGetCurrentSchedule(t *testing.T){
	prepTestInsertScheduleData("test_scheduleId")

	req,_ := http.NewRequest("GET", "/schedules/current", nil)

	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder,req)

	assert.Equal(t, 200, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body, "Body is not empty")

	var scheduleData []schedule2.ResponseScheduleData

	err := json.Unmarshal(responseRecorder.Body.Bytes(), &scheduleData)

	if err != nil {
		log.Fatal(err)
	}
	scheduleLength := len(scheduleData)

	assert.Equal(t, scheduleData[scheduleLength - 1].ScheduleId, "test_scheduleId" )
	assert.Equal(t,scheduleData[scheduleLength - 1].Status, "live" )
	assert.Equal(t, scheduleData[scheduleLength- 1].Title,"TestScheduleInsertTitle")

	deleteSavedSchedule("test_scheduleId")
}

func TestTypeOfTime(t *testing.T){
	prepTestInsertScheduleData("test_scheduleId")

	req,_ := http.NewRequest("GET", "/schedules/current", nil)

	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder,req)

	var scheduleData []schedule2.ResponseScheduleData

	err := json.Unmarshal(responseRecorder.Body.Bytes(), &scheduleData)
	if err != nil {
		log.Fatal(err)
	}


	currentTime := primitive.NewDateTimeFromTime(time.Now())

	assert.IsType(t, reflect.TypeOf(scheduleData[0].StartScheduled), reflect.TypeOf(currentTime))
	assert.IsType(t, reflect.TypeOf(scheduleData[0].AvailableAt), reflect.TypeOf(currentTime))
	assert.IsType(t, reflect.TypeOf(scheduleData[0].PublishedAt), reflect.TypeOf(currentTime))
}

