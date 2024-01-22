/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package websocket

import (
	"context"
	"strings"
	"time"

	"github.com/samsul96maarif/auth-service/lib/logger"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type client struct {
	*websocket.Conn
	username string
	send     chan Message
}

func NewClient(conn *websocket.Conn, u string) *client {
	return &client{conn, u, make(chan Message, 256)}
}

type Message struct {
	From     Client `json:"from"`
	To       Client
	Content  interface{} `json:"Message"`
	Username string      `json:"From"`
	Language string      `json:"language"`
	Type     string      `json:"type"`
}

func (c *client) GetUsername() string { return c.username }

func (c *client) Leave() {
	close(c.send)
}

func (c *client) ReadPump(pool ClientPool) {
	defer func() {
		pool.Eject(c)
		if c.Conn != nil {
			c.Conn.Close()
		}
	}()
	c.Conn.SetPongHandler(func(string string) error {
		return nil
	})
	for {
		var payload Message
		if err := c.Conn.ReadJSON(&payload); err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				break
			}
			logger.Error(context.Background(), c.username+" error readJSON : "+err.Error(), map[string]interface{}{
				"tags":  []string{"websocket", "readJSON"},
				"error": err,
			})
			continue
		}
		if payload.To == nil {
			pool.Broadcast(payload)
		}
	}
}

func (c *client) SendMessage(msg Message) {
	c.send <- msg
}

func (c *client) WritePump(pool ClientPool) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if c.Conn != nil {
			c.Conn.Close()
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.Conn.WriteJSON(message)
			if err != nil {
				logger.Error(context.Background(), err.Error()+" to: "+c.GetUsername(), map[string]interface{}{
					"tags": []string{"websocket", "WritePump"},
				})
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
