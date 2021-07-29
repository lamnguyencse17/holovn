package channel

type ChannelData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Org string `json:"org"`
	Type string `json:"type"`
	Photo string `json:"photo"`
	EnglishName string `json:"english_name"`
}