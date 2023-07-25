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
}
