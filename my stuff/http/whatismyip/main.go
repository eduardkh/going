package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.Split(r.RemoteAddr, ":")[0])
}

func main() {
	http.HandleFunc("/", index)
	log.Println("whatismyip Listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
