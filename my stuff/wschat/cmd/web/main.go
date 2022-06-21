package main

import (
	"log"
	"net/http"
	"wschat/handlers"
)

func main() {
	mux := routes()
	log.Println("starting channel listener")
	go handlers.ListenToWsChannel()
	log.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
