package routes

import (
	"gorillamux/handlers"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/methods", handlers.GetHandler).Methods("GET")
	router.HandleFunc("/methods", handlers.PostHandler).Methods("POST")
	router.HandleFunc("/methods", handlers.PutHandler).Methods("PUT")

	router.HandleFunc("/name/{name}/age/{age}", handlers.PathParamHandler).Methods("GET")
}
