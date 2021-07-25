package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const requestPrefix = "https://holodex.net/api/v2/videos/"
const requestPostfix = "/chats?lang=en&verified=1&moderator=1&limit=10000"

func getTl(liveId string) ([]chatData, error){
	client := http.DefaultClient
	requestUrl := requestPrefix + liveId + requestPostfix
	resp, err := client.Get(requestUrl)
	var parsedBody []chatData
	if err != nil {
		log.Println(err)
		return parsedBody, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return parsedBody, err
	}
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		log.Println(err)
		return parsedBody, err
	}
	return parsedBody, nil
}
