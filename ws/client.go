package ws

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type wsClient struct {
	conn *websocket.Conn
	host string
	path string
	tags []string
}

type wsHandshake struct {
	Ready bool
}

//New Returns a new instance of wsClient
func New() *wsClient {
	return &wsClient{}
}

//Setup Configures a wsClient instance
func (ws *wsClient) Setup(host string, path string, tags []string) {
	ws.host = host
	ws.path = path
	ws.tags = tags

	ws.connect()
	ws.sendHandshake()
}

func (ws *wsClient) connect() {
	u := url.URL{Scheme: "ws", Host: ws.host, Path: ws.path}

	start := time.Now()
	log.Printf("Connecting to %s.", u.String())
	WSConnection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	log.Printf("Connected in %s.", time.Since(start)/time.Nanosecond)

	if err != nil {
		log.Fatalf("Could not connect to %s, reason: %s", u.String(), err)
	}

	ws.conn = WSConnection
}

// func (ws *wsClient) closeConnHandler() func(code int, text string) error {
// 	return ws.conn.handleClose
// }

func (ws *wsClient) writeJSON(data interface{}) {
	ws.conn.WriteJSON(data)
}

func (ws *wsClient) readMessage() (int, []byte, error) {
	return ws.conn.ReadMessage()
}

func (ws *wsClient) sendHandshake() {
	log.Println("Sending initial message.")
	ws.writeJSON(wsHandshake{Ready: true})
	log.Println("Message sent.")
}

// func (ws *wsClient) run() {
// 	log.Println("Starting to listening for messages from the server.")
// 	for {
// 		msgType, msgData, msgErr := ws.readMessage()
// 	}
// }
