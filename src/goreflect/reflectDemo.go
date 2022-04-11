package goreflect

import (
	"fmt"
	"reflect"
)

// Monster 定义一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
	//address string
}

// Print Print方法
func (m Monster) Print() {
	fmt.Println("-----start----")
	fmt.Println(m)
	fmt.Println("-----end----")
}
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2

}
func (m Monster) set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}
func TestStruct(a interface{}) {
	rtype := reflect.TypeOf(a)
	fmt.Println("rtype", rtype)
	rval := reflect.ValueOf(a)
	fmt.Println("rval", rval)
	rkind := rval.Kind()
	fmt.Println("rkind", rkind)
	if rkind != reflect.Struct {
		fmt.Println("不是结构体，退出函数")
		return
	}
	// 获取该结构体有几个字段
	num := rval.NumField()
	fmt.Printf("该结构体有%v个字段 \n", num)
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("file %d:值为=%v \n", i, rval.Field(i))
		//读取struct标签，注意要通过reflect.Type的Field来获取tag标签， Type 和value 的返回值不一样
		// 如果没有tag 则是"" 值
		tagVal := rtype.Field(i).Tag.Get("json")
		fmt.Printf("Field%d，Tag=%v \n", i, tagVal)
	}
	//methodName := rval.MethodByName("Print")
	//methodName.Call(nil)
	// 这种私有的方法访问不到
	numofMethod := rval.NumMethod()
	fmt.Printf("结构体有%v个方法 \n", numofMethod)
	// 调用结构体中的方法
	// 方法的排序默认是按照函数名(ASCII码)排序
	rval.Method(1).Call(nil) // 获取第二个方法调用它，没有参数传nil
	// 调用结构体第一个方法
	var params []reflect.Value // 声明reflect.Value 切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := rval.Method(0).Call(params) // 传入的类型是reflect.Value 切片， 返回的也是reflect.Value切片
	fmt.Println("res", res[0].Int())

}
func ReflectDemo() {
	// 创建实列
	a := Monster{Age: 200, Name: "牛魔王", Score: 30.90}
	TestStruct(a)
}
