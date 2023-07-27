package main

import "fmt"

func main()  {
	f := addUpper()
	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 3
	fmt.Println(f(3)) // 6
}

// 使用闭包实现一个累加器
type fn func(int) int
func addUpper() fn  {
	n := 0 // 闭包使用到的变量不会被内存释放
	return func(n1 int) int {
		n += n1
		return n
	}
}
