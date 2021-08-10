package main

import (
	"server/cmd/server/eventStore"
	"server/cmd/server/liveRoom"
	"server/cmd/server/mainFunctions"
	"server/cmd/server/structure/eventStruct"
)

func main() {
	initServer()
	liveChannel := make(chan eventStruct.ChannelEvent)
	ginChannel := make(chan bool)
	go initGin(ginChannel)
	go eventStore.PollEvents(liveChannel)
	go liveRoom.ManageRoom(liveChannel)
	go mainFunctions.LoopGetSchedule()
	<-ginChannel
}
