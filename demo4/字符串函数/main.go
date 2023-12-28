package main

import (
	"fmt"
	"strconv"
	"strings"
	)

func main()  {
	str1 := "hi哇"
	// 中文占用三个字节
	fmt.Println("str1长度=", len(str1)) // 5

	str2 := "hello哇"
	str3 := []rune(str2)
	// 想要打印带中文字符串真实长度，需要转切片
	fmt.Println("str1长度=", len(str3)) // 6

	/* 字符串转整数 */
	n, err := strconv.Atoi("66")
	if err == nil {
		fmt.Println("转换成功， n =", n)
	} else {
		fmt.Println("转换失败", err)
	}
	fmt.Printf("")

	/* 整数转字符串 */
	str4 := strconv.Itoa(45)
	fmt.Printf("整数转字符串str4=%v\n", str4)

	/* 字符串转 byte */
	bytes := []byte("hello world")
	fmt.Printf("字符串转byte, bytes=%v\n", bytes)

	/* byte 转字符串 */
	str5 := string([]byte{97, 98, 99})
	fmt.Printf("byte 转字符串, str5=%v\n", str5)

	/* 查找字符串中是否包含指定的字符 */
	b := strings.Contains("hello, world", "or") // 返回一个布尔值
	fmt.Printf("查找字符串中是否包含指定的字符%v\n", b) // true

	/* 统计字符串中出现过几个指定字符 */
	strCount := strings.Count("hello, world", "l") // 返回一个整型
	fmt.Printf("统计字符串中出现过几个指定字符%v\n", strCount) // 3

	/* 不区分字母大小写的比较 */
	fmt.Println(strings.EqualFold("Abc", "abc")) // true ，返回一个布尔值

	/* 查询字符串第一次出现的下标 */
	fmt.Println(strings.Index("hello, world", "ll")) // 2 ,返回一个整型,如果没有返回 -1
}
