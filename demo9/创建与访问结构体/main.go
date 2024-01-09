package main

import "fmt"

// 创建一个结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	// 第一种声明方式
	var cat1 Cat
	cat1.Name = "凉粉"
	cat1.Age = 1
	cat1.Color = "银色"
	fmt.Println(cat1)

	// 第二种方式
	cat2 := Cat{
		Name:  "绿茶",
		Age:   2,
		Color: "咖啡色",
	}
	fmt.Println(cat2)

	// 第三种方式, new 出来的结构体是一个指针
	cat3 := new(Cat)
	(*cat3).Name = "未知" // 标准写法
	cat3.Age = 0 // 简化写法 => (*cat3).Age = 0
	cat3.Color = "未知"
	fmt.Println(*cat3)

	// 第四种方式，也是获得一个指针
	cat4 := &Cat{
		Name:  "小花",
		Age:   99,
		Color: "三花",
	}
	fmt.Println(*cat4)
}
