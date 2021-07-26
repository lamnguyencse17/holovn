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