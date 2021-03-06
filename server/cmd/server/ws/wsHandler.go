package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server/cmd/server/eventStore"
	"server/cmd/server/structure/eventStruct"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	var welcomeMsg message
	welcomeMsg.MessageType = "info"
	welcomeMsg.Message = "Connected successfully"
	err = conn.WriteJSON(welcomeMsg)
	if err != nil {
		log.Println(err)
		return
	}
	defer handleDisconnection(conn)
	for {
		var payloadEvent wsPayloadEvent
		_, rawPayload, _ := conn.ReadMessage()
		err = json.Unmarshal(rawPayload, &payloadEvent)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(payloadEvent)
		switchEvent(conn, payloadEvent.Event, rawPayload)
	}
}

func handleDisconnection(conn *websocket.Conn) {
	var channelEvent eventStruct.ChannelEvent
	var channelEventData eventStruct.ChannelEventData
	channelEventData.LiveId = "ALL"
	channelEventData.Socket = conn
	channelEvent.Type = "LEAVE_ALL"
	channelEvent.Data = channelEventData
	eventStore.WriteEvent(channelEvent)
}
