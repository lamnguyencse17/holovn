package liveRoom

import (
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

func DoesRoomExist(name string) bool {
	for _, channel := range Room.store {
		if channel.Name == name{
			return true
		}
	}
	return false
}

func CreateRoom(name string, socket *websocket.Conn) ([]RoomData, RoomData)  {
	var newChannel RoomData
	for _, channel := range Room.store {
		if channel.Name == name {
			return nil, newChannel
		}
	}
	var newChannelList = make([]RoomData, len(Room.store)+1)
	newChannel.sockets = append(newChannel.sockets, socket)
	newChannel.Connections = 1
	newChannel.Name = name
	newChannel.LastGet = 0
	newChannelList = append(Room.store, newChannel)
	return newChannelList, newChannel
}

func GetRoom(name string) (RoomData, error){
	var foundChannel RoomData
	for _, channel := range Room.store {
		if channel.Name == name {
			channel.LastGet = time.Now().Unix()
			return foundChannel, nil
		}
	}
	return foundChannel, errors.New("NOT FOUND")
}