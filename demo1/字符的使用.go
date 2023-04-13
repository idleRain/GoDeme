package main

import "fmt"

func main() {
	/* go 中的字符是保存的字符码值 */

	var a byte = 'a'       // 字符使用单引号包裹
	fmt.Println("a的值:", a) // 结果为 97, 因为字符 a 的码值为 97
	// 若想正确输出 a 则需要格式化输出
	fmt.Printf("a格式化后的结果: %c \n", a) // 结果为 a

	// 因为保存的是十进制码值，所以可以进行数学运算
	var b int = 100 + 'b'
	fmt.Println("b的值:", b) // 结果为 198

	/**
	值得注意的是，因为 byte 的范围为 -128 ~ 127
	所以声明码值超过 127 的字符变量则会溢出
	这里可以用 int 类型来声明
	*/
	var c int = '棒'
	fmt.Println("c未格式化的结果:", c)      // 结果为 26834
	fmt.Printf("c格式化后的结果: %c \n", c) // 结果为 棒
}
