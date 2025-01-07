package local

import (
	"os/exec"
	"fmt"
)

func Removehost(ip string,hostname string){
	command := "./script/hostRemove.sh "+ip+" "+hostname
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("list host:%s failed with error:%s", command, err.Error())
        return
    }
}