package liveRoom

import (
	"log"
	"server/cmd/server/event"
)

func ManageRoom(event chan event.ChannelEvent) {
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

func switchRoomEvent(event event.ChannelEvent) {
	switch event.Type {
	case "LEAVE_ALL":
		{
			log.Println("Helloooo", Room)
			LeaveAllRoom(event.Data.Socket)
			removeEmptyRoom()
			log.Println("Room after disconnect", Room)

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
			//joinResult := JoinRoom(event.Data.LiveId, event.Data.Socket)
			//if !joinResult {
			//	CreateRoom(event.Data.LiveId, event.Data.Socket)
			//	return
			//}
		}
	}
}
