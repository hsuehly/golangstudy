package recursion

import (
	"fmt"
)

func Test(n int) {
	if n > 2 {
		n--
		Test(n)
	}
	fmt.Println("n", n)
}
func Dizhi(k *int) int {
	var l = k
	var h = &l
	var g = &h
	***g = 100
	fmt.Println("k", *k, k)
	fmt.Println("l", *l)
	fmt.Println("h", **h)
	fmt.Println("g", ***g)
	return 10

}

func Test2(sum int, args ...int) {
	fmt.Println("sum", sum)
	fmt.Println("args", args)
	fmt.Printf("%T", args)
}
