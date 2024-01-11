package student

type Student struct {
	Name string
	Age int
}

type BigStudent struct {
	Student // 嵌入匿名结构体，相当于继承
	Intelligence int
}
