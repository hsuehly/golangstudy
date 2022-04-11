package goslice

import "fmt"

func Slices() {
	var arr3 = []int{1, 2, 3, 4, 5}
	var slice4 = arr3[:3]
	var slice5 = arr3[1:]
	var slice6 = arr3[:]
	fmt.Println("slice4", slice4)
	fmt.Println("slice5", slice5)
	fmt.Println("slice6", slice6)
	var slice1 []float64
	fmt.Println("slice1", slice1)
	slice2 := make([]int, 2, 4)
	slice2[0] = 2
	slice2[1] = 4
	slice2 = append(slice2, arr3...)
	//slice2[3] = 6
	//slice2[4] = 7
	//slice2[5] = 8
	fmt.Println("slice2", slice2)
	var slice7 = make([]int, 10)
	copy(slice7, arr3)
	arr3[2] = 100
	fmt.Println("slice6", slice7)
	var arr4 = []int{1, 2, 3, 5, 6}
	var slice8 = make([]int, 1)
	// 代码不回报错， 只会拷贝切变的最大长度
	copy(slice8, arr4)
	fmt.Println("slice8", slice8)
	var str = "hsueh"
	var slice9 = str[1:4]
	fmt.Println("slice9", slice9)
	//str[1] = 'y'
	// 转成byte字节后可以处理数字和字母， 但是不能处理中文
	// 原因 byte 是按照一个字节来处理，而中文是三个字节， 因此会出现乱码
	// 解决 将string转为 run[] 数组即可， 因为run是按字符处理，兼容中文
	arr6 := []rune(str)
	arr6[1] = '北'
	str = string(arr6)
	fmt.Println("str=", str)
	fbnSlice := fbn(10)
	fmt.Println("fbnSlice", fbnSlice)
	fmt.Println("二维数组-------------------")
	var arr2 [5][6]int
	arr2[1][3] = 3
	arr2[3][2] = 5
	arr2[4][5] = 7
	fmt.Println(arr2, "arr2二维数组")
	var arr31 [5][6][3]int
	fmt.Println(arr31, "arr31三维数组")
	var arr32 = [2][2]int{1: {1, 3}, 0: {9, 5}}
	fmt.Println("arr32", arr32)

	for i := 0; i < len(arr2); i++ {
		//fmt.Println(arr2[i])
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], "")

		}
		fmt.Println()

	}
}
func fbn(n int) []uint64 {
	slices := make([]uint64, n)
	slices[0] = 1
	slices[1] = 1
	for i := 2; i < n; i++ {
		slices[i] = slices[i-1] + slices[i-2]

	}
	return slices

}
