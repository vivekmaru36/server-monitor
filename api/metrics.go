package api

import (
	"log"
	"time"
	"server-monitor/monitor"
	"github.com/prometheus/client_golang/prometheus"
)

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

func startMetricsUpdater() {
	go func() {
		for {
			cpu, err1 := monitor.GetCPUUsage()
			mem, err2 := monitor.GetMemoryUsage()
			disk, err3 := monitor.GetDiskUsage()

			if err1 == nil {
				cpuGauge.Set(cpu)
			}
			if err2 == nil {
				memGauge.Set(mem)
			}
			if err3 == nil {
				diskGauge.Set(disk)
			}

			log.Printf("Updated metrics - CPU: %.2f%%, MEM: %.2f%%, DISK: %.2f%%", cpu, mem, disk)
			time.Sleep(3 * time.Second)
		}
	}()
}
