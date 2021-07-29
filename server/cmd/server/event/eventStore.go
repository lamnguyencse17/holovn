package event

import (
	"server/cmd/server/structure/event"
	"sync"
)

type IEventStore struct {
	mu    sync.Mutex
	store []event.ChannelEvent
}

var Store IEventStore

func InitStore() {
	Store = IEventStore{store: make([]event.ChannelEvent, 0)}
}

func (eventStore *IEventStore) AddEvent(event event.ChannelEvent) {
	eventStore.mu.Lock()
	eventStore.store = append(eventStore.store, event)
	eventStore.mu.Unlock()
}

func (eventStore *IEventStore) DrainEvent() []event.ChannelEvent {
	var drainedStore []event.ChannelEvent
	eventStore.mu.Lock()
	drainedStore = Store.store
	eventStore.store = make([]event.ChannelEvent, 0)
	eventStore.mu.Unlock()
	return drainedStore
}
