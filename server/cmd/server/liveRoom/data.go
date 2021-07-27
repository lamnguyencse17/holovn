package liveRoom

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type IRoomStore struct {
	mu        sync.Mutex
	store     map[string]RoomData
	emptyRoom []string
}

var Room IRoomStore

func InitRoomStore() {
	Room = IRoomStore{store: make(map[string]RoomData), emptyRoom: make([]string, 0)}
}

func RemoveRoom(name string) {
	Room.mu.Lock()
	delete(Room.store, name)
	Room.mu.Unlock()
}

func removeRoomWithoutMutex(name string) {
	delete(Room.store, name)
}

func DoesRoomExist(name string) bool {
	Room.mu.Lock()
	if _, ok := Room.store[name]; ok {
		Room.mu.Unlock()
		return true
	}
	Room.mu.Unlock()
	return false
}

func LeaveAllRoom(conn *websocket.Conn) {
	Room.mu.Lock()
	for name, roomData := range Room.store {
		var newSockets []*websocket.Conn
		for _, socket := range roomData.sockets {
			if conn != socket {
				newSockets = append(newSockets, socket)
			}
		}
		newRoom := Room.store[name]
		newRoom.sockets = newSockets
		Room.store[name] = newRoom
		log.Println(Room.store[name].sockets)
	}
	Room.mu.Unlock()
}

func AddRoom(roomData RoomData) {
	Room.mu.Lock()
	Room.store[roomData.Name] = roomData
	Room.mu.Unlock()
}

func UpdateRoomLastChat(name string, lastChat int64) RoomData {
	Room.mu.Lock()
	selectedRoom := Room.store[name]
	selectedRoom.LastChat = lastChat
	Room.store[name] = selectedRoom
	Room.mu.Unlock()
	return selectedRoom
}

func removeEmptyRoom() {
	Room.mu.Lock()
	for _, room := range Room.store {
		if len(room.sockets) == 0 {
			Room.emptyRoom = append(Room.emptyRoom, room.Name)
			removeRoomWithoutMutex(room.Name)
		}
	}
	Room.mu.Unlock()
}