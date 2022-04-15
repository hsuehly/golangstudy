package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
	"main/chatroom/server/model"
	"main/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// 通知所有用户在线的方法
// userId 通知其他在线用户，我上线

func (up *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历 onineUsers,然后一个个的发送NotifyUserStatusMes
	for id, um := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		// 开始通知【】写一个方法
		um.NotifyMeOnllien(userId)
	}
}
func (up *UserProcess) NotifyMeOnllien(userId int) {
	var mes = message.Message{Type: message.NotifyUserStatusMesType}
	var NotifyUserStatusMes = message.NotifyUserStatusMes{
		UserId: userId,
		Status: message.UserOnline,
	}

	// 将notifyUserSatusMes 序列化
	data, err := json.Marshal(NotifyUserStatusMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyMeOnline err=", err)
		return
	}
}
func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("jsonUnmarshal mes.Data 失败，err=", err)
		return
	}

	var resMes = message.Message{
		Type: message.RegisterMesType,
	}
	var registerResMes message.RegisterResMes
	// 去数据库完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		fmt.Println("err----", err)
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
		// 这里登录成功
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fail ", err)
	}
	// 发送data,
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//编写一个函数serverPROCESSLOGIN 函数，专门处理登录请求

func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
		up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(up)
		// 通知其他在线用户，我上线了
		up.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户的id放入到loginresmes
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)

		}

	}

	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200
	//} else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在，请注册后在使用"
	//}
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
	tf := &utils.Transfer{Conn: up.Conn}
	err = tf.WritePkg(data)
	return
}
