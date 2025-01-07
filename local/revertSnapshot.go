package local

import (
	"os/exec"
	"fmt"
)

func Revertsnapshot(ip string,hostname string,snapshot string){
	command := "./script/snapshotRevert.sh "+ip+" "+hostname+" "+snapshot
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("list host:%s failed with error:%s", command, err.Error())
        return
    }
}