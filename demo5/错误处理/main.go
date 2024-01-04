package main

import "fmt"

func main()  {
	/**
	 * Go 中可以抛出一个 panic 异常
	 * 然后在 defer 中通过 recover 捕获异常
	 */

	test(true) // 这行代码一定会报错
	// 如果上面代码出错了且没有处理异常，则不会执行后面的
	fmt.Println("test()之后该执行的代码")
}

// 传入 true 代码会报错
func test(isErr bool)  {
	defer func() {
		err := recover() // 内置函数 recover() 可以捕获代码块里面的异常
		if err != nil {
			fmt.Printf("test()执行错误，错误是%v\n", err)
		}
	}()

	var dangerousValues int
	if isErr {
		dangerousValues = 0
	} else {
		dangerousValues = 1
	}
	newNum := 100 / dangerousValues
	fmt.Println("newNum的值=", newNum)
}