package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	// 因为 test 函数执行出错了，所以后面的代码不会继续执行
	fmt.Println("test()后面的代码")
}

func readFile(fileName string) (err error) {
	if fileName != "test.txt" {
		// 如果文件名不正确，则返回一个错误
		return errors.New("文件读取错误")
	} else {
		return nil
	}
}

func test()  {
	err := readFile("test2.txt")
	if err != nil {
		panic(err) // 内置函数 panic() 可以输出错误并终止程序
	}
	fmt.Println("test()继续执行。。。")
}
