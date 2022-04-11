package goinherit

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}
type Pupil struct {
	Student
	beiCheng string
	// 如果父集和子集 都有相同的字段， 那么现寻找子集的如果子集没有，那么会寻找父集
	Age int
}
type Graduate struct {
	Student
	mingZi string
}
type A struct {
	Name string
	Age  int
}
type B struct {
	Name string
	Age  int
}
type C struct {
	// 如果两个匿名结构体都有相同的字段和方法，并且钩体本身没有同名的字段和方法，在访问时就必须指定匿名结构体的名字
	// 如果c 中有有名结构体 那么在访问中必须带上有名结构体的名字
	A
	B
	//b B // 有名机构体
}
type Goods struct {
	Name    string
	Address string
}
type Brand struct {
	Name  string
	Price float64
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}
type Tv3 struct {
	TV2
	int
	//int // 不能有重复的字段 如果要有多个则必须指定字段名称
}

func (stu *Student) ShowInfo() {
	fmt.Printf("名字：%v 年龄：%v 分数：%v", stu.Name, stu.Age, stu.Score)

}
func (p *Pupil) testing() {
	fmt.Println(p.Student.Name, "小学生正在考试中")
	//p.stu.ShowInfo()

}
func (g *Graduate) testing() {
	fmt.Println(g.Name, "大学生正在考试中")

}
func Inherit() {
	var pupil = &Pupil{Student{Name: "hhh", Age: 12, Score: 12}, "sss", 13}
	//pupil.Age = 12
	pupil.Student.Age = 15
	pupil.beiCheng = "小学生"
	fmt.Println("pupil", *pupil)

	fmt.Println("pupil", pupil.Age)

	pupil.testing()
	var c = C{}
	c.B.Age = 1
	var tv = TV{Goods{Name: "创维", Address: "北京"}, Brand{
		Name:  "LG",
		Price: 120.3,
	}}
	var tv2 = TV2{&Goods{"创维", "北京"}, &Brand{
		Name:  "LG",
		Price: 120.3,
	}}
	fmt.Println("tv", tv)
	fmt.Println("tv2", *tv2.Brand)
}
