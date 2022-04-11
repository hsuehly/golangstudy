package service

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 关闭conn
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		// 等待客户端通过conn发送消息
		// 如果客户端没有write[发送] 那么协程就会阻塞在这里
		fmt.Println("服务器等待客户端发送消息" + conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn 读取
		if err == io.EOF {
			fmt.Println("客户端退出", err)
			return // !!!! 一旦出问题就直接退出
		}
		fmt.Print(string(buf[:n]))
	}
}
func Server() {
	fmt.Println("服务器开始监听。。。")
	// tcp表示协议
	// address 监听的地址和端口
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("服务器启动失败", err)
		return
	}
	fmt.Println("listen", listen)
	defer listen.Close() // 延时关闭
	for {
		fmt.Println("等待客户端链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accper err", err)
		} else {
			fmt.Printf("conn 链接成功 conn=%v,客户端ip=%v \n", conn, conn.RemoteAddr().String())
		}
		// 这里准备一个协程，为客户端服务
		go process(conn)

	}

}
