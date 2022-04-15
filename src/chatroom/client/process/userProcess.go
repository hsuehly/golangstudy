package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"main/chatroom/client/utils"
	"main/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

// 注册

func (up *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	// 创建连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接客户端失败，err=", err)
		return errors.New("连接客户端失败")
	}
	defer conn.Close()
	// 创建消息结构体
	mes := message.Message{Type: message.RegisterMesType}
	user := message.User{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}
	registerMes := message.RegisterMes{User: user}
	// 将registerMes 序列化
	regisData, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
	}
	mes.Data = string(regisData)
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
	}
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(mesData)
	if err != nil {
		fmt.Println("注册发送信息错误，err=", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readpkg err=", err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	fmt.Println("registerResMes", registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error, "----")
		os.Exit(0)
	}
	return

}
func (up *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := utils.Transfer{Conn: conn}
	resMes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("readpkg 出错,err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 初始化curuser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		//fmt.Println("登录成功")
		// 显示我们登录成功后的菜单[循环]
		// 遍历logonresmes里的userid
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			// 完成 客户端的onlineUsers 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
