package domain

type Hub struct {
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnRegisterCh chan *Client
	BroadcastCh  chan Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan Message),
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.register(client)

		case client := <-h.UnRegisterCh:
			h.unregister(client)

		case msg := <-h.BroadcastCh:
			h.broadCastToAllClient(msg)
		}
	}
}

func (h *Hub) register(c *Client) {
	h.Clients[c] = true
}

func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}

func (h *Hub) broadCastToAllClient(msg Message) {
	for c := range h.Clients {
		c.sendCh <- msg
	}
}
