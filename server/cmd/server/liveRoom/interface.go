package liveRoom

import "github.com/gorilla/websocket"

type ChannelData struct {
	Name string
	Connections int
	sockets []*websocket.Conn
}