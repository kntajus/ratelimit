package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	s := &http.Server{Addr: ":8080"}
	s.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Printf("Request received from %s\n", id)
}
