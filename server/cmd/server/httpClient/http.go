package httpClient

import (
	"net/http"
)

var client *http.Client

func CreateHttpClient() *http.Client {
	//if client != nil {
	//	return client
	//}
	client = http.DefaultClient
	return client
}

func GetHttpClient() *http.Client {
	return client
}

func DestroyHttpClient() {
	client = nil
}
