package main

import "fmt"

func main() {
	/**
	 * Golang 中有三个位运算
	 * 按位与 & : 两位全为1，结果为1，否则为0
	 * 按位或 | : 两位有一个为1，结果为1，否则为0
	 * 按位异或 ^  : 一个为0 一个为1，结果为1，否则为0
	 */

	var a byte = 2 // 补码 [0000 0010]
	var b byte = 3 // 补码 [0000 0011]

	fmt.Println("a&b: ", a&b)   // 2
	fmt.Println("a|b: ", a|b)   // 3
	fmt.Println("a^b: ", a^b)   // 1
	fmt.Println("-2^2: ", -2^2) // -4

	/**
	 * Golang 中有两个移位运算
	 * >> 位右移：低位溢出，符号位不变，并用符号位补溢出的高位
	 * << 位左移：符号位不变，地位补 0
	 */

	var c byte = 1
	var d byte = 2
	fmt.Println("c>>d: ", c>>d) // 0
	fmt.Println("c<<d: ", c<<d) // 4
}
