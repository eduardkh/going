package main

import (
	"log"
	"net/http"

	"gorillamux/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/methods", handlers.GetHandler).Methods("GET")
	r.HandleFunc("/methods", handlers.PostHandler).Methods("POST")
	r.HandleFunc("/methods", handlers.PutHandler).Methods("PUT")

	r.HandleFunc("/name/{name}/age/{age}", handlers.PathParamHandler).Methods("GET")

	log.Println("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
