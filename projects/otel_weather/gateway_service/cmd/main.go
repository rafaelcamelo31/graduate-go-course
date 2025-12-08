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

	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/gateway_service/internal/adapter"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/gateway_service/internal/handler"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/observability"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var isShuttingDown atomic.Bool

func main() {
	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")

	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	shutdown, err := observability.InitOpenTelemetry(rootCtx, "gateway-service", "otel-collector:4317")
	if err != nil {
		slog.Error("failed to initialize open telemetry", "error", err)
	}
	defer func() {
		slog.Info("shutting down tracer provider...")
		shutdown(rootCtx)
	}()

	client := &http.Client{
		Timeout:   3 * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	ta := adapter.NewHttpTemperatureServiceAdapter(client)
	h := handler.NewHandler(ta)

	ongoingCtx, stopOngoingGracefully := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:         ":" + port,
		BaseContext:  func(_ net.Listener) context.Context { return ongoingCtx },
		ReadTimeout:  time.Second,
		WriteTimeout: 5 * time.Second,
	}

	http.Handle("POST /api/temperature",
		otelhttp.NewHandler(
			http.HandlerFunc(h.TemperatureHandler),
			"POST /api/temperature",
		),
	)

	go func() {
		slog.Info("starting server", "port", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "error", err)
			os.Exit(1)
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
