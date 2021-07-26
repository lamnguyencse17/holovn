package event

import (
	"log"
	"time"
)

func PollEvents (liveChannel chan ChannelEvent){
	for range time.Tick(time.Millisecond * 100) {
		drainedStore := Store.DrainEvent()
		if len(drainedStore) != 0{
			log.Println("FOUND EVENTS")
			for _, event := range drainedStore{
				liveChannel <- event
			}
		}
	}
}

func WriteEvent (event ChannelEvent){
	Store.AddEvent(event)
}