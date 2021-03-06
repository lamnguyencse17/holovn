package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"server/cmd/server/eventStore"
	"server/cmd/server/structure/eventStruct"
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
		var channelEvent eventStruct.ChannelEvent
		var channelEventData eventStruct.ChannelEventData
		channelEventData.LiveId = joinPayload.Data.LiveId
		channelEventData.Socket = conn
		channelEvent.Type = "JOIN"
		channelEvent.Data = channelEventData
		eventStore.WriteEvent(channelEvent)
	}
}
