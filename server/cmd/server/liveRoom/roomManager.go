package liveRoom

import (
	"log"
	"server/cmd/server/models/translation"
	"server/cmd/server/structure"
	"strconv"
	"time"
)

func pollRoom(roomData RoomData) {
	defer log.Println("NOT LONGER EXIST")
	for range time.Tick(time.Second * 20) {
		if !DoesRoomExist(roomData.Name) {
			return
		}
		log.Println("POLLING ROOMS")
		limit := 10
		if roomData.LastChat == 0 {
			translation.CreateTranslation(roomData.Name)
			limit = 10000
		}
		chatData, err := GetTl(roomData.Name, limit)
		if err != nil {
			log.Println("GET TL ERROR")
			log.Println(err)
			continue
		}
		newestTimeStamp, _ := strconv.ParseInt(chatData[len(chatData)-1].Timestamp, 10, 64)
		if roomData.LastChat < newestTimeStamp {
			filteredChatData := chatData
			if limit != 10000 {
				filteredChatData = filterChatData(chatData, roomData.LastChat)
			}
			roomData = UpdateRoomLastChat(roomData.Name, newestTimeStamp)
			announceNewData(roomData, filteredChatData)
			translation.InsertToTranslationStore(roomData.Name, filteredChatData)
		}
	}
}

func announceNewData(roomData RoomData, chatData []structure.TranslationData) {
	var newChatData updateChatData
	newChatData.NewChat = chatData
	for _, socket := range roomData.sockets {
		err := socket.WriteJSON(newChatData)
		if err != nil {
			log.Println("SEND DATA ERROR")
			log.Println(err)
			continue
		}
	}
}

func filterChatData(chatData []structure.TranslationData, timestamp int64) (filteredChatData []structure.TranslationData) {
	for _, chat := range chatData {
		chatTimestamp, _ := strconv.ParseInt(chat.Timestamp, 10, 64)
		if chatTimestamp > timestamp {
			filteredChatData = append(filteredChatData, chat)
		}
	}
	return filteredChatData
}
