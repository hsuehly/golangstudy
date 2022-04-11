package main

import (
	"flag"
	"fmt"
)

func main() {
	// os.args 是一个切片
	//fmt.Println("命令行的参数有：", os.Args)
	//fmt.Println("命令行的参数长度：", len(os.Args))
	//for i, v := range os.Args {
	//	fmt.Printf("args[%v]:%v\n", i, v)
	//
	//}
	var user string
	var pwd string
	var host string
	var port int
	//     		 	  传递的值  参数名      参数值       说明
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "p", "", "密码")
	flag.StringVar(&host, "h", "loclhost", "地址")
	flag.IntVar(&port, "P", 3306, "端口号")
	// 非常重要的操作，转换， 必须调用该方法
	//从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	fmt.Println(port, "port")

	fmt.Printf("用户名%v,密码%v, 地址%v, 端口号%v\r\n", user, pwd, host, port)

}
