package main

func main()  {
	// goto 可以无条件的跳转到指定的行
	// 不推荐使用 goto ，可能会造成程序运行的混乱
	println("我是1")
	println("我是2")
	println("我是3")
	goto label1  // 跳转到 label1 的位置(11 行)
	println("我是4")
	println("我是5")
	label1:
	println("我是6")
	println("我是7")
}
