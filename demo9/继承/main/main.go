package main

import (
	"GoDemo/demo9/继承/student"
	"fmt"
)

func main() {
	bs := student.BigStudent{}
	bs.Student.Name = "小丁"
	bs.Student.Age = 18
	bs.Intelligence = 250

	fmt.Println(bs)
}
