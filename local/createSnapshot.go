package local

import (
	"os/exec"
	"fmt"
)

func Createsnapshot(ip string,hostname string){
	command := "./script/snapshotCreate.sh "+ip+" "+hostname
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("clone a exist host:%s failed with error:%s", command, err.Error())
        return
    }
}