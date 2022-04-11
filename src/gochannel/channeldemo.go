package gochannel

import (
	"fmt"
	"time"
)

func writeData(numChan chan int) {
	for i := 1; i < 50; i++ {
		numChan <- i
		fmt.Println("i", i)
		time.Sleep(time.Second)

	}
	close(numChan)
}
func readData(numChan chan int, exitChan chan bool) {
	for {
		v, ok := <-numChan
		if !ok {
			fmt.Printf("ok= %v", ok)
			break
		}
		// 每隔一秒读一次
		//time.Sleep(time.Second)
		fmt.Printf("readData 读的数据=%v boolea=%v\n", v, ok)
	}
	exitChan <- true
	close(exitChan)
}
func ChannelDemo() {
	// 如果编译器（运行），发现一个管道只有写，而没有读，则该管道会阻塞
	//写管道和读管道的频率不一致，无所谓
	numChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	// 只写不读，会阻塞，发生死锁
	go readData(numChan, exitChan)
	// 读和写频率不一样没关系，但是必须要有读
	go writeData(numChan)
	//time.Sleep(time.Second * 5)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}

	}
}
