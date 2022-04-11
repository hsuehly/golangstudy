package breakDemo

import "fmt"

func Test(n *int) {
	*n++
	fmt.Println("n", *n)
}

//func BreakDemo() {}

func init() {
	fmt.Println("我先执行4")
}
