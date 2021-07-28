package main

import (
	"github.com/gin-gonic/gin"
	"server/cmd/server/event"
	"server/cmd/server/liveRoom"
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
	initRedisClient()
	event.InitStore()
	liveRoom.InitRoomStore()
	liveChannel := make(chan event.ChannelEvent)
	ginChannel := make(chan bool)
	go initGin(ginChannel)
	go event.PollEvents(liveChannel)
	go liveRoom.ManageRoom(liveChannel)
	<-ginChannel
}
