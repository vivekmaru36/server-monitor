package api

import (

	// imports for proper shutdown
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"encoding/json"
	// "net/http"
	"server-monitor/monitor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// prometheus variables
var (
	cpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage_percent",
		Help: "Current CPU usage percent",
	})
	memGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage_percent",
		Help: "Current memory usage percent",
	})
	diskGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "disk_usage_percent",
		Help: "Current disk usage percent",
	})
)

type Stats struct{
	Uptime string `json:"uptime"`
	CPUUsage float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage float64 `json:"disk_usage"`
} 

func statsHandler(w http.ResponseWriter,r *http.Request)  {
	cpu, _ := monitor.GetCPUUsage()
	mem, _ := monitor.GetMemoryUsage()
	disk, _ := monitor.GetDiskUsage()
	uptime, _ := monitor.GetUptime()

	// prmoetheus variables settings
	cpuGauge.Set(cpu)
	memGauge.Set(mem)
	diskGauge.Set(disk)

	
	stats := Stats{
		Uptime: uptime,
		CPUUsage: cpu,
		MemoryUsage:mem,
		DiskUsage:disk,
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(stats)
}

func uptimeHandler(w http.ResponseWriter, r *http.Request)  {
	uptime,err:=monitor.GetUptime()
	if err != nil {
		http.Error(w,"Failed to get uptime",http.StatusInternalServerError)
		return
	}
	w.Write([]byte(uptime))
}

// imp stuff
// func StartServer()  {

// 	// auto updating calling
// 	startMetricsUpdater()

// 	http.HandleFunc("/stats",statsHandler)
// 	http.HandleFunc("/uptime",uptimeHandler)
// 	http.HandleFunc("/health",healthHandler)

// 	http.Handle("/metrics",promhttp.Handler())

// 	// http.ListenAndServe(":8080",nil)

// 	log.Println("Server started at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(http.DefaultServeMux)))

// }


// health handler
func healthHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// func for auto updating /metrics tab
func startMetricsUpdater()  {
	go func() {
		for{
			cpu,err1:=monitor.GetCPUUsage()
			mem,err2:=monitor.GetMemoryUsage()
			disk,err3:=monitor.GetDiskUsage()

			if err1==nil {
				cpuGauge.Set(cpu)
			}
			if err2==nil {
				memGauge.Set(mem)
			}
			if err3==nil {
				diskGauge.Set(disk)
			}

			

			log.Printf("Updated metrics - CPU: %.2f%%, MEM: %.2f%%, DISK: %.2f%%", cpu, mem, disk)


			time.Sleep(3 * time.Second)  // every 5 seconds
		}
	}()
}

// proper shuttdown addition implemeted in new func
func StartServer() {
	startMetricsUpdater()

	mux := http.NewServeMux()
	mux.HandleFunc("/stats", statsHandler)
	mux.HandleFunc("/uptime", uptimeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/metrics", promhttp.Handler())

	// Wrap with logging middleware
	loggedMux := loggingMiddleware(mux)

	// Create server instance
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	// Graceful shutdown setup
	go func() {
		log.Println("Server started at http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait here

	log.Println("⚠️ Shutdown signal received. Cleaning up...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server shutdown complete.")
}


// this store the log of all request that are hitted to us
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s %s", time.Now().Format(time.RFC3339), r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
