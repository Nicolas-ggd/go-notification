package main

import (
	"github.com/Nicolas-ggd/go-notification/pkg/config"
	"github.com/Nicolas-ggd/go-notification/pkg/handlers"
	"github.com/Nicolas-ggd/go-notification/pkg/http/ws"
	"github.com/Nicolas-ggd/go-notification/pkg/storage"
	"github.com/nats-io/nats.go/micro"
	"log"
	"net/http"
)

func main() {
	nc, err := storage.NewNatsConn()
	if err != nil {
		log.Fatal(err)
	}

	wss := ws.NewWebsocket()

	_, err = micro.AddService(nc, micro.Config{
		Name: config.StreamName,
		Endpoint: &micro.EndpointConfig{
			Subject:    config.SubjectSendNotification,
			Handler:    handlers.SendNotification(wss),
			Metadata:   nil,
			QueueGroup: "",
		},
		Version:     config.SubjectVersion,
		Description: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /ws", wss.ServeWs)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
