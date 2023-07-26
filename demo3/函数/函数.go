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

	sum, sub := getSumAndSub(2, 1) // 接收多个返回值
	fmt.Printf("sum%v和sub%v:", sum, sub)
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

// 函数可以有多个返回值
func getSumAndSub(n1 int, n2 int) (int, int)  {
	return n1 + n2, n1 - n2
}