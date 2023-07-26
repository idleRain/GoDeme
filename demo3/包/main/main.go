package main

import (
	"GoDemo/demo3/包/utils" // 在 import 中导入要使用的包
	"fmt"
)

func main() {
	fmt.Println("请输入两个数和一个运算符，得到运算结果：")
	var n1 float64
	var n2 float64
	var operation string
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&operation)
	fmt.Printf("最后的结果是：%.2f", utils.Cal(n1, n2, operation))

	fmt.Println("utils下面的变量Num：", utils.Num)
	utils.SayHi()
}
