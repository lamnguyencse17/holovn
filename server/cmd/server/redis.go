package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
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

func SetKeyValue (key string, value string) bool{
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetValue (key string) string{
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println(key, " does not exist")
		return ""
	} else if err != nil {
		log.Println(err)
		return ""
	} else {
		return value
	}
}

func RemoveKey (key string) bool{
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}