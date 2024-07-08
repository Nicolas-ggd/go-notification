package main

import (
	"flag"
	"github.com/Nicolas-ggd/go-notification/pkg/config"
	"github.com/Nicolas-ggd/go-notification/pkg/handlers"
	"github.com/Nicolas-ggd/go-notification/pkg/http/ws"
	"github.com/Nicolas-ggd/go-notification/pkg/storage"
	"github.com/nats-io/nats.go/micro"
	"log"
	"net/http"
)

func main() {
	httpPort := flag.String("http-server-port", "8080", "a string")
	natsUrl := flag.String("nats-url", "nats://127.0.0.1:4222", "a string")

	flag.Parse()

	nc, err := storage.NewNatsConn(*natsUrl)
	if err != nil {
		log.Fatal(err)
	}

	wss := ws.NewWebsocket()

	_, err = micro.AddService(nc, micro.Config{
		Name: config.StreamName,
		Endpoint: &micro.EndpointConfig{
			Subject:    config.SubjectBroadcastNotification,
			Handler:    handlers.BroadcastNotification(wss),
			Metadata:   nil,
			QueueGroup: "",
		},
		Version:     config.SubjectVersion,
		Description: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = micro.AddService(nc, micro.Config{
		Name: config.StreamName,
		Endpoint: &micro.EndpointConfig{
			Subject:    config.SubjectClientNotification,
			Handler:    handlers.ClientBasedNotification(wss),
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

	err = http.ListenAndServe(":"+*httpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
