package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"server/cmd/server/env"
)

var ctx = context.Background()

var rdb *redis.Client

func InitRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     env.ReadEnv("RedisAddress"),
		Username: env.ReadEnv("RedisUsername"),
		Password: env.ReadEnv("RedisPassword"),
		DB:       0,
	})
}

func SetKeyValue(key string, value string) bool {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetValue(key string) string {
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

func RemoveKey(key string) bool {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
