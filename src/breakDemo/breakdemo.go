package breakDemo

import (
	"fmt"
)

var p = 20

func init() {
	fmt.Println("我先执行2", p)
}
func init() {
	fmt.Println("我先执行3")
}
func BreakDemo() {
	//var count int
	//rand.Seed(time.Now().Unix())
	//
	//for {
	//	n := rand.Intn(100) + 1
	//	fmt.Println("n=", n)
	//	count++
	//	if n == 99 {
	//		break
	//	}
	//}
	//fmt.Println("生成99 一共", count)
	//lable1:
	//for i := 1; i <= 10; i++ {
	//lable2:
	//	for j := 0; j <= 10; j++ {
	//		if j == 2 {
	//			break lable2
	//		}
	//		fmt.Println("j", j)
	//	}
	//
	//}
	//here:
	//for i := 1; i < 4; i++ {
	//	for j := 1; j < 4; j++ {
	//		if i == 2 {
	//			//continue here
	//			return
	//		}
	//		fmt.Println("i", i)
	//	}
	//
	//}
	//	fmt.Println("ok1")
	//	fmt.Println("ok2")
	//label2:
	//	fmt.Println("ok3")
	//	fmt.Println("ok4")
	//	goto label2
	//	fmt.Println("ok5")
	//	fmt.Println("ok6")
	//	fmt.Println("ok7")
	//	//label:
	//	fmt.Println("ok8")
	//	fmt.Println("ok9")
	//	fmt.Println("ok10")
}
