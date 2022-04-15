package main

import (
	"fmt"
	"main/chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	//var loop = true
	for {
		fmt.Println("---------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			//loop = false
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)

			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)

			fmt.Println("请输入用户名字(昵称)")
			fmt.Scanf("%s\n", &userName)
			//loop = false
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			//loop = false

		default:
			fmt.Println("你的输入有误，请重新输入")

		}
	}
	//if key == 1 {
	//	//fmt.Println("请输入用户id")
	//	//fmt.Scanf("%d\n", &userId)
	//	//fmt.Println("请输入用户密码")
	//	//fmt.Scanln(&userPwd)
	//	login(userId, userPwd)
	//
	//} else if key == 2 {
	//	fmt.Println("进行用户注册的逻辑")
	//}

}
