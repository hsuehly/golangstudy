package gotypeassertion

import "fmt"

type Ponint struct {
	x int
	y int
}
type Usb interface {
	start()
	stop()
}
type Phone struct {
	Name string
}
type Camera struct {
	Name string
}
type Computer struct {
}

func (p Phone) start() {
	fmt.Println("手机启动中。。。。")

}
func (p Phone) stop() {
	fmt.Println("手机关机中。。。。")

}
func (p Phone) call() {
	fmt.Println("手机打电话中。。。")

}
func (c Camera) start() {
	fmt.Println("相机启动中。。。")

}
func (c Camera) stop() {
	fmt.Println("相机关机中。。。")

}
func (com Computer) work(usb Usb) {
	usb.start()
	// 类型断言
	if phone, ok := usb.(Phone); ok {
		phone.call()
	}
	usb.stop()

}

func Assertion() {
	//var a interface{}
	//var point = Ponint{
	//	x: 1,
	//	y: 2,
	//}
	//a = point
	//var b Ponint
	//b = a.(Ponint) // 类型断言
	//fmt.Println(b)
	// 类型断言，带检测
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	y, ok := x.(float64) // 类型断言
	if ok == true {
		fmt.Println("cpnvert sucess")
		fmt.Printf("y 的类型是 %T 值为%v", y, y)
	} else {
		fmt.Println("断言失败")
	}
	fmt.Println("接着执行", y)
	var usbarr [3]Usb
	usbarr[0] = Phone{"vivo"}
	usbarr[1] = Phone{"小米"}
	usbarr[2] = Camera{"尼康"}
	var computer Computer
	for _, usb := range usbarr {
		computer.work(usb)
	}
	fmt.Println("类型断言------------------")
	var n1 float32 = 1.1
	var n2 float64 = 2.33
	var n3 int32 = 30
	var name string = "tom"
	address := "beijing"
	n4 := 300
	var com1 = Computer{}
	var com2 = &Computer{}
	TypeJudge(n1, n2, n3, name, address, n4, com1, com2)
}

func TypeJudge(item ...interface{}) {
	fmt.Println("item", item)
	for i, v := range item {
		fmt.Println(v)
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v \n", i, v)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v \n", i, v)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v \n", i, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v \n", i, v)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v \n", i, v)
		case Computer:
			fmt.Printf("第%v个参数是 Computer 类型，值是%v \n", i, v)
		case *Computer:
			fmt.Printf("第%v个参数是 *Computer 类型，值是%v \n", i, v)

		default:
			fmt.Printf("第%v个参数是类型不确定，值是%v \n", i, v)

		}

	}
}
