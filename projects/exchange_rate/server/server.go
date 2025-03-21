package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting at :8080")
	http.HandleFunc("/exchange-rate", exchangeRateHandler)
	http.ListenAndServe(":8080", nil)
}
