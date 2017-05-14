package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/migueleliasweb/pingo-agent/ws"
)

func main() {
	conn := ws.Connect("sockb.in", "/repeat/5")

	go func() {
		conn.WriteMessage(websocket.TextMessage, []byte("FOOOOO"))
		time.Sleep(1500 * time.Millisecond)
		conn.WriteMessage(websocket.TextMessage, []byte("BARRR"))
	}()

	for {
		_, msg, err := conn.ReadMessage()
		log.Println("Err:", err)
		log.Println("Message:", string(msg))
	}
}
