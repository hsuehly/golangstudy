package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"main/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)

	fmt.Println("读取客户端发送的数据。。。")
	// conn.Read 在conn没有别关闭的情况下，才会阻塞
	// 如果客户端关闭了conn 则不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("读取客户端信息失败")
		return
	}
	// 根据buf[:4] 转成一个uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	// 根据pkglen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("conn.Read fail err")
		return
	}
	// 把paklen 反序列化成message.Message
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		//err = errors.New("反序列化失败")
		return
	}
	return
}

// 编写serverprocessmes 函数
// 根据客户端发送的消息种类不同，决定调用那个函数来处理
func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	// 循环读取客户端发送过来的消息

	for {
		// 直接封装一个函数返回message err
		msg, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，f服务端也退了")
				return
			} else {
				fmt.Println("读取数据失败,err=", err)
				return

			}
		}
		fmt.Println("msg,", msg)
	}

}

func main() {

	fmt.Println("服务a在8080 端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器监听失败，err=", err)
		return
	}
	// 监听成功需要循环 等待客户端来连接
	for {
		fmt.Println("等待客户端来连接服务")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 起一个协程
		go process(conn)
	}

}
