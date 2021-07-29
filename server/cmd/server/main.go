package main

import (
	"server/cmd/server/event"
	"server/cmd/server/liveRoom"
	"server/cmd/server/models"
	"server/cmd/server/redis"

	"github.com/gin-gonic/gin"
)

func initGin(quit chan bool) {
	r := gin.Default()
	runGinRouter(r)
	err := r.Run()
	if err != nil {
		quit <- true
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func main() {
	redis.InitRedisClient()
	models.InitMongoDb()
	event.InitStore()
	liveRoom.InitRoomStore()
	liveChannel := make(chan event.ChannelEvent)
	ginChannel := make(chan bool)
	go initGin(ginChannel)
	go event.PollEvents(liveChannel)
	go liveRoom.ManageRoom(liveChannel)
	go getSchedule("25", 12)
	<-ginChannel
}
