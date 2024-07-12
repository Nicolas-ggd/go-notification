package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type Websocket struct {
	// per client represent map[string]*Client type, each client is provided with key
	Clients map[string]*Client

	Broadcast chan []byte

	Register chan *Client

	UnRegister chan *Client
}

// NewWebsocket returns new Websocket
func NewWebsocket() *Websocket {
	return &Websocket{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
}

func (ws *Websocket) Run() {
	for {
		select {
		// handle register client case
		case client := <-ws.Register:
			ws.Clients[client.ClientId] = client

			p := Message{
				Event: "rand",
				Data:  "PONG",
			}

			// marshal packet and send in to the channel
			symbolByte, err := json.Marshal(p)
			if err != nil {
				log.Println(err)
				return
			}

			client.Send <- symbolByte

		// unregister client case
		case client := <-ws.UnRegister:
			if _, ok := ws.Clients[client.ClientId]; ok {
				// delete client
				close(client.Send)
				delete(ws.Clients, client.ClientId)
			}

		// handle case to receiving broadcast
		case message := <-ws.Broadcast:
			for _, client := range ws.Clients {
				select {
				case client.Send <- message:
					fmt.Println("Broadcasting client.send")
				default:
					close(client.Send)
					delete(ws.Clients, client.ClientId)
				}
			}
		}
	}
}

func (ws *Websocket) ServeWs(res http.ResponseWriter, req *http.Request) {
	var query string
	if req.URL.Query().Get("key") == "" {
		http.Error(res, "Missing key query", http.StatusBadRequest)
		return
	}
	query = req.URL.Query().Get("key")

	conn, err := ConnectionUpgrader.Upgrade(res, req, nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// initialize websocket client
	client := &Client{
		Ws:       ws,
		Conn:     conn,
		ClientId: strconv.Itoa(id),
		Send:     make(chan []byte, 256),
	}

	// register initialized client
	client.Ws.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// another goroutines.
	go client.WritePump()
	go client.ReadPump()
}
