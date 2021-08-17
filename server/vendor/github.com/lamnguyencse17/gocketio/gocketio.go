package gocketio

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func establishConnection(scheme string, host string, path string, rawQuery string) *websocket.Conn {
	u := url.URL{Scheme: scheme, Host: host, Path: path, RawQuery: rawQuery}
	log.Printf("connecting to %s\n", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return c
}

func StartSocketIO(scheme string, host string, path string, rawQuery string) {
	connection := establishConnection(scheme, host, path, rawQuery)
	defer connection.Close()
	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		parsedMessage := parseMessage(message)
		switch parsedMessage.code {
		case SOCKETIO_OPEN:
			{
				log.Println("CONNECTION ESTABLISHING")
				connectMessage := []byte(SOCKETIO_CONNECT)
				err := connection.WriteMessage(WS_MESSAGE_TYPE, connectMessage)
				if err != nil {
					log.Println(err)
				}
			}
		case SOCKETIO_CONNECT:
			{
				log.Println("CONNECTION ESTABLISHED")
				subscribePayload, err := json.Marshal(ISubscribePayload{VideoId: "k_oMkblkB9k", Lang: "en"})
				if err != nil {
					return
				}
				subscribeMessage := wrapMessage(SOCKETIO_EMIT, EVENT_SUBSCRIBE, string(subscribePayload))
				err = connection.WriteMessage(WS_MESSAGE_TYPE, subscribeMessage)
				if err != nil {
					log.Println(err)
				}
			}
		case SOCKETIO_EMIT:
			{
				log.Println("EMIT SUCCESS")
			}
		case SOCKETIO_PING:
			{
				pongSocketIO(connection)
			}
		default:
			return
		}
	}
}

func pongSocketIO(connection *websocket.Conn) {
	pongMessage := []byte(SOCKETIO_PONG)
	err := connection.WriteMessage(WS_MESSAGE_TYPE, pongMessage)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("PONGED")
}
