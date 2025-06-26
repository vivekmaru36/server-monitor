package main

import (
	"fmt"
	"server-monitor/monitor"
)

func main() {
	cpu, _ := monitor.GetCPUUsage()
	mem, _ := monitor.GetMemoryUsage()
	disk, _ := monitor.GetDiskUsage()
	uptime, _ := monitor.GetUptime()

	fmt.Printf("Uptime: %s\n", uptime)
	fmt.Printf("CPU Usage: %.2f%%\n", cpu)
	fmt.Printf("Memory Usage: %.2f%%\n", mem)
	fmt.Printf("Disk Usage: %.2f%%\n", disk)
}
