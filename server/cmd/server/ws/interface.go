package ws

type message struct {
	MessageType string `json:"type"`
	Message     string `json:"message"`
}

type wsPayloadEvent struct {
	Event string `json:"eventStore"`
}

type wsJoinPayload struct {
	Event string   `json:"eventStore"`
	Data  joinData `json:"data"`
}

type joinData struct {
	LiveId string `json:"liveId"`
}
