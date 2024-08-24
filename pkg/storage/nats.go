package storage

import "github.com/nats-io/nats.go"

// NewNatsConn function open new nats client connection
func NewNatsConn(natsUrl string) (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
