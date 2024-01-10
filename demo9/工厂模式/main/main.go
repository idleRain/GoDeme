package main

import (
	"GoDemo/demo9/工厂模式/model"
	"fmt"
)

func main() {
	// 这里会报错，因为 model 包下面结构体是小写开头，无法直接引用
	// s1 := model.student{"小丁", 100.0}

	// 使用工厂模式解决这个问题
	s2 := model.NewStudent("小戴", 99.9) // 得到一个指针
	fmt.Println("学生二：", *s2)
	fmt.Println("学生二的成绩：", s2.GetScore())
}
