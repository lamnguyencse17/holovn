package translation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/structure/translation"
	"time"
)

type ITranslationStore struct {
	ID           primitive.ObjectID            `bson:"_id,omitempty"`
	LiveId       string                        `bson:"liveId,omitempty"`
	Translations []translation.TranslationData `bson:"translations,omitempty"`
	LastUpdated  primitive.DateTime            `bson:"last_updated,omitempty"`
}

var database = env.ReadEnv("DatabaseName")
var translationCollection = models.GetMongoClient().Database(database).Collection("translations")

func CreateTranslation(liveId string) {
	initialTranslationStore := ITranslationStore{LiveId: liveId, Translations: make([]translation.TranslationData, 0)}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	translationResult, err := translationCollection.InsertOne(ctx, initialTranslationStore)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(translationResult)
}

func InsertToTranslationStore(liveId string, translations []translation.TranslationData) {
	opts := options.Update()
	updateFilter := bson.D{{"liveId", liveId}}
	updateOperation := bson.D{{"$push", bson.D{{"translations", bson.D{{"$each", translations}}}}}, {"$set", bson.D{{"LastUpdate", time.Now()}}}}
	result, err := translationCollection.UpdateOne(context.TODO(), updateFilter, updateOperation, opts)
	if err != nil {
		log.Println(err)
		return
	}
	if result.MatchedCount == 0 {
		log.Println("NO MATCH FOUND")
		return
	}
	if result.ModifiedCount != 1 {
		log.Println("NONE MODIFIED")
		return
	}
}
