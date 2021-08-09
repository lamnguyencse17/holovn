package test

import (
	"gopkg.in/h2non/gock.v1"
	"server/cmd/server/httpClient"
	"testing"
)

func TestGetLiveStatus(t *testing.T) {
	// SETUP
	liveId := "live_status_test"
	gock.New("https://holodex.net/api/v2/videos/" + liveId).Reply(200).JSON(nil) //Mock response here
	defer gock.Off()
	client := httpClient.CreateHttpClient()
	gock.InterceptClient(client)
	// TEST

}
