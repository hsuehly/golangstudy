package gochannel

import "fmt"

type Cat struct {
	name string
	age  int
}

var a int

func testchan(chan3 chan<- int) {
	// 如果函数只操作写入可以声明只写的管道
	chan3 <- 90
	//<-chan3

}
func Channel() {

	// 创建一个管道存放三个
	// 只能存放指定的数据类型
	var intChan chan int
	// 管道使用必须先make
	intChan = make(chan int, 3)
	// 可以创建只读，或者只写的管道 make 类型不能添加读写
	// 只能写
	var chan1 chan<- int
	chan1 <- 10
	//<-chan1              // 不能写
	var chan2 <-chan int // 只读
	<-chan2
	//chan2 <- 29

	fmt.Printf("intChan的值 %v，intchan的地址 %p \n", intChan, &intChan)
	// 向管道写入数据
	intChan <- 10
	num := 20
	intChan <- num
	// 可以取出不要
	<-intChan

	intChan <- 90
	// ！！ 写入数据不能超过管道的长度 否者会报错 deadlock!
	intChan <- 79
	// 管道的长度和容量
	fmt.Printf("长度%v, 容量%v \n", len(intChan), cap(intChan))
	// 取管道里的数据
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	// 在没有使用携程的情况下如果管道数据已经全部取出，在取就会报错
	//num4 := <-intChan
	fmt.Println("num1", num1)
	fmt.Println("num2", num2)
	fmt.Println("num3", num3)
	//fmt.Println("num4", num4)
	fmt.Printf("长度%v, 容量%v \n", len(intChan), cap(intChan))

	// 定义一个存放任意数据类型的管道 3个 创建管道要写容量
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "toom"
	cat := Cat{
		name: "小花猫",
		age:  12,
	}
	// 要想得到管道的第三个元素，需要把前两个推出来
	<-allChan
	<-allChan
	allChan <- cat
	newCat := <-allChan
	fmt.Printf("newCat=%T, newCat=%v \n", newCat, newCat)
	// 因为管道类型是空接口，任何都可以满足，
	// 要调用里面的方法需要类型断言
	fmt.Printf("newCat.name = %v \n", newCat.(Cat).name)
	// 管道的关闭
	numChan := make(chan int, 3)
	numChan <- 10
	numChan <- 20
	//numChan <- 30
	// 关闭管道 就不能再写入数据， 但是可以读取数据
	close(numChan)
	<-numChan
	//numChan <- 30
	int1, ok := <-numChan
	fmt.Println("int1", int1, ok)

	// 管道的遍历
	numChan2 := make(chan int, 200)
	for i := 0; i < 200; i++ {
		numChan2 <- i * 2
	}
	fmt.Println("numChan2", numChan2)
	// 在遍历是如果管道没有关闭会报deadlock 错误
	// 在遍历时如果管道关闭，这回正常遍历数据，遍历完退出遍历
	//for i := 0; i < len(numChan2); i++ {
	//	fmt.Println("v=", i)
	//
	//}
	close(numChan2)
	for v := range numChan2 {
		fmt.Println("v=", v)
	}
}
