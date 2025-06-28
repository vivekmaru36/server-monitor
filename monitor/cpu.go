package monitor

import(
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	// "time"
)

/*
func GetCPUUsage()(float64,error) {
	percent,err:=cpu.Percent(time.Second,false) // false specifies all cores not just one core
	if err != nil{
		return 0 , err
	}
	return percent[0],nil
}
*/

func GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) == 0 {
		return 0, nil
	}
	return percentages[0], nil
}

func GetCPUCoreCount() (int, error) {
	count, err := cpu.Counts(true)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetLoadAverage() (*load.AvgStat, error) {
	return load.Avg()
}