package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"server/cmd/server/models/schedule"
	"testing"
)

func TestRequestGetCurrentSchedule(t *testing.T){
	req,_ := http.NewRequest("GET", "/schedules/current", nil)

	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder,req)

	scheduleData, _ := schedule.GetCurrentSchedule()

	assert.Equal(t, 200, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body, "Body is not empty")

	parsedScheduledData, _ := json.Marshal(scheduleData)

	assert.ElementsMatch(t, responseRecorder.Body.Bytes(),parsedScheduledData)

}
