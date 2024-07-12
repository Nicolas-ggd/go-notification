// SPDX-License-Identifier: MIT
// Copyright (c) 2024 TOMIOKA
//
// This file is part of the gonotification project.

package server

import (
	"github.com/Nicolas-ggd/go-notification/pkg/server/ws"
	"log"
	"net/http"
)

// todo: maybe it's better to use http2 package for better security
func NewServer(wss *ws.Websocket, httpPort *string) {
	http.HandleFunc("GET /ws", wss.ServeWs)

	err := http.ListenAndServe(":"+*httpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
