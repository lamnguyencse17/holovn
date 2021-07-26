package liveRoom

import (
	"github.com/gorilla/websocket"
)

func CreateRoom(name string, socket *websocket.Conn) RoomData {
	var newChannel RoomData
	newChannel.sockets = append(newChannel.sockets, socket)
	newChannel.Connections = 1
	newChannel.Name = name
	newChannel.LastChat = 0
	return newChannel
}