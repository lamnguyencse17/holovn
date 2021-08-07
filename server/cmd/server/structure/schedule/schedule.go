package schedule

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/cmd/server/structure/channel"
)

type ScheduleData struct {
	ScheduleId string `bson:"scheduleId" json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	PublishedAt string  `bson:"publishedAt" json:"published_at"`
	AvailableAt string `bson:"availableAt" json:"available_at"`
	StartScheduled string `bson:"startScheduled" json:"start_scheduled"`
	Duration int `json:"duration"`
	Status string `json:"status"`
	Channel channel.ChannelData `json:"channel"`
 }

type ResponseScheduleData struct {
	ScheduleId string `bson:"scheduleId" json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	PublishedAt primitive.DateTime  `bson:"publishedAt" json:"published_at"`
	AvailableAt primitive.DateTime `bson:"availableAt" json:"available_at"`
	StartScheduled primitive.DateTime `bson:"startScheduled" json:"start_scheduled"`
	Duration int `json:"duration"`
	Status string `json:"status"`
	Channel channel.ChannelData `json:"channel"`
}