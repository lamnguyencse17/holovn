package liveRoom

import (
	"server/cmd/server/event"
)

func ManageChannel(event chan event.ChannelEvent){
	for {
		select {
		case channelEvent := <- event:
			switchChannelEvent(channelEvent)
		default: continue
	}
}
}

func switchChannelEvent(event event.ChannelEvent){
	switch event.Type {
	case "JOIN":
		{
			channelExist := DoesChannelExist(event.Data.LiveId)
			if !channelExist {
				CreateChannel(event.Data.LiveId, event.Data.Socket)
				return
			}
			joinResult := JoinChannel(event.Data.LiveId, event.Data.Socket)
			if !joinResult {
				CreateChannel(event.Data.LiveId, event.Data.Socket)
				return
			}
		}
	}
}