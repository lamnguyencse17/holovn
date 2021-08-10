package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/cmd/server/env"
	"server/cmd/server/models/scheduleModel"
	"server/cmd/server/structure/scheduleStruct"
	"time"
)

const requestSchedulePrefix = "https://holodex.net/api/v2"
const requestSchedulePostfix = "/live?lang=en&max_upcoming_hours=48&org=Hololive"

var apiKey = env.ReadEnv("HolodexKey")

func getSchedule() ([]scheduleStruct.ScheduleData, error) {
	client := http.Client{}
	requestUrl := requestSchedulePrefix + requestSchedulePostfix
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.Header.Set("x-api-key", apiKey)
	resp, err := client.Do(request)

	var parsedBody []scheduleStruct.ScheduleData

	if err != nil {
		return parsedBody, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return parsedBody, err
	}

	err = json.Unmarshal(body, &parsedBody)
	scheduleModel.CreateSchedule(parsedBody)
	if err != nil {
		return parsedBody, err
	}

	return parsedBody, nil
}

func loopGetSchedule() {
	ticker := time.NewTicker(5 * time.Minute)
	for _ = range ticker.C {
		getSchedule()
	}
}
