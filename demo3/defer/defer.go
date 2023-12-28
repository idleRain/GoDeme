package main

import "fmt"

func main()  {
	res := sum(10, 20)
	fmt.Println("res=", res) // 第 4 步打印
}

func sum(n1 int, n2 int) int  {
	/**
	defer 会把后面的代码压入异步栈，并在函数执行完之后再去执行栈里的代码
	先入栈的最后出栈
	 */
	defer fmt.Println("n1=", n1)// 第 3 步打印
	defer fmt.Println("n2=", n2) // 第 2 步打印
	ok := n1 + n2
	fmt.Println("ok=", ok) // 第 1 步打印
	return ok
}
