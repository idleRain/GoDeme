package main

import "fmt"

func main()  {
	// 压栈和弹栈
	test(4) // 2, 2, 3
	test2(4) // 2, 2

	fmt.Println(fibonacci(8))
}

func test(num int)  {
	if num > 2 {
		num--
		test(num)
	}
	fmt.Println("num: ", num)
}

func test2(num int)  {
	if num > 2 {
		num--
		test(num)
	} else {
		fmt.Println("num: ", num)
	}
}

// 斐波那契数
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}