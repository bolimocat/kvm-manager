package local

import (
	"os/exec"
	"fmt"
//	"golang.org/x/crypto/ssh"
//    "io"
//    "log"
)

func Listhost(ip string,status string){
	command := "./script/hostList.sh "+ip+" "+status
	cmd := exec.Command("/bin/bash", "-c", command)
    out, err := cmd.Output()
    if err != nil {
        fmt.Printf("list host:%s failed with error:%s", command, err.Error())

    }
    fmt.Println(string(out))

}

//func SSHExecute(user string, password string, host string, command string) error {
//	fmt.Println("new list method")
//    // 配置SSH连接的相关信息
//    config := &ssh.ClientConfig{
//        User: user,
//        Auth: []ssh.AuthMethod{
//            ssh.Password(password),
//        },
//        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//        
//    }
//    fmt.Println("配置SSH连接的相关信息 完成")
//    // 连接SSH服务器
//    client, err := ssh.Dial("tcp", host+":22", config)
//    if err!= nil {
//        return err
//    }
//    defer client.Close()
//	fmt.Println("--------------1-------------")
//    // 创建一个会话
//    session, err := client.NewSession()
//    if err!= nil {
//        return err
//    }
//    defer session.Close()
//    fmt.Println("--------------2-------------")
//
//    // 创建一个管道用于读取命令执行结果
//    var result []byte
//    r, w := io.Pipe()
//    session.Stdout = w
//    session.Stderr = w
//	fmt.Println("--------------a-------------"+command)
//    // 开始执行命令
//	err = session.Run(command)
//	fmt.Println("--------------a1-------------")
//    w.Close()
//    fmt.Println("--------------a2-------------")
//    if err!= nil {
//        return err
//        fmt.Println("--------------a3-------------")
//    }
////    fmt.Println(string(info))
//	fmt.Println("--------------3-------------")
//    // 从管道读取命令执行结果到result变量
//    _, err = io.ReadFull(r, result)
//    if err!= nil {
//        return err
//    }
//	fmt.Println("--------------4-------------")
//    log.Println("命令执行结果:", string(result))
//    return nil
//}
