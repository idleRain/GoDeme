package main

import "fmt"

func main()  {
	num1 := 99
	fmt.Printf("num1的内心%T, num1的值%v, num1的地址%v\n", num1, num1, &num1)

	// new 用来分配内存，主要用来分配值类型，int float struct
	// new 函数接受的是一个类型，返回一个该类型的零值
	num2 := new(int)
	// 修改 num2 需要通过指针修改
	*num2 = 100
	fmt.Printf("num2的类型%T, num2的值%v, num2的地址%v, 指向的值%v\n", num2, num2, &num2, *num2)
}
