package mainFunctions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"server/cmd/server/httpClient"
	"server/cmd/server/models/scheduleModel"
	"server/cmd/server/structure/scheduleStruct"
	"time"
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
	response, err := client.Get("https://holodex.net/api/v2/live?status=past%2C%20missing&lang=all&sort=available_at&order=desc&limit=100&offset=0&paginated=%3Cempty%3E&id=" + liveIdParams)
	if err != nil {
		log.Println(err)
		return
	}
	var parsedBody verifyResponse
	if err != nil {
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &parsedBody)
	log.Println(parsedBody)
	if err != nil {
		log.Println(err)
		return
	}
	scheduleModel.CreateSchedule(parsedBody.Items)
}

func VerifySchedulesRoutine() {
	ticker := time.NewTicker(5 * time.Minute)
	for _ = range ticker.C {
		VerifySchedules()
	}
}

type verifyResponse struct {
	Total string                        `json:"total"`
	Items []scheduleStruct.ScheduleData `json:"items"`
}
