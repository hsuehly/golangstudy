package gochannel

import "fmt"

func Details() {
	// 定义一个管道10 个int
	intChan := make(chan int, 10)
	strChan := make(chan string, 5)

	for i := 0; i < 10; i++ {
		intChan <- i

	}
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭管道会阻塞而导致deadlock
	// 问题 在实际开发中，我们可能不好确定什么关闭管道
	// 可以使用select方式解决
	//labe2:
	for {
		// 注意： 这里如果intchan一直没有关闭不会一直阻塞管道而deadlock
		// 会自动到下一个case匹配
		select {
		case v := <-intChan:
			fmt.Printf("从intchan取到的值%d \n", v)
		case v := <-strChan:
			fmt.Printf("从strchan取到的值%v \n", v)
		default:
			fmt.Println("都取不到值，程序员可以写一些逻辑")
			//break labe2
			return
		}

	}
}
