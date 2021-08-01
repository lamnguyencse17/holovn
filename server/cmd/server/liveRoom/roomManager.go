package liveRoom

import (
	"log"
	"server/cmd/server/models/translation"
	"server/cmd/server/redis"
	"server/cmd/server/structure/room"
	translation2 "server/cmd/server/structure/translation"
	"strconv"
	"time"
)

func pollRoom(roomData room.RoomData) {
	defer log.Println("NOT LONGER EXIST")
	for range time.Tick(time.Second * 20) {
		if !DoesRoomExist(roomData.Name) {
			return
		}
		log.Println("POLLING ROOMS")
		limit := 10
		if roomData.LastTranslation == 0 {
			translation.CreateTranslation(roomData.Name)
			limit = 10000
		}
		pullingStatus := redis.GetValue(roomData.Name + "-pull")

		var chatData []translation2.TranslationData
		var err error

		if pullingStatus == "" {
			chatData, err = GetTl(roomData.Name, limit)

			if err != nil {
				log.Println("GET TL ERROR")
				log.Println(err)
				continue
			}
			redis.SetKeyValue(roomData.Name + "-pull", "pulled")

		} else {
			log.Println("Get translation from DB")
		}
		newestTimeStamp, _ := strconv.ParseInt(chatData[len(chatData)-1].Timestamp, 10, 64)
		if roomData.LastTranslation < newestTimeStamp {
			filteredChatData := chatData
			if limit != 10000 {
				filteredChatData = filterChatData(chatData, roomData.LastTranslation)
			}
			roomData = UpdateRoomLastChat(roomData.Name, newestTimeStamp)
			announceNewData(roomData, filteredChatData)
			if pullingStatus == "" {
				translation.InsertToTranslationStore(roomData.Name, filteredChatData)
			}
		}
	}
}

func announceNewData(roomData room.RoomData, chatData []translation2.TranslationData) {
	var newChatData room.UpdateTranslationData
	newChatData.NewTranslation = chatData
	for _, socket := range roomData.Sockets {
		err := socket.WriteJSON(newChatData)
		if err != nil {
			log.Println("SEND DATA ERROR")
			log.Println(err)
			continue
		}
	}
}

func filterChatData(chatData []translation2.TranslationData, timestamp int64) (filteredChatData []translation2.TranslationData) {
	for _, chat := range chatData {
		chatTimestamp, _ := strconv.ParseInt(chat.Timestamp, 10, 64)
		if chatTimestamp > timestamp {
			filteredChatData = append(filteredChatData, chat)
		}
	}
	return filteredChatData
}


