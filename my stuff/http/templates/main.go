package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func main() {
	// routes
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/suka", suka)
	http.ListenAndServe(":8080", nil)
}

// handlers
func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.gohtml", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "about.gohtml", nil)
}
func suka(w http.ResponseWriter, r *http.Request) {
	context := map[string]string{
		"Name": "Suka",
		"Age":  "20",
	}
	templates.ExecuteTemplate(w, "suka.gohtml", context)
}
