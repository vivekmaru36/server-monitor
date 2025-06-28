package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer() {
	startMetricsUpdater() // from metrics.go

	mux := http.NewServeMux()
	mux.HandleFunc("/stats", statsHandler)
	mux.HandleFunc("/uptime", uptimeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./static"))))

	// Wrap with logging middleware
	loggedMux := loggingMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	go func() {
		log.Println("Server started at http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("⚠️ Shutdown signal received. Cleaning up...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("✅ Server shutdown complete.")
}
