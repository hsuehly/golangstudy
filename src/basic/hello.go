package basic

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

var c int = 9
var (
	name = "hsueh"
	age  = 20
)
var city = "北京"

func BaseFunc() {

	var j = 1279 + 'a'
	var k string = `sssss`
	//var cv byte = 'a'
	k = "kkkkkkkk"
	fmt.Println("k的地址", &k)
	fmt.Println(len(k))
	fmt.Println("j", j)
	fmt.Printf("j 的类型是 %T", j)
	fmt.Println("天灵吧不\r张飞")
	valuetype("kkkkkkk")
	yuan()
	enum()
	readFiles()
	fmt.Println(grand(88))

}
func yuan() {
	a := 3
	b := 4
	//c := 8
	//var c int = 8

	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c, "yuan")
}
func enum() {
	const (
		javascript = 1 << (10 * iota)
		typescript
		java
	)
	fmt.Println(javascript, typescript, java)
	q, r, _ := div(13, 3, 4)
	fmt.Println(q, r, "返回多个参数")
}
func valuetype(str string) {
	fmt.Println(str)
	var a int = 10
	var b string = ""
	var c int
	var d string
	fmt.Println(a, b)
	fmt.Println(c, d)

}
func readFiles() {
	const filename = "abc.txt"
	str, errer := os.Open(filename)
	if errer != nil {
		fmt.Println(errer)
	} else {
		fmt.Sprintln(str, "str")
	}
	value, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err, "err")

	} else {
		fmt.Printf("%s\n", value)

	}
}
func grand(score int) string {
	f := ""
	switch {
	case score < 60:
		f = "A"
	case score < 80:
		f = "B"
	case score < 90:
		f = "D"
	default:
		panic(fmt.Sprintf("Wrong Scord %d", score))
	}
	return f
}
func div(a, b, c int) (q, r, t int) {
	q = a / b
	r = a % b
	t = c
	return
}
