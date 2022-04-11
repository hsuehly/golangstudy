package goobjectoriented

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (stu *Student) say() string {
	info := fmt.Sprintf("学生名字：%v 学生性别：%v 年龄: %v ID: %v 分数：%v", stu.Name, stu.Gender, stu.Age, stu.Id, stu.Score)
	return info
}
func ObjectOriented() {
	var stu = Student{
		Name:   "小白",
		Gender: "男",
		Age:    10,
		Id:     01,
		Score:  10.3,
	}
	str := stu.say()
	fmt.Println(str)

}
