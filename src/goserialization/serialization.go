package goserialization

import (
	"encoding/json"
	"fmt"
)

// 结构体序列化

type Marshal struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func serialStruct() {
	marshal := Marshal{
		Name:     "牛魔王",
		Age:      1000,
		Birthday: "1000-10-10",
		Sal:      70000.00,
		Skill:    "牛魔拳",
	}
	data, err := json.Marshal(&marshal)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(data), "data")
}

// map 序列化

func testMap() {
	var a map[string]interface{}
	// 使用map 切片，先make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["Age"] = 30
	a["address"] = "洪崖洞"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(data), "data")
}

// 切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 20
	m1["address"] = "北京"
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "marry"
	m2["age"] = 18
	m2["address"] = "南京"
	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("data", string(data))
}

// 基本结构序列化  意义不大
func inttype() {
	var num = 134.90
	data, err := json.Marshal(num)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("data", string(data))
}

// 对数组序列化
func testarr() {
	var arr [5]string
	arr[0] = "jack"
	arr[2] = "mary"
	arr[1] = "gher"
	arr[3] = "ioio"
	data, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("err", err)

	}
	fmt.Println("data", string(data))

}

//结构体反序列化

type Marshals struct {
	Name     string `json:"Name"`
	Age      int
	Birthday string
	Sal      float64
	Skill    string
	Mk       string
}

func unMarshalStruct() {
	var str = `{"name":"牛魔王","age":1000,"birthday":"1000-10-10","sal":70000,"skill":"牛魔拳"}`
	var marshal Marshals
	err := json.Unmarshal([]byte(str), &marshal)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("data", marshal)
}

// map 反序列化
func unMarshalMap() {
	var str = `{"Age":30,"address":"洪崖洞","name":"红孩儿"}`
	var a map[string]interface{}
	// 反序列化
	// 注意： 反序列化map，不能需要make，因为make操作封装到 unmarshal中
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("a", a)
}

// 切片反序列化
func unMarshalSlice() {
	var str = `[{"address":"北京","age":20,"name":"jack"},{"address":"南京","age":18,"name":"marry"}]`
	var slice []map[string]interface{}
	// 反序列化
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("slice", slice)
	fmt.Println(slice[0]["name"])
}

func Serialization() {
	// 序列化
	//serialStruct()
	//testMap()
	//testSlice()
	//inttype()
	//testarr()
	// 反序列化
	//unMarshalStruct()
	//unMarshalMap()
	unMarshalSlice()
}
