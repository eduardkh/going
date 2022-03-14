package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func main() {
	// initiate FS
	fs := http.FileServer(http.Dir("."))
	// routes
	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

// handlers
func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.gohtml", nil)
}
