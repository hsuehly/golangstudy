package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Sprintf("client链接失败错误信息%v", err)
		return
	}
	//fmt.Println("conn 成功=", conn)
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入【终端】
	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err=", err)
		}

		if strings.Trim(line, "\r\n") == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//fmt.Println(reader, "reader")
		// 再将line 发送给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.write, err=", err)
		}
		//fmt.Printf("客户端发送了%d字节，并退出", n)

	}
}
