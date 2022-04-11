package goconst

import "fmt"

func Consts() {
	// 常量在声明的时候必赋值
	// 常量声明后就不能修改
	// 常量只能修饰 bool 数值类型(int,float系列),string 系列
	// 常量可以通过首字母大小写来控制访问的范围
	const a = 1.08
	const b = "yyy"
	const c = true
	var arr = []int{1, 2, 3, 4}
	fmt.Println(arr)
	//const d = arr // 不能把数组赋值给常量
	//const e = func() {
	//
	//}
	const (
		f = 1
		g = 4
		h = iota
	)

}
