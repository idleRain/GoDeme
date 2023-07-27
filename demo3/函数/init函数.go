package main

import "fmt"

// 执行顺序 全局变量 > init函数 > main函数
var a int = test()

// 每个源文件都可以有一个 init 函数，它会在 main 函数之前调用
func init()  {
	fmt.Println("init...") // 2
}

func main()  {
	fmt.Println("main...") // 3
}

func test() int  {
	fmt.Println("test...") // 1
	return 0
}
