package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 布尔类型只有两个值，true 与 false，同其他语言一样
	var a bool = true
	// 与 JavaScript 不同的是， 数字、字符串、对象在判断中无法自动转成布尔类型
	fmt.Println("a的值：", a)
	fmt.Println("布尔a占用的字节：", unsafe.Sizeof(a))
}
