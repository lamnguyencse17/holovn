package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)


var ctx = context.Background()

var rdb *redis.Client

func initRedisClient (){
	rdb = redis.NewClient(&redis.Options{
		Addr:     ReadEnv("RedisAddress"),
		Username: ReadEnv("RedisUsername"),
		Password: ReadEnv("RedisPassword"),
		DB:       0,
	})
}