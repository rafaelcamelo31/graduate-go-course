package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rafaelcamelo31/graduate-go-course/projects/multithreading/internal/handler"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handler.GetAddress)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
