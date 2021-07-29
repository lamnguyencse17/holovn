package translation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/cmd/server/liveRoom"
	"server/cmd/server/models"
	"time"
)

type ITranslationStore struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty"`
	LiveId       string  `bson:"liveId,omitempty"`
	Translations []liveRoom.ChatData `bson:"translations,omitempty"`
	LastUpdated  primitive.DateTime  `bson:"last_updated,omitempty"`
}

var translationCollection = models.GetMongoClient().Database("holovn").Collection("translations")

func CreateTranslation(liveId string) {
	initialTranslationStore := ITranslationStore{LiveId: liveId, Translations: make([]liveRoom.ChatData, 0)}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	translationResult, err := translationCollection.InsertOne(ctx, initialTranslationStore)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(translationResult)
}

func InsertToTranslationStore(liveId string, translations []liveRoom.ChatData) {
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
