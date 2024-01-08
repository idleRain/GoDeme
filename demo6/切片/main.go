package main

import "fmt"

func main() {
	/**
	 * 切片是数组的引用
	 * 可以当作是动态的数组
	 * 切片是引用类型，而数组是基本类型
	 */
	arr := [...]int{1, 2, 3, 4, 5, 6}
	// 从下标 1 开始引用，到下标 3 结束(不包括3)
	s := arr[1: 3]
	fmt.Println("intArr=", arr)
	fmt.Println("slice元素=", s)
	fmt.Println("slice元素个数=", len(s))
	// 使用内置函数 cap 获得容量，注：切片容量是可变的
	fmt.Println("slice容量=", cap(s))
}
