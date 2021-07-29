package main

import (
	"log"
	"net/http"
)

const requestSchedulePrefix = "https://holodex.net/api/v2"
const requestSchedulePostfix = "/live?lang=en&limit=25&max_upcoming_hours="

func getSchedule(maxHours string, limit int) {
	client := http.DefaultClient
	requestUrl := requestSchedulePrefix + requestSchedulePostfix + maxHours

	resp, err := client.Get(requestUrl)

	log.Println("REsponse", resp)

	log.Println("err", err)
}
