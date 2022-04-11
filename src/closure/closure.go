package closure

import "fmt"

// 闭包

func Bibao() func(int) int {
	var n = 10
	defer fmt.Println("nnnnnn", n)
	defer fmt.Println("nnnnnn2222", n)
	return func(sum int) int {
		defer println("sumsum", sum)
		defer println("sumsum2222", sum)
		n = n + sum
		return n
	}
}
func Defaut(n *int, n1 *int) int {
	// defer 进行的是值拷贝，用指针也是一样
	defer fmt.Println("*n", *n)
	defer fmt.Println("*n1", *n1)
	*n++
	*n1--
	fmt.Println("n", *n)
	fmt.Println("n1", *n1)
	return *n + *n1

}
