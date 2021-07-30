package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/cmd/server/env"
	"server/cmd/server/structure/schedule"
)

const requestSchedulePrefix = "https://holodex.net/api/v2"
const requestSchedulePostfix = "/live?lang=en&max_upcoming_hours=48"

var apiKey = env.ReadEnv("HolodexKey")

func getSchedule() ([]schedule.ScheduleData, error){
	client := http.Client{}
	requestUrl := requestSchedulePrefix + requestSchedulePostfix
	request , _ := http.NewRequest("GET", requestUrl, nil)
	request.Header.Set("x-api-key", apiKey)
	resp, err := client.Do(request)

	var parsedBody []schedule.ScheduleData

	if err != nil {
		return parsedBody , err
	}

	body, err:= ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return parsedBody, err
	}

	err = json.Unmarshal(body, &parsedBody)

	if err != nil {
		return parsedBody, err
	}

	return parsedBody, nil
}
