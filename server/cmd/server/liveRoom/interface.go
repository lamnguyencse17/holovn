package liveRoom

import "github.com/gorilla/websocket"

type RoomData struct {
	Name string
	Connections int
	sockets []*websocket.Conn
	LastGet int64
}

type updateChatData struct {
	NewChat []ChatData `json:"newChat"`
}

type ChatData struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	IsTl      bool   `json:"is_tl"`
}