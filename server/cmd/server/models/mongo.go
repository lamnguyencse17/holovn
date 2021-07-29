package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"server/cmd/server/env"
	"time"
)

var mongoClient *mongo.Client

func InitMongoDb(){
	uri := env.ReadEnv("MongoUri")
	if uri == ""{
		log.Fatalln("MongoDB URI is not defined")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
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
	cancelFunc()
}

func GetMongoClient () *mongo.Client {
	if mongoClient == nil{
		InitMongoDb()
	}
	return mongoClient
}
