package gomap

import "fmt"

func Maps() {
	// map 是引用类型 遵守引用类型传递机制
	var stud = make(map[string]map[string]string)
	stud["1"] = make(map[string]string, 2)
	stud["1"]["name"] = "hsueh"
	stud["1"]["age"] = "12"
	stud["2"] = make(map[string]string, 2)
	stud["2"]["name"] = "hsuehly"
	stud["2"]["age"] = "13"
	// for range 遍历复杂的map
	for s, m := range stud {
		fmt.Println("s", s, "m", m)
		for s2, s3 := range m {
			fmt.Println("s2", s2, "s3", s3)

		}

	}
	// 三种创建方式
	// 先声明 在make，  直接make 创建，  直接创建并赋值
	//var a map[string]string
	// map 使用前一定要make创建内存空间
	var a = make(map[string]string, 10)
	a["aum1"] = "a"
	a["bey"] = "f"

	a["cum2"] = "b"
	a["i"] = "d"

	a["num3"] = "c"
	a["v"] = "e"
	fmt.Println(a, "a")
	// 声明完之后可以继续赋值
	heros := map[string]string{"key": "hs", "key2": "klk", "klj": "opo", "loi": "uyt"}
	heros["klg"] = "jip"
	delete(heros, "klg")
	delete(heros, "yui")
	fmt.Println("heros", heros)
	// map 增加，如果key不存在就直接增加，如果key存在，就直接覆盖

	// map 查找
	v, res := heros["klg"]

	fmt.Println("v", v, "f", res)
	fmt.Println("heros的长度", len(heros))
	// map 遍历 只能使用for range 进行遍历， 应为没有下标
	for i, v := range heros {
		fmt.Printf("heros的i %v， heros的v %v \n", i, v)

	}
	//for i := 0; i < len(heros); i++ {
	//	fmt.Println(heros)
	//
	//}
	var numMap = map[int]string{0: "1", 1: "2", 2: "3"}
	for i := 0; i < len(numMap); i++ {
		fmt.Println(numMap[i], "ppp")

	}
}
