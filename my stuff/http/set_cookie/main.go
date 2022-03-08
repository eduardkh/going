package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Println("server running")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "TestCookie",
		Value:  "test",
		MaxAge: 300,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(200)
	w.Write([]byte("you've got a cookie"))
}
