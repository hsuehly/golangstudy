package gostruct

import (
	"encoding/json"
	"fmt"
)

// 结构体默认值 int 为0 string为 "" boolean false

type Cat struct {
	Name  string
	Age   int
	Color string
}

// 结构体为值类型 默认是值拷贝 如果字段是引用类型，不进行重新创建空间，那么会把原来的地址传递给下一个
// 指针本身的地址还是连续的，但是指向的地址不一定是连续的

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针 new
	slice  []int             //切片
	map1   map[string]string // 切片
	// 如果结构体的字段类型是： 指针，slice，map 的零值都是nil，即还没有分配空间
	// 如果需要使用这样的字段，需要先make，才能使用
}
type Animal struct {
	Name  string
	Age   int
	Hoppy map[string]string
}
type Per struct {
	Name string
	Age  int
}
type Per2 Per

// 结构体标签， 一般用来前后交互数据

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func Structs() {
	cat2 := new(Cat)
	var cat3 = Cat{Name: "小花"}
	var cat1 Cat
	cat1.Name = "小白"

	fmt.Println("cat2的地址", &cat2, "cat1的地址", &cat1, &cat1.Name)
	fmt.Printf("cat2的地址%p cat1的地址%p \n", &cat2, &cat1.Color)
	cat1.Name = "小白"
	cat1.Age = 10
	cat1.Name = "白色"
	//cat1.op = "lll"
	//cat1.a = "uuii"
	fmt.Println("cat1", cat1)
	fmt.Println(cat2, "cat2")
	fmt.Println(cat3, "cat3")
	var k int = 10
	fmt.Println("k的地址", &k)
	var arr1 = []int{1, 2, 3, 4}
	var slicearr = arr1[:]
	fmt.Println("arr1的地址", &arr1)
	fmt.Println("arrslice的地址", &slicearr[0])
	//slice 和 map 没有分配空间 默认值为nil
	var person Person
	fmt.Println("person", person)

	// 使用slice 之前一定要先make分配空间
	person.slice = make([]int, 5)
	person.slice[0] = 12
	// 使用map前一定要先make分配空间
	person.map1 = make(map[string]string)
	person.map1["key"] = "Hui"
	if person.ptr == nil {
		fmt.Println("ok1")
	}
	if person.slice == nil {
		fmt.Println("ok2")
	}
	if person.map1 == nil {
		fmt.Println("ok3")
	}
	fmt.Println("person", person)
	//
	var dog Animal
	dog.Hoppy = make(map[string]string)
	dog.Hoppy["eat"] = "骨头"
	var tiger = dog
	tiger.Hoppy = make(map[string]string)
	tiger.Hoppy["drink"] = "水"
	fmt.Println("dog", dog)
	fmt.Println("tiger", tiger)
	// 结构体赋值 4 种方式
	// 方式1
	var per Per
	// 方式2
	var per2 = Per{}
	per2.Name = "bui"
	per2.Age = 12
	var per3 = Per{"jkl", 12}
	// 方式3
	var per4 = new(Per)
	// per4 是一个指针 标准的给字段赋值方式
	// (*pr4).Age = 12 也可以这样写 pre4.Age = 12
	//  原因： 为了程序员使用方便， 底层还是会转变为（*pre4）的写法
	(*per4).Age = 12
	per4.Name = "ghj"
	// 方式4
	var per5 = &Per{}
	per5.Name = "io"
	var per6 = per5
	per6.Name = "op"
	var per7 = &Per{}
	per7.Name = "ui"
	fmt.Printf("per %v per2 %v per3 %v per4 %v per5 %v per6 %v per7 %v \n", per, per2, per3, per4, per5, per6, per7)

	var monstr = Monster{"牛魔王", 500, "牛魔拳"}
	jsonstr, err := json.Marshal(monstr)
	if err != nil {
		fmt.Println("处理错误")
	}

	fmt.Println("monstr", monstr)
	fmt.Println("jsonstr", string(jsonstr))
	// 结构体创建的方式
	// 1
	stu1 := Per{"小白", 30}
	var stu2 = Per{
		Name: "小白",
		Age:  10,
	}
	// 2
	var stu3 = &Per{"小白", 30}
	stu4 := &Per{
		Name: "小白",
		Age:  4,
	}
	fmt.Println(stu1, stu2, *stu3, *stu4)
}
