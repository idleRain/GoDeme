package main

import "fmt"

func main() {
	var num byte
	fmt.Println("请输入一个小于3的正整数：")
	fmt.Scanln(&num)

	// Golang 中的 switch 不需要 break !!
	switch num {
	case 1:
		fmt.Println("你输入了1")
	case 2:
		fmt.Println("你输入了3")
	default:
		fmt.Println("输入有误")
	}

	// switch 后面也可以不带表达式，类似 if else 分支来使用
	age := 20
	switch {
	case age == 10:
		fmt.Println("age == 10")
		fallthrough // switch 分支穿透
	case age == 20:
		fmt.Println("age == 20")
	case age > 20:
		fmt.Println("age > 20")
	default:
		fmt.Println("没有匹配到")
	}
}
