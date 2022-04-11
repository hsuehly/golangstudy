package gochannel

import (
	"fmt"
	"time"
)

func putnum(intChan chan int) {
	// 管道引用类型
	for i := 1; i <= 100000; i++ {
		intChan <- i
	}
	// 放完关闭管道
	close(intChan)

}

// 从intchan取出数据，并判断是否为素数，如果是放入primechan中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用for循环
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true // 假设是素数
		// 判读num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}

		}
		if flag {
			//如果是素数，向primer管道添加
			primeChan <- num
		}

	}
	fmt.Println("有一个primeNum协程取不到数据退出")
	// 这里不能直接关闭prime管道， 应为比人还需要用
	// 向exitChan 写入true
	exitChan <- true
}
func GetPrime() {
	intChan := make(chan int, 1000)
	// 放入结果管道
	primeChan := make(chan int, 4000)
	// 标识退出管道
	exitChan := make(chan bool, 4)
	start := time.Now().Unix()
	// 开启一个协程向intchan放入1-8000个数
	go putnum(intChan)
	// 开启四个协程
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)

	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时：", end-start)
		// 当我们从exitchan取出4个结果，就可以关闭primeChan管道
		close(primeChan)
	}()
	// 遍历primechan管道取出结果

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数%d \n", res)
	}
	fmt.Println("main 主线程退出")
}
