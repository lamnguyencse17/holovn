package room

import (
	"github.com/gorilla/websocket"
	"server/cmd/server/structure/translation"
)

type RoomData struct {
	Name            string
	Connections     int
	Sockets         []*websocket.Conn
	LastTranslation int64
}

type UpdateTranslationData struct {
	NewTranslation []translation.IAnnouncingTranslation `json:"newTranslation"`
}
