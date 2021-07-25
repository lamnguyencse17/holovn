package ws


type message struct {
	MessageType string `json:"type"`
	Message     string `json:"message"`
}

type wsPayload struct {
	Event string `json:"event"`
	Data map[string]interface{} `json:"data"`
}