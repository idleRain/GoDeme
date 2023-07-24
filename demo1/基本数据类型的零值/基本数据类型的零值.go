package main

import "fmt"

func main() {
	/**
	基本数据类型的零值（默认值）
	整形的零值 0
	单浮点的零值 0
	双浮点的零值 0
	布尔的零值 false
	字符串的零值 ""
	*/
	var (
		a int
		b float32
		c float64
		d bool
		e string
	)

	fmt.Printf("整形a的零值： %v \n", a)
	fmt.Printf("单浮点b的零值： %v \n", b)
	fmt.Printf("双浮点c的零值： %v \n", c)
	fmt.Printf("布尔d的零值： %v \n", d)
	fmt.Printf("字符串e的零值： %v \n", e)
}
