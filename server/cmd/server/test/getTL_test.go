package test

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"server/cmd/server/models/translationStore"
	"server/cmd/server/structure/translation"
	"testing"
	"time"
)

func TestGetAllTL(t *testing.T) {
	liveId := "TEST_GET_ALL_TL"
	prepGetTLTestData(liveId)
	// Done prepping

	fetchedTranslations, err := translationStore.GetTranslation(liveId, 0)
	if err != nil {
		return
	}
	assert.Equal(t, liveId, fetchedTranslations.LiveId)
	assert.Equal(t, 2, len(fetchedTranslations.Translations))

	//Cleaning up
	cleanUpGetTLTest(liveId)
}

func TestGetWithTimeStamp(t *testing.T) {
	liveId := "TEST_GET_WITH_TIMESTAMP"
	firstTimestamp := prepGetTLTestData(liveId)
	// Done prepping
	fetchedTranslations, err := translationStore.GetTranslation(liveId, firstTimestamp+20)
	if err != nil {
		return
	}
	assert.Equal(t, liveId, fetchedTranslations.LiveId)
	assert.Equal(t, 1, len(fetchedTranslations.Translations))
	//Cleaning up
	cleanUpGetTLTest(liveId)
}

func prepGetTLTestData(liveId string) (firstTimestamp int64) {
	translationStore.CreateTranslation(liveId)
	var translations []translation.IDatedTranslation
	for i := 0; i <= 1; i = i + 1 {
		currentTime := time.Now()
		if i == 0 {
			firstTimestamp = currentTime.UnixNano()
			time.Sleep(time.Millisecond * 200)
		}
		translations = append(translations, translation.IDatedTranslation{Name: "LAM", Translated: "Translated Message", Timestamp: primitive.NewDateTimeFromTime(currentTime), Original: "Test Message" + string(rune(i))})
	}
	translationStore.InsertToTranslationStore(liveId, translations)
	return firstTimestamp
}

func cleanUpGetTLTest(liveId string) {
	err := translationStore.DeleteTranslation(liveId)
	if err != nil {
		log.Fatal(err)
	}
}
