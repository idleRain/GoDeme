package main

import "fmt"

func main() {
	fmt.Println("请输入两个数和一个运算符，得到运算结果：")
	var n1 float64
	var n2 float64
	var operation string
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&operation)
	fmt.Println("最后的结果是:", cal(n1, n2, operation))
}

func cal(n1 float64, n2 float64, operation string) float64 {
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
