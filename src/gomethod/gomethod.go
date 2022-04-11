package gomethod

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// 推荐传递指针进去，如果数据量很大，会导致值很多

func (per *Person) test(n string) {
	fmt.Printf("per的地址 %p \n", per)
	fmt.Println(per.Name)
	per.Name = "kkk"

}

//type

func (per Person) speak(str string) {
	fmt.Println(str)

}

type Student struct {
	Name string
	Age  int
}

// 如果一个类型实现了String()这个方法，那么fmt.Println(&..) 默认会调用这个变量的string() 进行输出
// 方法不管调用形式如何，真正决定是值拷贝还是地址拷贝看这个方法和那个类型绑定
// 如果是和值类型，比如（p Person),则是值拷贝，如果是指针类型（p *Person）则是地址拷贝
func (S *Student) String() string {
	str := fmt.Sprintf("名字= %v 年龄= %v", S.Name, S.Age)
	return str

}

func Methods() {
	var p = Person{"lll", 12}
	p.test("lll")
	p.speak("hah")
	fmt.Println("p", p)
	fmt.Printf("p的地址 %p \n", &p)
	var stu = Student{"小明", 12}
	fmt.Println(&stu)
}
