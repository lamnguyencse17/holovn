package liveRoom

import (
	"log"
	"server/cmd/server/structure/eventStruct"
)

func ManageRoom(event chan eventStruct.ChannelEvent) {
	for {
		select {
		case channelEvent := <-event:
			log.Println("FOUND EVENT")
			switchRoomEvent(channelEvent)
		default:
			continue
		}
	}
}

func switchRoomEvent(event eventStruct.ChannelEvent) {
	switch event.Type {
	case "LEAVE_ALL":
		{
			LeaveAllRoom(event.Data.Socket)
			removeEmptyRoom()
		}
	case "JOIN":
		{
			channelExist := DoesRoomExist(event.Data.LiveId)
			if !channelExist {
				newRoom := CreateRoom(event.Data.LiveId, event.Data.Socket)
				AddRoom(newRoom)
				go pollRoom(newRoom)
				return
			}
			joinResult := JoinRoom(event.Data.LiveId, event.Data.Socket)
			if !joinResult {
				CreateRoom(event.Data.LiveId, event.Data.Socket)
				return
			}
		}
	}
}
