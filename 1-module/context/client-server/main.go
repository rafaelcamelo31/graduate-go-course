package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request received")
	defer log.Println("Request processed")

	select {
	case <-time.After(5 * time.Second):
		// Prints in comand line stdout
		log.Println("Request processed in 5 seconds")
		// Prints in browser
		w.Write([]byte("Request processed in 5 seconds"))
	case <-ctx.Done():
		// Prints in comand line stdout
		log.Println("Request cancelled by client")
	}
}
