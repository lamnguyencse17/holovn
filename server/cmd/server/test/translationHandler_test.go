package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"net/http/httptest"
	"server/cmd/server/models/translationStore"
	"server/cmd/server/redis"
	"server/cmd/server/structure/translation"
	"strconv"
	"testing"
	"time"
)

type translationHandlerResponse struct {
	ObjectId     string                          `json:"_id"`
	LiveId       string                          `json:"liveId"`
	LastUpdated  primitive.DateTime              `json:"lastUpdated"`
	Translations []translation.IDatedTranslation `json:"translations"`
}

func TestTranslationHandler(t *testing.T) {
	liveId := "test_liveId"
	prepData(liveId)
	// First request. Should be in DB
	req, _ := http.NewRequest("GET", "/live/"+liveId+"/translations", nil)
	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 200, responseRecorder.Code)
	var responseObject translationHandlerResponse
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObject)
	if err != nil {
		return
	}
	assert.Equal(t, "test_liveId", responseObject.LiveId)
	assert.Equal(t, "Test Message", responseObject.Translations[0].Original)
	assert.Equal(t, "Translated Message", responseObject.Translations[0].Translated)
	assert.Len(t, responseObject.Translations, 1)
	translations := redis.GetValue(liveId)
	assert.NotEqual(t, "", translations)
	// Clean up
	cleanUp(liveId)
}

func prepData(liveId string) {
	intTime, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		log.Println(err)
	}
	translationStore.CreateTranslation(liveId)
	var translations []translation.IDatedTranslation
	translationEntry := translation.IDatedTranslation{Name: "LAM", Translated: "Translated Message", Timestamp: primitive.NewDateTimeFromTime(time.Unix(intTime, 0)), Original: "Test Message"}
	translations = append(translations, translationEntry)
	translationStore.InsertToTranslationStore(liveId, translations)
}

func cleanUp(liveId string) {
	err := translationStore.DeleteTranslation(liveId)
	if err != nil {
		log.Fatal(err)
	}
	redis.RemoveKey(liveId)
}
