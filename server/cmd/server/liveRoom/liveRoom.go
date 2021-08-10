package liveRoom

import (
	"github.com/gorilla/websocket"
	"server/cmd/server/structure/roomStruct"
)

func CreateRoom(name string, socket *websocket.Conn) roomStruct.RoomData {
	var newChannel roomStruct.RoomData
	newChannel.Sockets = append(newChannel.Sockets, socket)
	newChannel.Connections = 1
	newChannel.Name = name
	newChannel.LastTranslation = 0
	return newChannel
}
