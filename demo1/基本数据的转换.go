package main

import "fmt"

func main() {
	// go 中没有类型自动转换或隐式转换，必须显式转换
	var a int8 = 12
	var b int32 = int32(a) // a 的值本身不会改变
	fmt.Println("int8转换int32后b的值：", b)

	var c int8 = int8(b) + 127 // warning: 编译通过但运行时会发生溢出错误
	// var d int8 = int8(b) + 128 // error: 编译不通过
	fmt.Println("c的值：", c)

	var e float64 = 3.14159
	var f float32 = float32(e)
	fmt.Println("float64转float32后f的值：", f)
}
