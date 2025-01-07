package local

import (
	"os/exec"
	"fmt"
)

func Listsnapshot(ip string,hostname string){
	command := "./script/snapshotList.sh "+ip+" "+hostname
	cmd := exec.Command("/bin/bash", "-c", command)

    out, err := cmd.Output()
    if err != nil {
        fmt.Printf("list host:%s failed with error:%s", command, err.Error())

    }
    fmt.Println(string(out))

}