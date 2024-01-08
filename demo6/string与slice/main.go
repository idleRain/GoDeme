package main

import "fmt"

func main() {
	// string 本质是一个 byte 数组，所以 string 也可以切片
	str := "hello world!"
	strSlice := str[6:] // world!
	fmt.Println("strSlice", strSlice)

	// string 是不可变的，通过下标修改字符串会报错
	// str[0] = '2' // error

	// 如果要修改字符串，需要先转成切片，修改完毕后再转回来
	strArr := []rune(str)
	strArr[0] = '好'
	str = string(strArr)
	fmt.Println("str=", str) // 好ello world!
}

