package main

import (
	"log"
	"net/http"

	"gorillamux/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/methods", handlers.GetHandler).Methods("GET")
	router.HandleFunc("/methods", handlers.PostHandler).Methods("POST")
	router.HandleFunc("/methods", handlers.PutHandler).Methods("PUT")

	log.Println("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
