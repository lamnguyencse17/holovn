package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/cmd/server/structure/channel"
)

const requestSchedulePrefix = "https://holodex.net/api/v2"
const requestSchedulePostfix = "/live?lang=en&limit=25&max_upcoming_hours="

func getSchedule(maxHours string, limit int) ([]channel.ChannelData, error){
	client := http.DefaultClient
	requestUrl := requestSchedulePrefix + requestSchedulePostfix + maxHours

	resp, err := client.Get(requestUrl)

	var parsedBody []channel.ChannelData

	if err != nil {
		return parsedBody , err
	}

	log.Println("log 1",parsedBody)

	body, err:= ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return parsedBody, err
	}
	log.Println("log 2",parsedBody)

	err = json.Unmarshal(body, &parsedBody)

	if err != nil {
		return parsedBody, err
	}
	log.Println("log 3",parsedBody[0])
	return parsedBody, nil
}
