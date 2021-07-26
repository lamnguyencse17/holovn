package liveRoom

import (
	"log"
	"strconv"
	"time"
)

func pollRoom (roomData RoomData){
	for range time.Tick(time.Millisecond * 15000) {
		limit := 10
		if roomData.LastChat == 0 {
			limit = 10000
		}
		chatData, err := GetTl(roomData.Name, limit)
		if err != nil {
			log.Println("GET TL ERROR")
			log.Println(err)
			continue
		}
		newestTimeStamp, _ := strconv.ParseInt(chatData[0].Timestamp, 10, 64)
		if roomData.LastChat < newestTimeStamp {
			roomData = UpdateLastChatInRoom(roomData.Name, newestTimeStamp)
			announceNewData(roomData, chatData)
		}
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