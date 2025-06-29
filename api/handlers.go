package api

import (
	"encoding/json"
	"net/http"
	"server-monitor/monitor"
)

type Stats struct {
	Uptime      string  `json:"uptime"`
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	CPUCore     int     `json:"cpu_cores"`
	Load1       float64 `json:"load_1"`
	Load5       float64 `json:"load_5"`
	Load15      float64 `json:"load_15"`
	Hostname    string  `json:"hostname"`
	OS          string  `json:"os"`
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	cpu, _ := monitor.GetCPUUsage()
	mem, _ := monitor.GetMemoryUsage()
	disk, _ := monitor.GetDiskUsage()
	uptime, _ := monitor.GetUptime()
	cores, _ := monitor.GetCPUCoreCount()
	load, _ := monitor.GetLoadAverage()
	info, _ := monitor.GetSysInfo()

	cpuGauge.Set(cpu)
	memGauge.Set(mem)
	diskGauge.Set(disk)

	stats := Stats{
		Uptime:      uptime,
		CPUUsage:    cpu,
		MemoryUsage: mem,
		DiskUsage:   disk,
		CPUCore:     cores,
		Load1:       load.Load1,
		Load5:       load.Load5,
		Load15:      load.Load15,
		Hostname:    info.Hostname,
	    OS:          info.Platform,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	uptime, err := monitor.GetUptime()
	if err != nil {
		http.Error(w, "Failed to get uptime", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(uptime))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
