package main

import "fmt"

// 创建一个结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func edit(cat Cat)  {
	// 结构体是值类型, 这里修改并不会影响到原来的
	cat.Name = "绿茶"
}

func main() {
	cat1 := Cat{
		Name:  "凉粉",
		Age:   1,
		Color: "银色",
	}
	fmt.Println("猫猫的名字：", cat1.Name)
	fmt.Println(cat1.Name + "的信息：", cat1)
	edit(cat1)
	// 不受影响
	fmt.Println(cat1.Name + "的信息：", cat1)
}
