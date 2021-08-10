package test

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"log"
	"server/cmd/server/httpClient"
	"server/cmd/server/liveRoom"
	"testing"
)

func TestGetLiveStatus(t *testing.T) {
	// SETUP
	liveId := "live_status_test"
	gock.New("https://holodex.net/api/v2/videos/" + liveId).Reply(200).JSON(map[string]string{"status": "past"}) //Mock response here
	defer gock.Off()
	client := httpClient.CreateHttpClient()
	gock.InterceptClient(client)
	// TEST
	statusType, err := liveRoom.GetStatusLiveRoom(liveId)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t,  "past",statusType.Status)
	assert.Equal(t, true, gock.IsDone())
}
