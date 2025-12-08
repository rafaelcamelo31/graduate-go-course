package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/observability"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/internal/adapter"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/internal/handler"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var isShuttingDown atomic.Bool

func main() {
	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	shutdown, err := observability.InitOpenTelemetry(rootCtx, "temperature-service", "otel-collector:4317")
	if err != nil {
		slog.Error("failed to initialize open telemetry", "error", err)
	}
	defer shutdown(rootCtx)

	httpClient := &http.Client{
		Timeout:   5 * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	cfg := &config.WeatherConfig{}
	cfg.LoadConfig()

	viacepAdapter := adapter.NewHttpViaCepAdapter(httpClient)
	weatherAdapter := adapter.NewHttpWeatherApiAdapter(httpClient, cfg.ApiKey)

	h := handler.NewHandler(viacepAdapter, weatherAdapter, cfg)
	http.Handle("GET /api/temperature",
		otelhttp.NewHandler(
			http.HandlerFunc(h.GetTemperature),
			"GET /api/temperature",
		),
	)

	port := os.Getenv("PORT")

	ongoingCtx, stopOngoingGracefully := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:         ":" + port,
		BaseContext:  func(_ net.Listener) context.Context { return ongoingCtx },
		ReadTimeout:  time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		slog.Info("starting server", "port", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "error", err)
			panic(err)
		}
	}()

	<-rootCtx.Done()
	stop()
	isShuttingDown.Store(true)
	slog.Info("Received shutdown signal, shutting down.")

	time.Sleep(5 * time.Second)
	slog.Info("Readiness check propagated, now waiting for ongoing requests to finish.")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	stopOngoingGracefully()
	if err != nil {
		slog.Error("Failed to wait for ongoing requests to finish, waiting for forced cancellation.")
		time.Sleep(3 * time.Second)
	}

	slog.Info("Server shut down gracefully.")
}
