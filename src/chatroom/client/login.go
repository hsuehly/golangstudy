package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"main/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	fmt.Printf("userId%v, userPwd%v \n", userId, userPwd)
	// 创建连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接客户端失败，err=", err)
		return errors.New("连接客户端失败")
	}
	defer conn.Close()
	// 创建消息结构体
	mes := message.Message{Type: message.LoginMesType}
	loginMes := message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	// 将loginmes 序列化
	loginData, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes 序列化失败", err)
		return errors.New("loginMes 序列化失败")
	}
	fmt.Println("loginData", string(loginData))
	mes.Data = string(loginData)
	// 再将mes 进行序列化
	mesData, err := json.Marshal(mes)
	if err != nil {
		return errors.New("mes序列化失败")
	}
	fmt.Println("mesData", string(mesData))
	// 先把字节的长度发给客户端
	// 先获取data的长度 -> 转成一个表示长度的切片
	var pkgLen uint32
	// 获取发送信息的长度
	pkgLen = uint32(len(mesData))
	fmt.Println("包的长度", pkgLen)
	// 创建一个byte切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	fmt.Println("buf", buf)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		return errors.New("向服务端发送数据错误")
	}
	_, err = conn.Write(mesData)
	if err != nil {
		return errors.New("向服务端发送数据错误")
	}
	//println("参数", string(mesData))
	//fmt.Printf("向客户端发送数据长度=%d,内容是=%s", len(mesData), string(mesData))
	return
}
