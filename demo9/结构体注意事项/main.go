package main

import "fmt"

/**
 * 如果结构体类型是：指针、slice、map、func，则其零值都是 nil，即还没有分配空间
 * 如果要使用这样的字段，需要先 make，才能使用
 */

type person struct {
	name   string
	age    int
	scores [2]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	m      map[string]string // 映射
	fn     func(s string)    // 方法
}

func main() {
	var p person
	fmt.Println(p)

	if p.slice == nil {
		fmt.Println("p.slice为空")
	}

	// 使用 slice，要先 make，并且指定初始长度
	p.slice = make([]int, 10)
	p.slice[0] = 1

	p.fn = func(s string) {
		fmt.Println(s)
	}
	p.fn("好好好")
}
