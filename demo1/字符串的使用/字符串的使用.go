package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// go 中的字符串统一使用 UTF-8 编码格式
	// 字符串用双引号包裹
	var a string = "牛蛙牛蛙！"
	// a[2] = '好'  // error: 不可直接对字符串进行修改！
	fmt.Println("a的值：", a)
	fmt.Println("a占用的字节：", unsafe.Sizeof(a))

	// 字符串拼接使用 + 号
	b := "好耶！" + a + "哈哈哈……"
	fmt.Println("b的值：", b)

	// 反引号包裹的字符串类似于 JavaScript 中的模板字符串
	// 能够保留一些特殊字符和缩进
	c := `
      ⠀⠀⠀⠀⠰⢷⢿⠄
      ⠀⠀⠀⠀⠀⣼⣷⣄
      ⠀⠀⣤⣿⣇⣿⣿⣧⣿⡄
      ⢴⠾⠋⠀⠀⠻⣿⣷⣿⣿⡀
      ○ ⠀  ⢀⣿⣿⡿⢿⠈⣿
      ⠀⠀⠀⢠⣿⡿⠁⠀⡊⠀⠙
      ⠀⠀⠀⢿⣿⠀⠀⠹⣿
      ⠀⠀⠀⠀⠹⣷⡀⠀⣿⡄
      ⠀⠀⠀⠀⣀⣼⣿⠀⢈⣧.
    `
	fmt.Println("c的值：", c)
}
