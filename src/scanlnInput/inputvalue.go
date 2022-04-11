package scanlnInput

import "fmt"

func GetValue() {
	var name string
	var age int
	var sel float32
	var isPass bool
	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sel)
	//fmt.Println("请输入是否通过考试")
	//fmt.Scanln(&isPass)
	//fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水为 %v \n 是否通过考试 %v \n", name, age, sel, isPass)
	// 使用空格隔开获取用户输入的内容
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sel, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水为 %v \n 是否通过考试 %v \n", name, age, sel, isPass)

}
