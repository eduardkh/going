package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func main() {
	// routes
	http.HandleFunc("/", index)
	http.HandleFunc("/getform", getform)
	http.HandleFunc("/postform", postform)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

// handlers
func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.gohtml", nil)
}
func getform(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	context := map[string]string{
		"Name":  name,
		"Email": email,
	}
	templates.ExecuteTemplate(w, "getform.gohtml", context)
}
func postform(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	context := map[string]string{
		"Name":  name,
		"Email": email,
	}
	templates.ExecuteTemplate(w, "postform.gohtml", context)
}
