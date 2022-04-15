package process

import (
	"fmt"
	"main/chatroom/client/model"
	"main/chatroom/common/message"
)

// 客户端要维护的map

var onlineUsers = make(map[int]*message.User, 10)
var CurUser model.CurUser // 用户登录成功后完成对curuser的初始化

//在客户端显示前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		//如果不是 显示自己
		fmt.Println("用户id：\t", id)

	}

}

// 编写一个方法，处理返回的notifyUserStatusmes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()

}
