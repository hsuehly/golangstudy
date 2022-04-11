package gointerface

import (
	"fmt"
	"math/rand"
	"sort"
)

// interface 默认是一个指针， 没有初始化默认值就是nil
// interface{} 没有任何方法， 所有类型都实现了空接口

type Usb interface {
	start()
	stop()
}
type AInterfac interface {
	say()
}
type Stu struct {
	Name string
}

func (stu Stu) say() {
	fmt.Println("stu say()")
}

type Phone struct {
}
type Camera struct {
}
type Computer struct {
}

func (com *Computer) work(usb Usb) {
	usb.start()
	usb.stop()

}
func (p *Phone) start() {
	fmt.Println("手机重启。。。。")

}
func (p *Phone) stop() {
	fmt.Println("手机关机。。。。")
}
func (c *Camera) start() {
	fmt.Println("相机重启。。。。")
}
func (c *Camera) stop() {
	fmt.Println("相机关机。。。。")
}

// interface 继承

type BInterfac interface {
	test1()
	test2()
}
type CInterfac interface {
	test2()
	test3()
}
type DInterfac interface {
	BInterfac
	CInterfac
	test2()
}
type Stu2 struct {
}

func (stu Stu2) test1() {
	fmt.Println("test1")

}

func (stu *Stu2) test2() {
	fmt.Println("test2")
}

func (stu Stu2) test3() {
	fmt.Println("test3")

}

// 接口编程
// 声明 hero结构体

type Hero struct {
	Name string
	Age  int
}

// 声明hero 切片

type HeroSlice []Hero

//实现interface 接口

// 实现len 方法

func (hs HeroSlice) Len() int {
	return len(hs)

}

// less 这个方法是使用什么标准进行排序

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age

}

// 交换值

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func Interfaces() {
	var com = Computer{}
	var p = Phone{}
	com.work(&p)
	var stu Stu
	//stu.Name = "jjj"
	var c AInterfac = stu
	//var b = stu
	c.say()
	var n Stu2
	var m DInterfac = &n
	m.test2()
	var t interface{} = stu
	fmt.Println("t", t)
	var g = []int{1, 4, 5, 78, -12}
	sort.Ints(g)
	fmt.Println(g, "g")
	// 定义接口体切片
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	for _, hero := range heros {
		fmt.Println(hero)

	}
	sort.Sort(heros)
	fmt.Println("排序后的值-----------------")
	for _, hero := range heros {
		fmt.Println(hero)

	}
}
