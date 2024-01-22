/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package websocket

type Client interface {
	GetUsername() string
	ReadPump(pool ClientPool)
	WritePump(pool ClientPool)
	SendMessage(msg Message)
	Leave()
}

type ClientPool interface {
	Add(c Client)
	Clients() map[Client]bool
	Broadcast(message Message)
	Eject(c Client)
	Notify(message Message)
	Run()
}
