package monitor

import(
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

func GetCPUUsage()(float64,error) {
	percent,err:=cpu.Percent(time.Second,false) // false specifies all cores not just one core
	if err != nil{
		return 0 , err
	}
	return percent[0],nil
}