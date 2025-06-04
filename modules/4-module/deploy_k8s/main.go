package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow World"))
	})

	log.Println("Server listening at port 8080")
	http.ListenAndServe(":8080", nil)
}
