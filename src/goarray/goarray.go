package goarray

import "fmt"

func test1(arr [3]int) {
	fmt.Println("test1 arr6", arr)
	arr[0] = 88
	//fmt.Println("test1, arr" arr)
	fmt.Println("test1 arr6 改变之后", arr)

}

func test2(arr *[3]int) {
	fmt.Println("test2 arr6", arr)
	(*arr)[0] = 88
	//fmt.Println("test1, arr" arr)
	fmt.Println("test2 arr6 改变之后", arr)

}

func Arrays() {
	var numarr1 = [3]interface{}{1, 2, "ooo"}
	fmt.Printf("numarr1的地址值 %v \n", &numarr1[0])
	var numarr2 = [3]int{1, 2, 3}
	fmt.Println(numarr2, "numarr2")
	var numarr3 = [3]int{4, 5, 6}
	fmt.Println(numarr3, numarr3)
	numarr4 := [...]int{1: 300, 0: 1, 2: 900}
	//numarr4[5] = 1111
	fmt.Println(numarr4, "numarr4")
	var numarr5 = [...]string{"haha", "你好", "hello"}
	for i, s := range numarr5 {
		fmt.Println(i, s, "index", "value")

	}
	arr3 := [3]int{1, 2}
	//arr4 := [...]interface{}{1, "3", true, numarr2}
	fmt.Println(arr3, "arr3")
	//fmt.Println(arr4[3][1], "arr4")
	// 长度是数据类型的一部分
	arr5 := [...]int{1, 2, 3}
	//test1(arr5)
	test2(&arr5)
	fmt.Println("main, arr5", arr5)
}
func ArraySlice() {
	arr1 := [5]int{1, 2, 3}
	// 切片长度， 包头不包尾
	// 从下标为0 开始结束于下标为3， 不包含4
	arrslice := arr1[0:4]
	fmt.Println("arrslice", arrslice)
	fmt.Println("arrslice 的长度", len(arrslice))
	fmt.Println("arrslice 的容量", cap(arrslice))
	fmt.Println("arr1[0] 的地址", &arr1[0])

	fmt.Println("arrslice 的地址", &arrslice[0])
	arrslice[1] = 45
	fmt.Println("arr1", arr1)
	fmt.Println("arrslice", arrslice)
}
