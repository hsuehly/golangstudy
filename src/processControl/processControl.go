package processControl

import (
	"fmt"
)

var age int

func ProcessControl() {
	//var age int
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//if age > 18 {
	//	fmt.Println("你年龄过大")
	//}
	//if age > 18 {
	//	fmt.Println("你年龄过大")
	//} else {
	//
	//}
	//var k int
	//for i := 0; i <= 10; i++ {
	//	fmt.Println("Hsueh", i)
	//	k = i
	//	fmt.Println(k, "k在循环中")
	//}
	//var j = 0
	//for j <= 10 {
	//	fmt.Println("jHsueh", j)
	//
	//	j++
	//}
	//var l = 0
	//for {
	//	fmt.Println("jjjjjjjj")
	//	if l <= 10 {
	//		fmt.Println("hahahah")
	//	} else {
	//		break
	//	}
	//	l++
	//}
	//var str string = "hello,golang, 你好"
	//var str2 = []rune(str)
	//for i := 0; i < len(str2); i++ {
	//	fmt.Println(i, str2[i])
	//	fmt.Printf("字符 %c \n", str2[i])
	//}
	//for index, val := range str {
	//	fmt.Println(index, val)
	//	fmt.Printf("下标 %v, 值 %c \n", index, val)
	//}
	// while
	var o = 1
	for {
		if o >= 10 {
			break
		}
		fmt.Println("hghghghgh")
		o++
	}
	// do while
	var p = 1
	for {
		fmt.Println("klklk")
		p++
		if p >= 10 {
			break
		}
	}
	var totalNum = 20
	for i := 1; i <= totalNum; i++ {
		for k := 1; k <= totalNum-i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || j == totalNum {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", i, j, i*j)
		}
		fmt.Println()
	}
}
