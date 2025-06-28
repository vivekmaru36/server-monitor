package api

import (
	"encoding/json"
	"net/http"
	"server-monitor/monitor"
)

type Stats struct{
	Uptime string `json:"uptime"`
	CPUUsage float64 `json:"cpu_usage`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage float64 `json:"disk_usage"`
} 

func statsHandler(w http.ResponseWriter,r *http.Request)  {
	cpu, _ := monitor.GetCPUUsage()
	mem, _ := monitor.GetMemoryUsage()
	disk, _ := monitor.GetDiskUsage()
	uptime, _ := monitor.GetUptime()
	
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
	http.HandleFunc("/stats",statsHandler)
	http.HandleFunc("/uptime",uptimeHandler)
	http.HandleFunc("/health",healthHandler)

	http.ListenAndServe(":8080",nil)
}


// health handler
func healthHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}