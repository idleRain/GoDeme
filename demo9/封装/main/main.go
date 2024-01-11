package main

import (
	"GoDemo/demo9/封装/model"
	"fmt"
)

func main() {
	p := model.NewPerson("小丁")
	p.SetAge(18)
	p.SetSal(300000)
	fmt.Println(p.ShowPerson())
}