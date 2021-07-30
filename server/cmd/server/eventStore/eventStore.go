package eventStore

import (
	"server/cmd/server/structure/eventStruct"
	"sync"
)

type IEventStore struct {
	mu    sync.Mutex
	store []eventStruct.ChannelEvent
}

var Store IEventStore

func InitStore() {
	Store = IEventStore{store: make([]eventStruct.ChannelEvent, 0)}
}

func (eventStore *IEventStore) AddEvent(event eventStruct.ChannelEvent) {
	eventStore.mu.Lock()
	eventStore.store = append(eventStore.store, event)
	eventStore.mu.Unlock()
}

func (eventStore *IEventStore) DrainEvent() []eventStruct.ChannelEvent {
	var drainedStore []eventStruct.ChannelEvent
	eventStore.mu.Lock()
	drainedStore = Store.store
	eventStore.store = make([]eventStruct.ChannelEvent, 0)
	eventStore.mu.Unlock()
	return drainedStore
}
