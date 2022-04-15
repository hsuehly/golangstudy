package main

import (
	"fmt"
	"io"
	"main/chatroom/common/message"
	process2 "main/chatroom/server/process"
	"main/chatroom/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送的消息种类不同，决定调用那个函数来处理
func (pro *Processor) serverProcessMes(mes *message.Message) (err error) {
	//是否能监听到消息
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建实列
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册流程
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 创建一个smsprocess 实力转发消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
func (pro *Processor) process2() (err error) {
	// 循环读取客户端发送过来的消息

	for {
		// 直接封装一个函数返回message err
		tf := &utils.Transfer{Conn: pro.Conn}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，f服务端也退了")
				return err
			} else {
				fmt.Println("读取数据失败,err=", err)
				return err

			}
		}
		//fmt.Println("msg,", msg)
		err = pro.serverProcessMes(&msg)
		if err != nil {
			return err
		}
	}
}
