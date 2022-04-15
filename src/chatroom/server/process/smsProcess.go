package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
	"main/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

// 写方法

func (sp *SmsProcess) SendGroupMes(mes *message.Message) {
	// 遍历服务器的map
	// 将消息转发
	// 取出mes 的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return

	}
	for id, up := range userMgr.onlineUsers {
		// 不需要发给自己
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}

}
func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	// 创建一个tf
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败,err=", err)
	}
}
