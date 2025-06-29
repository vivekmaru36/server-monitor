package monitor

import (
	"os"
	"github.com/shirou/gopsutil/v3/host"
)

type SysInfo struct {
	Hostname string `json:"hostname"`
	OS 		 string `json:"os"`
	Platform string `json:"platform"`

}

func GetSysInfo() (SysInfo,error)  {
	hoststats,err:=host.Info()
	if err != nil{
		return SysInfo{},err
	}

	name,_ :=os.Hostname()

	return SysInfo{
		Hostname:name,
		OS:hoststats.OS,
		Platform:hoststats.Platform,
	},nil

}