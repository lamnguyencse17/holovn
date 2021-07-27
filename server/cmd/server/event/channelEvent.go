package event

import "github.com/gorilla/websocket"

type ChannelEvent struct {
	Type string
	Data ChannelEventData
}

type ChannelEventData struct {
	LiveId string
	Socket *websocket.Conn
}
