package api

import (
	"encoding/json"
	"net/http"
	"server-monitor/monitor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
	"log"
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
func StartServer()  {

	// auto updating calling
	startMetricsUpdater()

	http.HandleFunc("/stats",statsHandler)
	http.HandleFunc("/uptime",uptimeHandler)
	http.HandleFunc("/health",healthHandler)

	http.Handle("/metrics",promhttp.Handler())

	http.ListenAndServe(":8080",nil)
}


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