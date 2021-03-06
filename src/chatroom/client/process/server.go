package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/client/utils"
	"main/chatroom/common/message"
	"net"
	"os"
)

// 显示登录成功后的界面

func ShowMenu() {
	var key int
	var content string
	smsProces := &SmsProcess{}
	fmt.Println("-----------登录成功-----------")
	fmt.Println("-----------1. 显示在线用户列表-----------")
	fmt.Println("-----------2. 发送消息-----------")
	fmt.Println("-----------3. 信息列表-----------")
	fmt.Println("-----------4. 退出系统-----------")
	fmt.Println("请选择(1-4):")
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
		fmt.Println("请输入你想要说的话")
		fmt.Scanf("%s\n", &content)
		smsProces.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("请输入正确的选项")

	}
}

// 和服务器保持通讯

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err", err)
			return
		}
		fmt.Println("mes=", mes)
		switch mes.Type {

		case message.NotifyUserStatusMesType: // 有人上线
			// 取出notifyUserstatusmes
			fmt.Println("有人上线")
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 把这个用户信息，状态保存到客户端map中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			ottputGroupMes(&mes)

		// 处理
		default:
			fmt.Println("服务器端返回了未知消息类型")

		}

	}
}
