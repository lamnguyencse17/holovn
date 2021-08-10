package test

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"server/cmd/server/httpClient"
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
	res, _ := client.Get("https://holodex.net/api/v2/videos/" + liveId)
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t,  `{"status":"past"}`,string(body)[:17])
	assert.Equal(t, "200 OK", res.Status)
	assert.Equal(t, true, gock.IsDone())
}
