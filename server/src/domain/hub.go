package domain

import "log"

type Hub struct {
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnRegisterCh chan *Client
	BroadcastCh  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.register(client)
		case client := <-h.UnRegisterCh:
			h.unregister(client)
		case msg := <-h.BroadcastCh:
			h.broadCast(msg)
		}
	}
}

func (h *Hub) register(c *Client) {
	log.Println("new client joined")
	h.Clients[c] = true
}

func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}

func (h *Hub) broadCast(msg []byte) {
	for c := range h.Clients {
		c.sendCh <- msg
	}
}