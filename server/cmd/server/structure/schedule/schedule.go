package schedule

import "server/cmd/server/structure/channel"

type ScheduleData struct {
	ScheduleId string `bson:"scheduleId" json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	PublishedAt string `bson:"publishedAt" json:"publishedAt"`
	AvailableAt string `bson:"availableAt" json:"availableAt"`
	Duration int `json:"duration"`
	Status string `json:"status"`
	Channel channel.ChannelData `json:"channel"`
 }