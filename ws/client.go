package ws

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func Connect(host string, path string) *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	log.Printf("Connecting to %s", u.String())

	WSConnection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatalf("Could not connect to %s, reason: %s", u.String(), err)
	}

	return WSConnection
}
