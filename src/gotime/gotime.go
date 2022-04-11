package gotime

import (
	"fmt"
	"strconv"
	"time"
)

func Times() {
	// 获取当前日期
	time := time.Now()
	fmt.Printf("当前时间：%v, type=%T \n", time, time)
	fmt.Printf("当前时间年：%v \n", time.Year())
	fmt.Printf("当前时间月：%v \n", int(time.Month()))
	fmt.Printf("当前时间日：%v \n", time.Day())
	fmt.Printf("当前时间时：%v \n", time.Hour())
	fmt.Printf("当前时间分：%v \n", time.Minute())
	fmt.Printf("当前时间秒：%v \n", time.Second())
	// 格式化时间
	fmt.Println("当前时间", time.Format("2006-01-02 15:04:05"))
	fmt.Println("当前时间", time.Format("2006/01/02"))
	fmt.Println("当前时间", time.Format("15:04:05"))
	var i = 0
	for {
		i++
		fmt.Println("i", i)
		//time.Sleep(time.UnixMilli())
		if i == 10 {
			break
		}
	}
}
func TestTime() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
