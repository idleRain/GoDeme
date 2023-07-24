package main

import (
	"fmt"
)

func main() {
	// 浮点数分为 float32(单精度) 和 float64(双精度)
	// 小数点后位数过多可能会造成精度丢失!
	// float32 范围为 3.4e38
	var a float32 = 3.14159
	fmt.Printf("浮点数a, 类型 %T", a)

	// float64 范围为 1.8e308
	var b float64 = 123.456789
	fmt.Printf("\n浮点数b, 类型 %T", b)

	// 如果不具体声明类型或者让类型自动推断，则默认为 float64
	var c = 321.321321321321 // type = float64
	fmt.Printf("\n浮点数c, 类型 %T", c)
}
