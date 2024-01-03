package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
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
	fmt.Printf("查找字符串中是否包含指定的字符%v\n", b)        // true

	/* 统计字符串中出现过几个指定字符 */
	strCount := strings.Count("hello, world", "l") // 返回一个整型
	fmt.Printf("统计字符串中出现过几个指定字符%v\n", strCount)    // 3

	/* 不区分字母大小写的比较 */
	fmt.Println(strings.EqualFold("Abc", "abc")) // true, 返回一个布尔值

	/* 查询字符串第一次出现的下标 */
	fmt.Println(strings.Index("hello, world", "ll")) // 2, 返回一个整型,如果没有返回 -1

	/* 查询字符串最后一次初选的未知 */
	fmt.Println(strings.LastIndex("hello, world", "l")) // 10, 返回一个整型,如果没有返回 -1

	/**
	 * 替换字符串
	 * 第一个参数：字符串
	 * 第二个参数：希望被替换的字符串
	 * 第三个参数：替换后的字符串
	 * 第四个参数：要替换几次， -1 表示全部替换
	 */
	str6 := strings.Replace("golang牛蛙牛蛙", "牛", "好", 1) // golang好蛙牛蛙
	fmt.Println(str6)

	/* 将字符串切割成数组 */
	strArr := strings.Split("hello,world,golang", ",")
	fmt.Printf("strArr=%v\n", strArr) // strArr=[hello world golang]

	/* 字符串转小写 */
	fmt.Println(strings.ToLower("hEllO")) // hello

	/* 字符转大写 */
	fmt.Println(strings.ToUpper("hEllO")) // HELLO

	/* 将字符串左右两边的空格去掉 */
	fmt.Println(strings.TrimSpace("  go language  ")) // go language

	/* 将左右两边指定的字符串去掉 */
	/* TrimLeft 去掉左边的; TrimRight 去掉右边的 */
	fmt.Println(strings.Trim(" !go! language! ", " !")) // go! language

	/* 判断字符串是否以指定字符串开头或结尾, 返回一个布尔值 */
	fmt.Printf("开头%v\n", strings.HasPrefix("hello world", "he")) // true
	fmt.Printf("结尾%v\n", strings.HasSuffix("hello world", "rl")) // false
}
