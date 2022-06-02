package main

import (
	"log"
	"net/http"
	"wschat/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	log.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
