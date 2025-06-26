package monitor

import (
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryUsage() (UsedPercent float64,err error) {
	v,err:=mem.VirtualMemory()
	if err!=nil{
		return 0,err
	}
	return v.UsedPercent,nil
}