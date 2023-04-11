package main

import (
	"fmt"
)

func main() {
	// int 声明的变量表示有符号整形， 分为 int8 int16 int32 int64
	/**
	  int8 取值范围为 -128 ~ 127
	  int16 取值范围为 -32768 ~ 32767
	  int32 取值范围为 -2147483648 ~ 2147483647
	  int64 取值范围为 -9223372036854775808 ~ 9223372036854775807
	  若声明时不添加后缀或让类型自动推断，则会根据操作系统自动声明为 int32 或 int64
	*/
	var a int8 = 100
	fmt.Println("int8类型的a：", a)

	// uint 声明的变量表示无符号整形， 分为 uint8 uint16 uint32 uint64
	/**
	  uint8 取值范围为 0 ~ 255
	  uint16 取值范围为 0 ~ 65535
	  uint32 取值范围为 0 ~ 4294967295
	  uint64 取值范围为 0 ~ 18446744073709551615
	  若声明时不添加后缀或让类型自动推断，则会根据操作系统自动声明为 uint32 或 uint64
	*/
	var b uint16 = 20
	fmt.Println("uint16类型的b：", b)

	// byte 类型与 uint8 完全相同， 可以当作是别名
	var c byte = 255
	fmt.Println("byte类型的c：", c)
}
