package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}

func (p *person) SetSal(sal float64) {
	if sal > 1000000 || sal < 100 {
		fmt.Println("薪水范围不正确")
	} else {
		p.sal = sal
	}
}

func (p *person) SetAge(age int) {
	if age > 180 || age < 0 {
		fmt.Println("年龄范围不正确")
	} else {
		p.age = age
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) ShowPerson() string {
	return fmt.Sprintf("姓名：%v，年龄：%v，薪资：%v", p.Name, p.age, p.sal)
}
