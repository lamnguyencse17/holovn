package eventStore

import (
	"log"
	"server/cmd/server/structure/eventStruct"
	"time"
)

func PollEvents(liveChannel chan eventStruct.ChannelEvent) {
	for range time.Tick(time.Millisecond * 200) {
		drainedStore := Store.DrainEvent()
		if len(drainedStore) != 0 {
			log.Println("EVENT QUEUED")
			for _, event := range drainedStore {
				liveChannel <- event
			}
		}
	}
}

func WriteEvent(event eventStruct.ChannelEvent) {
	Store.AddEvent(event)
}
