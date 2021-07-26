package liveRoom


import "sync"

type RoomStore struct {
	mu sync.Mutex
	store []RoomData
}

var Room RoomStore

func InitRoom(){
	Room = RoomStore{store: make([]RoomData, 0)}
}

func (roomStore *RoomStore) ReplaceRoom(newRoomStore []RoomData){
	roomStore.mu.Lock()
	roomStore.store = newRoomStore
	roomStore.mu.Unlock()
}

func UpdateLastChatInRoom(name string, lastChat int64) RoomData{
	var updatedRoom RoomData
	Room.mu.Lock()
	for roomIndex, room := range Room.store{
		if room.Name == name{
			Room.store[roomIndex].LastChat = lastChat
			updatedRoom = Room.store[roomIndex]
			break
		}
	}
	Room.mu.Unlock()
	return updatedRoom
}