package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	lock sync.Mutex
)

// 互斥锁

func test(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i

	}
	// 读写操作之前加锁
	lock.Lock()
	myMap[n] = res
	//写完之后解锁
	lock.Unlock()

}
func Test() {
	//for i := 0; i <= 1; i++ {
	//	fmt.Println("test() helloword" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//
	//}
	fmt.Println("test")
}
func test1() {
	go Test()
	//for i := 0; i <= 1; i++ {
	//
	//	fmt.Println("main() helloword" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//
	//}
	fmt.Println("test1")
	time.Sleep(time.Second)
}
func main() {
	//test1()
	//time.Sleep(time.Second * 3)
	//test1()
	for i := 0; i < 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	//fmt.Println("myMap", myMap)
	//lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%v]:%v \n", i, v)

	}
	//lock.Unlock()
}
