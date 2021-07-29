package liveRoom

import (
	"github.com/gorilla/websocket"
	"server/cmd/server/structure/room"
)

func CreateRoom(name string, socket *websocket.Conn) room.RoomData {
	var newChannel room.RoomData
	newChannel.Sockets = append(newChannel.Sockets, socket)
	newChannel.Connections = 1
	newChannel.Name = name
	newChannel.LastTranslation = 0
	return newChannel
}
