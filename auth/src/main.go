package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "81"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}
