package liveRoom

import (
	"github.com/gorilla/websocket"
	"server/cmd/server/structure"
)

type RoomData struct {
	Name        string
	Connections int
	sockets     []*websocket.Conn
	LastChat    int64
}

type updateChatData struct {
	NewChat []structure.TranslationData `json:"newChat"`
}
//
//type ChatData struct {
//	Name      string `json:"name"`
//	Timestamp string `json:"timestamp"`
//	Message   string `json:"message"`
//	IsTl      bool   `json:"is_tl"`
//}
