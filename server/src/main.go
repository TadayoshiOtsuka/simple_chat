package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TadayoshiOtsuka/simple_chat/src/domain"
	"github.com/TadayoshiOtsuka/simple_chat/src/handlers"
)

func main() {
	hub := domain.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}
