package ws

import (
	"encoding/json"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models"
	"log"
)

// SendEvent function send events to the client
func (ws *Websocket) SendEvent(clients []string, model *models.Notification) {
	v, err := json.Marshal(&model)
	if err != nil {
		log.Printf("Can't marshal action data")
		return
	}

	for _, client := range clients {
		c, ok := ws.Clients[client]
		if !ok {
			log.Printf("Client with ID %s not found", client)
			continue // Continue process
		}

		// Check if the Send channel is initialized
		if c.Send == nil {
			log.Printf("Send channel not initialized for client with ID %v", client)
			continue // Continue process
		}

		// Send data to the client
		c.Send <- v
	}

}

// BroadcastEvent function send events in broadcast
func (ws *Websocket) BroadcastEvent(model *models.Notification) {
	value, err := json.Marshal(&model)
	if err != nil {
		log.Printf("Can't marshal broadcast data")
		return
	}

	// Send data to the Broadcast channel
	ws.Broadcast <- value
}
