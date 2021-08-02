package main

import (
	"github.com/gin-gonic/gin"
	"server/cmd/server/eventStore"
	"server/cmd/server/liveRoom"
	"server/cmd/server/models"
	"server/cmd/server/redis"
	"server/cmd/server/routers"
	"server/cmd/server/structure/eventStruct"
)

func initGin(quit chan bool) {
	r := gin.Default()
	routers.RunGinRouter(r)
	err := r.Run()
	if err != nil {
		quit <- true
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func main() {
	redis.InitRedisClient()
	models.InitMongoDb()
	eventStore.InitStore()
	liveRoom.InitRoomStore()
	liveChannel := make(chan eventStruct.ChannelEvent)
	ginChannel := make(chan bool)
	go initGin(ginChannel)
	go eventStore.PollEvents(liveChannel)
	go liveRoom.ManageRoom(liveChannel)
	go loopGetSchedule()
	<-ginChannel
}
