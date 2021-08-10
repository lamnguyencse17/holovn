package mainFunctions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"server/cmd/server/httpClient"
	"server/cmd/server/models/scheduleModel"
	"server/cmd/server/structure/scheduleStruct"
)

func VerifySchedules() {
	schedules, err := scheduleModel.GetCurrentSchedule()
	if err != nil {
		return
	}
	liveIdParams := ""
	for _, schedule := range schedules {
		liveIdParams = liveIdParams + "," + schedule.ScheduleId
	}
	client := httpClient.GetHttpClient()
	defer httpClient.DestroyHttpClient()
	response, err := client.Get("https://holodex.net/api/v2/videos?lang=all&sort=available_at&order=desc&limit=100&offset=0&paginated=%3Cempty%3E&id=" + liveIdParams)
	if err != nil {
		return
	}

	var parsedBody verifyResponse
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		log.Println(err)
	}
	scheduleModel.CreateSchedule(parsedBody.Items)
}

type verifyResponse struct {
	Total string                        `json:"total"`
	Items []scheduleStruct.ScheduleData `json:"items"`
}
