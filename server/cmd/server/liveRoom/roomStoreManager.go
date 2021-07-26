package liveRoom

import (
	"server/cmd/server/event"
)

func ManageRoom(event chan event.ChannelEvent){
	for {
		select {
		case channelEvent := <- event:
			switchRoomEvent(channelEvent)
		default: continue
	}
}
}

func switchRoomEvent(event event.ChannelEvent){
	switch event.Type {
	case "JOIN":
		{
			channelExist := DoesRoomExist(event.Data.LiveId)
			if !channelExist {
				newRoomStore, newRoom := CreateRoom(event.Data.LiveId, event.Data.Socket)
				Room.ReplaceRoom(newRoomStore)
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