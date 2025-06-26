package monitor

import (
	// "fmt" 
	"os/exec" 
	"strings"
)

func GetUptime() (string,error) {
	out,err:=exec.Command("uptime","-p").Output()
	if err != nil {
		return "",err	
	}
	return strings.TrimSpace(string(out)), nil
}