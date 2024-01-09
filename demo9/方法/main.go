package main

import "fmt"

type A struct {
	Num int
}

// 前面括号里面的 A，表示这个方法跟 A 绑定
func (a A) test() {
	// 前面的形参 a 表示 A 结构体实例，值拷贝 修改后不会影响到原来的
	fmt.Println("test(), Num=", a.Num)
}

func (a *A) edit() {
	// 获得一个指针，可以直接修改原值
	a.Num *= 10
}

func main() {
	a := A{10}
	a.test()

	b := A{Num: 999}
	b.test()

	a.edit() // 这里 a 的 Num 属性会被修改
	a.test()
	// test() // error! 方法只能被结构体实例调用
}
