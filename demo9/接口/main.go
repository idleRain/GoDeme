package main

import "fmt"

// 接口，定义一种规范，接口里面的方法需要被实现
type Inter interface {
	start()
	stop()
}

type A struct {}
type B struct {}
type C struct {}

func (c *C) working(i Inter)  {
	i.start()
	i.stop()
}

func (a *A) start()  {
	fmt.Println("A开始工作")
}

func (a *A) stop()  {
	fmt.Println("A停止工作")
}

func (a *B) start()  {
	fmt.Println("B开始工作")
}


func main() {
	a := A{}
	//b := B{}
	c := C{}
	// a 实现了接口
	c.working(&a)
	// b 没有实现接口
	// c.working(&b) // error
}
