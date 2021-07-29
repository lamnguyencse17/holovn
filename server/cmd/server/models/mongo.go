package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"server/cmd/server/env"
	"time"
)

var mongoClient *mongo.Client

func InitMongoDb() {
	uri := env.ReadEnv("MongoUri")
	if uri == "" {
		log.Fatalln("MongoDB URI is not defined")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MONGODB CONNECTED")
	mongoClient = client
	CreateTranslationStoreIndex()
}

func CreateTranslationStoreIndex() {
	translationCollection := mongoClient.Database("holovn").Collection("translations")
	_, err := translationCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "liveId", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Panicln(err)
	}
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		InitMongoDb()
	}
	return mongoClient
}
