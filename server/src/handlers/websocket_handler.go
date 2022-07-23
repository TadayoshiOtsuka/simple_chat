package handlers

import (
	"log"
	"net/http"

	"github.com/TadayoshiOtsuka/simple_chat/src/domain"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	hub *domain.Hub
}

func NewWebsocketHandler(hub *domain.Hub) *WebsocketHandler {
	return &WebsocketHandler{hub: hub}
}

func (h *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	c := domain.NewClient(ws)
	go c.Read(h.hub.BroadcastCh, h.hub.UnRegisterCh)
	go c.Write()
	h.hub.RegisterCh <- c
}
