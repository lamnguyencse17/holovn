package event

import (
	"time"
)

func PollEvents (liveChannel chan ChannelEvent){
	for range time.Tick(time.Millisecond * 100) {
		if len(Store) != 0{
			for _, event := range Store{
				liveChannel <- event
			}
		}
	}
}

func WriteEvent (event ChannelEvent){}