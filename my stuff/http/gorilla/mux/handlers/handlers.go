package handlers

import "net/http"

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
