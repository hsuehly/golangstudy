package main

import "main/goredis"

//func init() {
//	fmt.Println("我先执行0")
//}
//func init() {
//	fmt.Println("我先执行1")
//}
//
//var b = 20
//var c = 10

func main() {
	//basic.BaseFunc()
	//var i = 'a'
	//var p = 'a'
	//fmt.Println(&i, &p, "llll")
	//fmt.Println("i的地址", &i)
	//var k *int = &i
	//var c = *k
	//i = 100
	//
	//fmt.Println("k的地址", k)
	//fmt.Println("k 的值", *k)
	//fmt.Println("c 的值", c)
	//
	//var u = 9
	//var t = 10
	//var ptr = &u
	//*ptr = 100
	//fmt.Println("u 最终值", u, *ptr)
	//var n float64 = 10 / 4.0
	//
	//fmt.Println(u >= t)
	//
	//fmt.Println(10.0/4, "kkk")
	//fmt.Println(n, "kkkjj")
	// 取模的本质 a%b = a-(a/b)*b
	//zhizen.ZhiFc()
	//scanlnInput.GetValue()
	//bitOperation.Operations("开始启动")
	//processControl.ProcessControl()
	//var n = 10
	//test.BreakDemo()
	//test.Test(&n)
	//fmt.Println("main n ", n)
	//recursion.Test(4)
	//var r = 10
	//var u = recursion.Test2
	//recursion.Dizhi(&r)
	//fmt.Println("r", r, &r)
	//fmt.Printf("u的类型 %T, dizhi的类型 %T \n", u, recursion.Test2)
	//u(10, 20, 30, 40)
	//var fun = closure.Bibao()
	//fmt.Println(fun(3))
	//fmt.Println(fun(4))
	//fmt.Println(fun(5))
	//var fun1 = closure.Bibao
	//fmt.Println(fun1()(3))
	//fmt.Println(fun1()(4))
	//fmt.Println(fun1()(5))
	//
	//var b = 10
	//var c = 20
	//var res = closure.Defaut(&b, &c)
	//fmt.Println("res", res)
	// 常量
	//goconst.Consts()
	// go 字符串
	//gostring.Str()
	// go 时间函数
	//gotime.Times()
	// 测试 循环耗费时间
	//start := time.Now().Unix()
	//gotime.TestTime()
	//end := time.Now().Unix()
	//fmt.Printf("函数执行耗费时间%v秒 \n", end-start)
	// 内置函数 new
	//num1 := 100
	//fmt.Printf("num1 的类型 %T num1 的值 %v num1 的内存地址 %v \n", num1, num1, &num1)
	//num2 := new(int)
	////*num2 = 100
	//fmt.Printf("num2 的类型 %T num2 的值 %v num2 的内存地址 %v num2指针指向的值%v \n", num2, num2, &num2, *num2)
	//goerror.Error()
	//fmt.Println("hahaahhaha")
	//goarray.Arrays()

	//goarray.ArraySlice()
	//goslice.Slices()
	// map
	//gomap.Maps()
	// go 结构体
	//gostruct.Structs()
	// go 方法
	//gomethod.Methods()
	//var ghj = gomethod.Person{
	//	Name: "jjj",
	//	Age:  0,
	//}
	//ghj.Test()
	// go 面向对象
	//goobjectoriented.ObjectOriented()
	// go 封装
	//var p = gopack.NewPerson("小白")
	//fmt.Println("p", *p)
	//fmt.Println(p.GetAge(), "age")
	// go 继承
	//goinherit.Inherit()
	// go 接口
	//gointerface.Interfaces()
	// go 接口继承
	//gointerfaceextends.InterfaceExtends()
	// go 类型断言
	//gotypeassertion.Assertion()
	// 家庭收入程序
	//testMyAccount.TestMyAccount()
	//FamilyAccount.NewFamAccount().MainMenu()
	// 客户管理系统
	//view.CustomerView()
	// 文件操作
	//gofile.Files()
	// 序列化操作
	//goserialization.Serialization()
	// go 并行， 携程
	//goroutine.Main()
	//
	//goruntime.Runtime()
	// 管道
	//gochannel.Channel()
	//gochannel.ChannelDemo()
	// 用管道和协程求素数
	//gochannel.GetPrime()
	// 普通方法耗时
	//gochannel.TestPrime()
	// 管道细节
	//gochannel.Details()
	// 管道中的错误
	//gochannel.Error()
	// 反射
	//goreflect.Reflect()
	// 反射用例
	//goreflect.ReflectDemo()
	// net
	//service.Server()
	// go 连接reids
	goredis.Redis()
	// go redis 连接池
	//goredis.RedisPool()
}
