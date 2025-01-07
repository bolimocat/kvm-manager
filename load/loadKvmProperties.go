package load

import (
	"os"
    "log"
    "bufio"
    "strings"	
//    "fmt"
)


func Loadproperties(profile string) []string{
//	fmt.Println("读取资源配置文件")
	var sliceinfo []string
	sliceinfo,err := HandleKvmProperties(profile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
	}
}

func HandleKvmProperties(textfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,9)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text() 
	     
	    if !strings.Contains(line, "#"){
	    	
	    	if strings.Contains(line,"x86_host"){
			infoSlice[0] = line
			}
	    	if strings.Contains(line,"arm_host"){
			infoSlice[1] = line
			}
	    	if strings.Contains(line,"disk_path"){
			infoSlice[2] = line
			}
	    	if strings.Contains(line,"iso_path"){
			infoSlice[3] = line
			}
	    	if strings.Contains(line,"templates_path"){
			infoSlice[4] = line
			}
	    	if strings.Contains(line,"x86_templates"){
			infoSlice[5] = line
			}
	    	if strings.Contains(line,"arm_templates"){
			infoSlice[6] = line
			}
	    	if strings.Contains(line,"x86_iso"){
			infoSlice[7] = line
			}
	    	if strings.Contains(line,"arm_iso"){
			infoSlice[8] = line
			}
	    	
	    }
	 }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

	return infoSlice,nil
}
