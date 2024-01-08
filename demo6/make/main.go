package main

import "fmt"

func main() {
	/**
	 * make 是一个内置函数
	 * make 可以用来创建切片、映射或通道
	 * make 有三个参数，param1: 类型；param2: 长度；param3?: 容量
	 * !!! 分配的容量必须不小于长度
	 * 创建好后默认为所有元素分配该类型的零值
	 */

	slice := make([]int, 2)
	fmt.Println("slice", slice)
	fmt.Println("slice长度", len(slice))
	fmt.Println("slice容量", cap(slice))

	// 为切片追加元素
	slice2 := append(slice, 100, 200, 300)
	fmt.Println("slice2", slice2)
	slice2 = append(slice2, slice2...)
	fmt.Println("slice2", slice2)
}
