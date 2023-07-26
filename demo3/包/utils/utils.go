package utils // 包名通常和文件夹名一致，但不是硬性规定

import "fmt"

var Num int = 12 // 变量首字母大写也可以被其他包使用

// 函数首字母大写可以被其他包使用，否则不能跨包
func Cal(n1 float64, n2 float64, operation string) float64 {
	var res float64
	switch operation {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	default:
		fmt.Println("运算符输入有误")
	}
	return res
}

func SayHi()  {
	fmt.Println("Hi~")
}
