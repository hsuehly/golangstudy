package goreflect

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Student1 struct {
	Name string
	Age  int
}

func reflectTestInt(a interface{}) {
	// 获取reflect 的类型
	rTyp := reflect.TypeOf(a)
	fmt.Printf("rType的类型%v,真正的类型%T \n", rTyp, rTyp)
	// 获取值
	rVal := reflect.ValueOf(a)
	// 基本类型取值
	n1 := 20 + rVal.Int()
	n2 := rVal.Float() // 报错 运行错误 在 int Value 上调用 reflect.Value.Float
	fmt.Printf("n1的值%v \n", n1)
	fmt.Printf("n2的值%v \n", n2)
	fmt.Printf("rVal 的值%v \n", rVal)
	// 转为需要的类型，需要先转为空接口，然后断言成想要的类型
	num2 := rVal.Interface().(int)
	fmt.Printf("num2 的值%v \n", num2)
	// kind 比 type 的范围大
	// type是类型， kind 是类别 可能相同 可能不相同
	fmt.Printf("rTypkind %v rValkind %v \n", rTyp.Kind(), rVal.Kind())

}

// 结构体的反射
func reflectTestStruct(a interface{}) {

	rTyp := reflect.TypeOf(a)
	fmt.Printf("rType的类型%v,真正的类型%T \n", rTyp, rTyp)
	rVal := reflect.ValueOf(a)
	// 接口体取值 用断言来取
	fmt.Printf("rVal 的值%v \n", rVal)
	iV := rVal.Interface()
	fmt.Printf("iV=%v,type=%T \n", iV, iV)
	// 断言 第二个参数可以确定是不是当前类型的
	//stu, ok := iV.(Student1)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("ok=%v stu.name=%v \n", ok, stu.Name)
	}
	fmt.Println("ok", ok)
	fmt.Printf("rTypkind %v rValkind %v \n", rTyp.Kind(), rVal.Kind())

}

// 通过反射修改变量值
func reflectSetValue(a interface{}) {
	// 获取到reflect.value
	rVal := reflect.ValueOf(a)
	fmt.Printf("rVal的值%v \n", rVal)
	fmt.Printf("rVal 的Kind=%v \n", rVal.Kind())
	//rVal.SetInt(120) // 传递指针 直接修改编译不会报错，运行会报错
	/*
		// 可以这样理解Elem()
		num := 10
		var ptr *int = &num
		num2 = *ptr ===> 类似 rVal.Elem()
	*/

	rVal.Elem().SetInt(120)
}
func Reflect() {
	//num := 100
	//reflectTestInt(num)
	//stu := Student{
	//	Name: "小明",
	//	Age:  19,
	//}
	//reflectTestStruct(stu)
	// 传递指针类型
	//reflectSetValue(&num)
	//fmt.Println("num的值", num)
	// test
	var str = "tom"
	fs := reflect.ValueOf(&str)
	fmt.Printf("fs的Kind=%v \n", fs.Kind())
	ok := fs.CanSet()
	fmt.Println("ok", ok)
	fmt.Println("ll", fs.Elem())
	fmt.Printf("fs elem 的类型%T \n", fs.Elem())
	// 只有kind 为当前类型 或者canset 为true 才可以修改，不然panic
	fs.Elem().SetString("jack")

	fmt.Println("str", str)
}
