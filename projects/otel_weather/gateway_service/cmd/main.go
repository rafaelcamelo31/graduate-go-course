package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/gateway_service/internal/handler"
)

func main() {
	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")

	http.HandleFunc("POST /api/temperature", handler.TemperatureHandler)

	slog.Info("starting server at:", "port", port)
	http.ListenAndServe(":"+port, nil)
}
