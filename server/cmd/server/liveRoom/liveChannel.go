package liveRoom

import "github.com/gorilla/websocket"

func DoesChannelExist(name string) bool {
	for _, channel := range ChannelList {
		if channel.Name == name{
			return true
		}
	}
	return false
}

func CreateChannel(name string, socket *websocket.Conn) ([]ChannelData, ChannelData)  {
	var newChannel ChannelData
	for _, channel := range ChannelList {
		if channel.Name == name {
			return nil, newChannel
		}
	}
	var newChannelList = make([]ChannelData, len(ChannelList)+1)
	newChannel.sockets = append(newChannel.sockets, socket)
	newChannel.Connections = 1
	newChannel.Name = name
	newChannelList = append(ChannelList, newChannel)
	return newChannelList, newChannel
}

func JoinChannel(name string, socket *websocket.Conn) bool {
	for _, channel := range ChannelList {
		if channel.Name == name {
			channel.sockets = append(channel.sockets, socket)
			return true
		}
	}
	return false
}

func RemoveChannel(name string) {
	var newChannelList = make([]ChannelData, len(ChannelList)-1)
	for _, channel := range ChannelList {
		if channel.Name != name {
			newChannelList = append(newChannelList, channel)
		}
	}
}