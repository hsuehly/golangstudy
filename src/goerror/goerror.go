package goerror

import (
	"errors"
	"fmt"
)

func Error() {
	err := readConf("config.ii")
	if err != nil {
		panic(err)
	}
	fmt.Println("kkkk")
	defer func() {
		err := recover() // recover内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res, "res")

}

// 自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}

}
