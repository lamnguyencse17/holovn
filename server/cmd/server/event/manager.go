package event

import (
	"log"
	"server/cmd/server/structure/event"
	"time"
)

func PollEvents(liveChannel chan event.ChannelEvent) {
	for range time.Tick(time.Millisecond * 100) {
		drainedStore := Store.DrainEvent()
		if len(drainedStore) != 0 {
			log.Println("EVENT QUEUED")
			for _, event := range drainedStore {
				liveChannel <- event
			}
		}
	}
}

func WriteEvent(event event.ChannelEvent) {
	Store.AddEvent(event)
}
