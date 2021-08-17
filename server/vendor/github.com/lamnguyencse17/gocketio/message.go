package gocketio

import (
	"strconv"
)

type IParsedMessage struct {
	code string
	data string
}

const WS_MESSAGE_TYPE = 1

func parseMessage(message []byte) (parsedMessage IParsedMessage) {
	messageString := string(message)
	if len(messageString) == 1 {
		return IParsedMessage{code: string(messageString)}
	}
	_, err := strconv.Atoi(string(messageString[1]))
	if err != nil {
		return IParsedMessage{code: string(messageString[0]), data: messageString[1:]}
	}
	return IParsedMessage{code: messageString[:2], data: messageString[3:]}
}

func wrapMessage(socketIOEvent string, event string, data string) (wrappedMessage []byte) {
	return []byte(socketIOEvent + "[" + "\"" + event + "\"" + "," + data + "]")
}
