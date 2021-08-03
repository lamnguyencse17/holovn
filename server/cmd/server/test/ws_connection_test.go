package test

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"server/cmd/server/ws"
	"strings"
	"testing"
)

func TestWsConnection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ws.HandleWS))
	defer server.Close()
	url := "ws" + strings.TrimPrefix(server.URL, "http")
	socket, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer socket.Close()
	for {
		_, payload, err := socket.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}
		assert.Equal(t, "{\"type\":\"info\",\"message\":\"Connected successfully\"}\n", string(payload))
		return
	}
}
