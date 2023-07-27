package main

import "fmt"

func main() {
	// 匿名函数，只执行一次
	func() {
		fmt.Println("我是匿名函数")
	}()
	func(n1 int, n2 int) {
		fmt.Println(n1 + n2)
	}(12, 20)

	// 将匿名函数赋值给变量，可以多次调用
	a := func(n1 int, n2 int) {
		fmt.Println(n1 * n2)
	}
	a(10, 2)
	a(22, 11)
}
