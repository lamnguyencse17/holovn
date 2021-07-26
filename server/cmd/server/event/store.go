package event

import "sync"

type EventStore struct {
	mu sync.Mutex
	store []ChannelEvent
}

var Store EventStore

func InitStore(){
	Store = EventStore{store: make([]ChannelEvent, 0)}
}

func (eventStore *EventStore) AddEvent(event ChannelEvent){
	eventStore.mu.Lock()
	eventStore.store = append(eventStore.store, event)
	eventStore.mu.Unlock()
}

func (eventStore *EventStore) DrainEvent() []ChannelEvent{
	var drainedStore []ChannelEvent
	eventStore.mu.Lock()
	drainedStore = Store.store
	eventStore.store = make([]ChannelEvent, 0)
	eventStore.mu.Unlock()
	return drainedStore
}