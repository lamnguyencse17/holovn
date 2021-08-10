package liveRoom

import (
	"encoding/json"
	"io/ioutil"
	"server/cmd/server/httpClient"
)

type StatusType struct {
	Status string `json:"status"`
}
const requestStatusLivePrefix = "https://holodex.net/api/v2/videos/"

func GetStatusLiveRoom(liveId string) (StatusType, error) {
	client := httpClient.GetHttpClient()
	defer httpClient.DestroyHttpClient()
	requestUrl := requestStatusLivePrefix + liveId

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
