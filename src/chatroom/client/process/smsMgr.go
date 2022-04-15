package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
)

func ottputGroupMes(mes *message.Message) { // 这个地方mes一定smsmes
	// 显示即可
	// 反序列化mes.data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}
	// 显示信息
	info := fmt.Sprintf("用户ID：\t%d 对大家说\t %s", smsMes.UserId, smsMes.Conntent)
	fmt.Println(info)
	fmt.Println()

}
