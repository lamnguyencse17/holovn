package translationStore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/cmd/server/env"
	"server/cmd/server/models"
	"server/cmd/server/structure/translation"
	"server/cmd/server/util"
	"time"
)

type ITranslationStore struct {
	ID           primitive.ObjectID              `bson:"_id,omitempty" json:"_id"`
	LiveId       string                          `bson:"liveId,omitempty" json:"liveId"`
	Translations []translation.IDatedTranslation `bson:"translations,omitempty" json:"translations"`
	LastUpdated  primitive.DateTime              `bson:"lastUpdated,omitempty" json:"lastUpdated"`
}

var database = env.ReadEnv("DatabaseName")
var translationCollection = models.GetMongoClient().Database(database).Collection("translations")

func CreateTranslation(liveId string) {
	initialTranslationStore := ITranslationStore{LiveId: liveId, Translations: make([]translation.IDatedTranslation, 0)}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := translationCollection.InsertOne(ctx, initialTranslationStore)
	if err != nil {
		log.Println(err)
		return
	}
}

func InsertToTranslationStore(liveId string, translations []translation.IDatedTranslation) {
	opts := options.Update()
	updateFilter := bson.D{{"liveId", liveId}}
	updateOperation := bson.D{{"$push", bson.D{{"translations", bson.D{{"$each", translations}}}}}, {"$set", bson.D{{"lastUpdated", time.Now()}}}}
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

func GetTranslation(liveId string, timestamp int64) (returnedTranslationStore ITranslationStore, err error) {
	var requestedTranslation []ITranslationStore
	convertedTimestamp, err := util.ConvertTimestampToPrimitiveDate(timestamp)
	if err != nil {
		return returnedTranslationStore, err
	}
	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunction()
	matchStage := bson.D{{"$match", bson.D{{"liveId", liveId}}}}
	filterCondition := bson.D{{"$gt", bson.A{"$$translation.timestamp", convertedTimestamp}}}
	filterOption := bson.D{{"input", "$translations"}, {"as", "translation"}, {"cond", filterCondition}}
	projectStage := bson.D{{"$project", bson.D{{"liveId", 1}, {"lastUpdated", 1}, {"translations", bson.D{{"$filter", filterOption}}}}}}
	translationCursor, err := translationCollection.Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
	if err != nil {
		log.Println(err)
		return returnedTranslationStore, err
	}
	if err = translationCursor.All(ctx, &requestedTranslation); err != nil {
		panic(err)
	}
	returnedTranslationStore = requestedTranslation[0]
	return returnedTranslationStore, nil
}

func DeleteTranslation(liveId string) error {
	filter := bson.D{{"liveId", liveId}}
	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunction()
	deleteResult, err := translationCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Println("DELETED: ", deleteResult.DeletedCount)
	return nil
}
