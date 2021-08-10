package scheduleStruct

import (
	"server/cmd/server/structure/channel"
	"time"
)

type ScheduleData struct {
	ScheduleId     string              `bson:"scheduleId" json:"id"`
	Title          string              `json:"title"`
	Type           string              `json:"type"`
	PublishedAt    string              `bson:"publishedAt" json:"published_at"`
	AvailableAt    string              `bson:"availableAt" json:"available_at"`
	StartScheduled string              `bson:"startScheduled" json:"start_scheduled"`
	Duration       int                 `json:"duration"`
	Status         string              `json:"status"`
	Channel        channel.ChannelData `json:"channel"`
}

type ResponseScheduleData struct {
	ScheduleId     string              `bson:"scheduleId" json:"id"`
	Title          string              `json:"title"`
	Type           string              `json:"type"`
	PublishedAt    time.Time           `bson:"publishedAt" json:"publishedAt"`
	AvailableAt    time.Time           `bson:"availableAt" json:"availableAt"`
	StartScheduled time.Time           `bson:"startScheduled" json:"startScheduled"`
	LastUpdated    time.Time           `bson:"lastUpdated" json:"lastUpdated"`
	Duration       int                 `json:"duration"`
	Status         string              `json:"status"`
	Channel        channel.ChannelData `json:"channel"`
}
