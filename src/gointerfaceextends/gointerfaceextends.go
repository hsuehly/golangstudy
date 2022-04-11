package gointerfaceextends

import "fmt"

type Monkey struct {
	Name string
}

func (mo *Monkey) climbing() {
	fmt.Println(mo.Name, "会爬树")
}

// 声明接口

type Birds interface {
	fly()
}
type Fish interface {
	swimming()
}

// 让titlemonkey 实现bidrs 接口
//当A结构体继承了B结构体，那么A结构体就自动继承了B结构体的字段和方法，并且可以直接使用
// 当A结构体需要扩展功能时，同时不希望破坏继承关系，则去实现某个接口即可，因此我啃可以认为实现接口是对继承机制的补充
/*
接口可继承的解决的解决问题不同
1.继承的价值主要在于，解决代码的复用性，可维护性
2.接口的继承价值主要在于，设计，设计各种规范（方法，让其他自定义类去实现这些方法
接口比继承更加灵活， 继承是满足 is -a 的关系， 而接口只需满足 like -a 的关系
接口在一定程度上实现代码解耦
*/

type TitleMonkey struct {
	Monkey // 继承
}

func (tmo *TitleMonkey) fly() {
	fmt.Println(tmo.Name, "通过学习，学会了飞行")
}
func (tmo *TitleMonkey) swimming() {
	fmt.Println(tmo.Name, "通过学习，学会了游泳")
}

func InterfaceExtends() {
	monkey := TitleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.fly()
	monkey.swimming()
}
