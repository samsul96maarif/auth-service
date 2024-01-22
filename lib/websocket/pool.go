/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package websocket

type pool struct {
	ClientConnections map[Client]bool
	sentMessage       chan Message
	register          chan Client
	unregister        chan Client
}

func NewPool() *pool {
	return &pool{make(map[Client]bool), make(chan Message), make(chan Client), make(chan Client)}
}

func (p *pool) Add(c Client) {
	p.register <- c
}

func (p *pool) Broadcast(payload Message) {
	p.sentMessage <- payload
}

func (p *pool) Clients() map[Client]bool {
	return p.ClientConnections
}

func (p *pool) Eject(c Client) {
	p.unregister <- c
}

func (p *pool) Notify(msg Message) {
	for c, _ := range p.ClientConnections {
		if c.GetUsername() == msg.Username {
			c.SendMessage(msg)
		}
	}
}

func (p *pool) Run() {
	for {
		select {
		case c := <-p.register:
			p.ClientConnections[c] = true
		case c := <-p.unregister:
			if _, ok := p.ClientConnections[c]; ok {
				delete(p.ClientConnections, c)
				c.Leave()
			}
		case msg := <-p.sentMessage:
			if msg.Type == "notification" {
				p.Notify(msg)
			}
		}
	}
}
