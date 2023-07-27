package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a原来的值：", a)
	numAdd(&a) // 将 a 的地址值传过去
	fmt.Println("a修改后的值：", a) // 此时 a 本身的值发生改变

	fmt.Println("多个参数", add(1, 2, 3, 4, 5))

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

// 通过指针的方式操控传进来的原值
func numAdd(num *int)  {
	*num += 10
}

// 剩余参数， args 实际是一个切片
func add(args... int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}