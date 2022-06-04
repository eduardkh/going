package main

import (
	"net/http"
	"wschat/handlers"

	"github.com/bmizerany/pat"
)

// https://mpolinowski.github.io/devnotes/2021-09-10--websocket-recconects
var fs = http.FileServer(http.Dir("./static/"))

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	mux.Get("/static/", http.StripPrefix("/static", fs))
	return mux
}
