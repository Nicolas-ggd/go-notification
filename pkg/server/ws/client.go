package ws

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	Ws       *Websocket
	Conn     *websocket.Conn
	ClientId string
	Send     chan []byte
}

// ReadPump pumps messages from the websocket connection to the websocket.
//
// The application runs ReadPump in a per-connection goroutine. The application
// ensures that there is a most one reader on a connection by executing all
// readers from the goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Ws.UnRegister <- c
	}()

	// set rate limit which use maximum message size to read message
	c.Conn.SetReadLimit(maxMessageSize)

	// set readDead line time by using time.Now using additional pongWait
	// allowed to read the next pong message from peer.
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return
	}

	// set pong handler which use readDeadline setter
	c.Conn.SetPongHandler(func(string) error {
		err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}

		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		fmt.Printf("Received message: %s\n", message)
	}
}

// WritePump pumps message from the hub to the websocket connection.
//
// A goroutine running WritePUmp is started for each connection. The
// application ensures that there is a most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	// Send piings to peer with this period.
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.Conn.Close(); err != nil {
			return
		}
	}()

	for {
		select {
		// receive message from the channel
		case message, ok := <-c.Send:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return
			}
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					return
				}
				return
			}

			// write message for websocket
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			// write message
			_, err = w.Write(message)
			if err != nil {
				return
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				_, err = w.Write(newLine)
				if err != nil {
					return
				}
			}
			// close websocket
			if err = w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return
			}
			err := c.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}
