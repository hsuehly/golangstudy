package main

import (
	"fmt"
	"main/chatroom/server/model"
	"net"
	"time"
)

/*
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

func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给客户端
	var pkgLen uint32
	// 获取发送信息的长度
	pkgLen = uint32(len(data))
	fmt.Println("包的长度", pkgLen)
	// 创建一个byte切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	fmt.Println("buf", buf)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.write(bytes) fail", err)
		return
	}
	// 发送数据
	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.write(bytes) fail", err)
		return
	}
	return

}

//编写一个函数serverPROCESSLOGIN 函数，专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 从mes种取出mes.Data,并反序列化成loginmes
	var resMes = message.Message{
		Type: message.LoginResMesType,
	}
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("jsonUnmarshal mes.Data 失败，err=", err)
		return
	}
	// 先声明一个resmes 返回给客户端
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册后在使用"
	}
	loginResData, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal resmes 失败,err=", err)
		return
	}
	resMes.Data = string(loginResData)
	// 序列化返回的字段
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal resmes 失败，err=", err)
		return
	}
	// 发送data
	err = writePkg(conn, data)
	return
}

// 编写serverprocessmes 函数
// 根据客户端发送的消息种类不同，决定调用那个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	// 处理注册流程
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

*/
func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	// 调用总控
	process := &Processor{Conn: conn}
	err := process.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协议错误err=", err)
		return
	}
}
func init() {
	initPool("39.107.92.94:6379", 16, 0, 300*time.Second)
	initUSerDao()

}
func initUSerDao() {
	// 这里的pool 本事就是一个全局的变量
	// 先initpoll 在 inituserpdao
	model.MyUserDao = model.NewUserDao(pool)

}

func main() {

	fmt.Println("服务b在8080 端口监听")
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
