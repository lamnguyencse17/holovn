package channel

type ChannelData struct {
	ChannelId string `bson:"channelId" json:"id"`
	Name string `json:"name"`
	Org string `json:"org"`
	Type string `json:"type"`
	Photo string `json:"photo"`
	EnglishName string `bson:"englishName" json:"english_name"`
}