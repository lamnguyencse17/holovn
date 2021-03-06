package liveRoom

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/cmd/server/structure/translation"
	"strconv"
)

const requestPrefix = "https://holodex.net/api/v2/videos/"
const requestPostfix = "/chats?lang=en&verified=1&moderator=1&limit="

func GetTl(liveId string, limit int) ([]translation.TranslationData, error) {
	client := http.DefaultClient
	requestUrl := requestPrefix + liveId + requestPostfix + strconv.Itoa(limit)
	log.Println(requestUrl)
	resp, err := client.Get(requestUrl)
	var parsedBody []translation.TranslationData
	if err != nil {
		return parsedBody, err
	}
	body, err := ioutil.ReadAll(resp.Body)
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
