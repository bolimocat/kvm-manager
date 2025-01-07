package local

import (
	"os/exec"
	"fmt"
)

func Starthost(ip string,hostname string){
	command := "./script/hostStart.sh "+ip+" "+hostname
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("list host:%s failed with error:%s", command, err.Error())
        return
    }
}