package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/cmd/server/eventStore"
	"server/cmd/server/liveRoom"
	"server/cmd/server/models"
	"server/cmd/server/redis"
	"server/cmd/server/routers"
)

func initServer() {
	redis.InitRedisClient()
	models.InitMongoDb()
	eventStore.InitStore()
	liveRoom.InitRoomStore()
}

func initGin(quit chan bool) {
	r := gin.Default()
	r.Use(cors.Default())
	routers.RunGinRouter(r)
	err := r.Run()
	if err != nil {
		quit <- true
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
