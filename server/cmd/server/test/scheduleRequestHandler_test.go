package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestGetCurrentSchedule(t *testing.T){
	req,_ := http.NewRequest("GET", "/schedules/current", nil)

	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder,req)

	assert.Equal(t, 200, responseRecorder.Code)
}
