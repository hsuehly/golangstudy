package gochannel

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("helloword")
	}
}
func test() {
	// 使用defale + recover 捕获错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	var mymap map[int]string
	mymap[0] = "golang"

}
func Error() {
	go sayHello()
	go test()
	//time.Sleep(time.Second * 10)
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)

	}
}
