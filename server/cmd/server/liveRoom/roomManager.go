package liveRoom

import (
	"log"
	"server/cmd/server/constants"
	"server/cmd/server/env"
	"server/cmd/server/models/translationStore"
	"server/cmd/server/redis"
	"server/cmd/server/structure/room"
	"server/cmd/server/structure/translation"
	"strconv"
	"time"
)

func pollRoom(roomData room.RoomData) {
	isPullingInstance := isGetTlInstance(roomData.Name)
	defer log.Println("NOT LONGER EXIST")
	for range time.Tick(time.Second * 20) {
		if !DoesRoomExist(roomData.Name) {
			return
		}
		log.Println("POLLING ROOM")
		status , _ := GetStatusLiveRoom(roomData.Name)

		if status.Status == constants.PAST_STATUS {
			return
		}

		limit := 10
		if roomData.LastTranslation == 0 && isPullingInstance { // NEWLY CREATED ROOM
			translationStore.CreateTranslation(roomData.Name)
			limit = 10000
		}
		if isPullingInstance {
			roomData = handleAsPullingInstance(roomData, limit)
			continue
		}
		roomData = handleAsGettingInstance(roomData)
	}
}

func handleAsGettingInstance(roomData room.RoomData) (newRoomData room.RoomData) {
	fetchedTranslationStore, err := translationStore.GetTranslation(roomData.Name, roomData.LastTranslation)
	if err != nil {
		log.Println(err)
		return
	}
	parsedChatData := fetchedTranslationStore.Translations
	newestTimeStamp := parsedChatData[len(parsedChatData)-1].Timestamp.Time().Unix()
	announcingData := translation.ConvertDatedTranslationsToAnnouncingTranslations(parsedChatData)
	announceNewData(roomData, announcingData)
	return UpdateRoomLastChat(roomData.Name, newestTimeStamp)
}

func handleAsPullingInstance(roomData room.RoomData, limit int) (newRoomData room.RoomData) {
	chatData, err := GetTl(roomData.Name, limit)

	if err != nil {
		log.Println(err)
		return
	}
	redis.SetKeyValue(roomData.Name+"-pull", env.ReadEnv("SID"))
	newestTimeStamp, _ := strconv.ParseInt(chatData[len(chatData)-1].Timestamp, 10, 64)
	if roomData.LastTranslation < newestTimeStamp {
		filteredChatData := chatData
		if limit != 10000 {
			filteredChatData = filterChatData(chatData, roomData.LastTranslation)
		}
		parsedChatData := translation.ConvertTranslationsToDatedTranslations(filteredChatData)
		translationStore.InsertToTranslationStore(roomData.Name, parsedChatData)
		announcingData := translation.ConvertDatedTranslationsToAnnouncingTranslations(parsedChatData)
		announceNewData(roomData, announcingData)
		return UpdateRoomLastChat(roomData.Name, newestTimeStamp)
	}
	return roomData
}

func announceNewData(roomData room.RoomData, chatData []translation.IAnnouncingTranslation) {
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

func filterChatData(chatData []translation.TranslationData, timestamp int64) (filteredChatData []translation.TranslationData) {
	for _, chat := range chatData {
		chatTimestamp, _ := strconv.ParseInt(chat.Timestamp, 10, 64)
		if chatTimestamp > timestamp {
			filteredChatData = append(filteredChatData, chat)
		}
	}
	return filteredChatData
}

func isGetTlInstance(liveId string) (pullingStatus bool) {
	pullValue := redis.GetValue(liveId + "-pull")
	return pullValue == "" || pullValue == env.ReadEnv("SID")
}
