package local

import (
	"os/exec"
	"fmt"
)

func Stophost(ip string,hostname string){
	command := "./script/hostStop.sh "+ip+" "+hostname
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("list host:%s failed with error:%s", command, err.Error())
        return
    }
}