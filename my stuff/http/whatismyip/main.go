package main

import (
	"fmt"
	"net/http"
	"strings"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.Split(r.RemoteAddr, ":")[0])
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
