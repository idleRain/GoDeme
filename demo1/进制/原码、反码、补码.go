package main

import (
	"fmt"
)

func main() {
	/*
		对于有符号的而言，二进制的最高位表示正负
		0 表示正数, 1表示负数

		正数的原码、反码、补码都一样
		byte 1 => 原码 [0000 00001] 反码 [0000 00001] 补码 [0000 00001]

		负数的反码：原码位不变，其他位取反  负数的补码：它的反码 +1
		byte -1 => 原码 [1000 0001] 反码 [1111 1110] 补码 [1111 1111]

		计算机运算的时候，都是以补码的方式运算的
	*/

	fmt.Println("牛蛙牛蛙")
}
