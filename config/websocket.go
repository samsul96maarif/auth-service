/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package config

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	WsUpgrader = websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	WebsocketTopic = "WEBSOCKET-TOPIC"
)
