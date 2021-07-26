package liveRoom

import (
	"log"
	"time"
)

func pollRoom (roomData RoomData){
	for range time.Tick(time.Millisecond * 5000) {
		limit := 10
		if roomData.LastGet == 0 {
			limit = 10000
		}
		chatData, err := GetTl(roomData.Name, limit)
		if err != nil {
			log.Println("GET TL ERROR")
			log.Println(err)
			continue
		}
		roomData, err = GetRoom(roomData.Name)
		if err != nil{
			log.Println("GET ROOM ERROR")
			log.Println(err)
			continue
		}
		announceNewData(roomData, chatData)
	}
}

func announceNewData (roomData RoomData, chatData []ChatData){
	var newChatData updateChatData
	newChatData.NewChat = chatData
	for _, socket := range roomData.sockets{
		err := socket.WriteJSON(newChatData)
		if err != nil {
			log.Println("SEND DATA ERROR")
			log.Println(err)
			continue
		}
	}
}