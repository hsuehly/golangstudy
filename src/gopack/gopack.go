package gopack

import "fmt"

type Person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *Person {
	return &Person{Name: name}

}
func (p *Person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}
func (p *Person) GetAge() int {
	return p.age
}

//
//func Pack() {
//
//}
