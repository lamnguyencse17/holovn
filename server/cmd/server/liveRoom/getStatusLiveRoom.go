package liveRoom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type StatusType struct {
	Status string `json:"status"`
}

func GetStatusLiveRoom(liveId string)(StatusType , error){
	client := http.DefaultClient
	requestUrl := requestPrefix + liveId

	response, err := client.Get(requestUrl)

	var statusType StatusType

	if err != nil {
		return statusType, err
	}

	body, err := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	if err != nil {
		return statusType, err
	}

	err = json.Unmarshal(body, &statusType)
	if err != nil {
		return statusType, err
	}

	return statusType, nil
}