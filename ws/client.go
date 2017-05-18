package ws

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

//WebsocketClient This struct is used for holding the main method for the agent
type WebsocketClient struct {
	conn *websocket.Conn
	host string
	path string
	tags []string
}

type wsHandshake struct {
	Ready bool
}

//New Returns a new instance of WebsocketClient
func New() *WebsocketClient {
	return &WebsocketClient{}
}

//Setup Configures a WebsocketClient instance
func (ws *WebsocketClient) Setup(
	host string,
	path string,
	tags []string,
    handlerMap map) {
	ws.host = host
	ws.path = path
	ws.tags = tags

	ws.connect()
	ws.sendHandshake()
}

func (ws *WebsocketClient) connect() {
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

// func (ws *WebsocketClient) closeConnHandler() func(code int, text string) error {
// 	return ws.conn.handleClose
// }

func (ws *WebsocketClient) writeJSON(data interface{}) {
	ws.conn.WriteJSON(data)
}

func (ws *WebsocketClient) readJSON(data interface{}) error {
	return ws.conn.ReadJSON(data)
}

func (ws *WebsocketClient) readMessage() (int, []byte, error) {
	return ws.conn.ReadMessage()
}

func (ws *WebsocketClient) sendHandshake() {
	log.Println("Sending initial message.")
	ws.writeJSON(wsHandshake{Ready: true})
	log.Println("Message sent.")
}

func (ws *WebsocketClient) handleMessage(message map[string]string) error {

}

func (ws *WebsocketClient) run() error {
	log.Println("Starting to listening for messages from the server.")
	for {
		message := make(map[string]string)

		if err := ws.readJSON(message); err != nil {
			log.Println("Got error reading the message: ", err)
			return err
		}

		if err := ws.handleMessage(message); err != nil {
			log.Println("Got error processing the message: ", err)
			return err
		}
	}
}
