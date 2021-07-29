package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"server/cmd/server/event"
	event2 "server/cmd/server/structure/event"
)

func switchEvent(conn *websocket.Conn, eventName string, rawPayload []byte) {
	switch eventName {
	case "JOIN":
		var joinPayload wsJoinPayload
		err := json.Unmarshal(rawPayload, &joinPayload)
		if err != nil {
			log.Println(err)
			return
		}
		var channelEvent event2.ChannelEvent
		var channelEventData event2.ChannelEventData
		channelEventData.LiveId = joinPayload.Data.LiveId
		channelEventData.Socket = conn
		channelEvent.Type = "JOIN"
		channelEvent.Data = channelEventData
		event.WriteEvent(channelEvent)
	}
}
