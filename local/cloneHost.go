package local

import (
	"os/exec"
	"fmt"
)

func Clone(ip string,sourcehost string,newhost string){
	command := "./script/cloneHost.sh "+ip+" "+sourcehost+" "+newhost
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("clone a exist host:%s failed with error:%s", command, err.Error())
        return
    }
}