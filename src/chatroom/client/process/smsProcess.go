package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/client/utils"
	"main/chatroom/common/message"
)

type SmsProcess struct {
}

// 发送群聊的消息

func (sm *SmsProcess) SendGroupMes(content string) (err error) {
	//  创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	// 创建一个smsmes 实力
	var smsMes message.SmsMes
	smsMes.Conntent = content //内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	// 将mes 发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	return
}
