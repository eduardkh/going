package main

import (
	"log"
	"net/http"

	"gorillamux/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.Routes(router)
	log.Println("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
