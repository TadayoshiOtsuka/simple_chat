package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TadayoshiOtsuka/simple_chat/src/domain"
	"github.com/TadayoshiOtsuka/simple_chat/src/handlers"
	"github.com/TadayoshiOtsuka/simple_chat/src/services"
)

func main() {
	pubsub := services.NewPubSubService()
	hub := domain.NewHub(pubsub)
	go hub.SubscribeMessages()
	go hub.RunLoop()

	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}
