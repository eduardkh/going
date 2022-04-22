package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET Method\n"))
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST Method\n"))
}
func PutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PUT Method\n"))
}
func PathParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	age := vars["age"]
	fmt.Fprintf(w, "Name : %s\nAge  : %s\n", name, age)
}
