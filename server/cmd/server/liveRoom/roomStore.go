package liveRoom

import (
	"log"
	"server/cmd/server/structure/room"
	"sync"

	"github.com/gorilla/websocket"
)

type IRoomStore struct {
	mu        sync.Mutex
	store     map[string]room.RoomData
	emptyRoom []string
}

var Room IRoomStore

func InitRoomStore() {
	Room = IRoomStore{store: make(map[string]room.RoomData), emptyRoom: make([]string, 0)}
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
		for _, socket := range roomData.Sockets {
			if conn != socket {
				newSockets = append(newSockets, socket)
			}
		}
		roomData.Connections = roomData.Connections - 1
		newRoom := Room.store[name]
		newRoom.Sockets = newSockets
		Room.store[name] = newRoom
		log.Println(Room.store[name].Sockets)
	}
	Room.mu.Unlock()
}

func AddRoom(roomData room.RoomData) {
	Room.mu.Lock()
	Room.store[roomData.Name] = roomData
	Room.mu.Unlock()
}

func UpdateRoomLastChat(name string, lastChat int64) room.RoomData {
	Room.mu.Lock()
	selectedRoom := Room.store[name]
	selectedRoom.LastTranslation = lastChat
	Room.store[name] = selectedRoom
	Room.mu.Unlock()
	return selectedRoom
}

func removeEmptyRoom() {
	Room.mu.Lock()
	for _, room := range Room.store {
		if len(room.Sockets) == 0 {
			Room.emptyRoom = append(Room.emptyRoom, room.Name)
			removeRoomWithoutMutex(room.Name)
		}
	}
	Room.mu.Unlock()
}

func JoinRoom(name string, conn *websocket.Conn) bool {
	Room.mu.Lock()
	if thisRoom, ok := Room.store[name]; ok {
		thisRoom.Sockets = append(Room.store[name].Sockets, conn)
		thisRoom.Connections = thisRoom.Connections + 1
		Room.mu.Unlock()
		return true
	}
	return false
}
