// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the gonotification project.

package app

import (
	"database/sql"
	"flag"
	"github.com/Nicolas-ggd/go-notification/pkg/microhandler"
	"github.com/Nicolas-ggd/go-notification/pkg/queue"
	"github.com/Nicolas-ggd/go-notification/pkg/repository"
	"github.com/Nicolas-ggd/go-notification/pkg/server"
	"github.com/Nicolas-ggd/go-notification/pkg/server/ws"
	"github.com/Nicolas-ggd/go-notification/pkg/services"
	"github.com/Nicolas-ggd/go-notification/pkg/storage"
	"log"
	"os"
)

func Run() {
	httpPort := flag.String("http-server-port", "8741", "a string")
	natsUrl := flag.String("nats-url", "nats://nats:4222", "a string")

	flag.Parse()

	err := os.Setenv("NATS_URL", *natsUrl)
	if err != nil {
		return
	}

	db, err := storage.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	nc, err := storage.NewNatsConn(*natsUrl)
	if err != nil {
		log.Fatal(err)
	}

	wss := ws.NewWebsocket()

	go wss.Run()

	priorityQueue := queue.NewPriorityQueue()

	repositories := repository.NewRepository(db)
	service := services.NewService(repositories)

	notificationHandler := microhandler.NewMicroHandler(service, priorityQueue, wss)

	microhandler.MicroServices(nc, notificationHandler)

	server.NewServer(wss, httpPort)

}
