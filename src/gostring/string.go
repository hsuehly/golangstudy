package gostring

import (
	"fmt"
	"strings"
)

func Str() {
	var Str = "kkkkkk"
	strs(Str)

	var strs = "llll"
	str1 := len(strs)
	fmt.Println(str1, "str1")
	var names = "hsuehly"
	// 判断字符串中是否包含指定的字符串
	ishave := strings.Contains(names, "j")
	fmt.Println(ishave, "ishave")
}

var strs = func(str string) {
	fmt.Println("kkkk", str)
}
