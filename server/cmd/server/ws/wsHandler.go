package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
	for {
		var payload wsPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			return
		}
		conn.WriteJSON(payload)
	}
}