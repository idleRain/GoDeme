package main

import "fmt"

func main() {
	/**
	 * 计算六个数的总和、平均值、最大值
	 */
	// 定义一个 float64 类型数组 arr
	// [] 里面是数组大小， 使用 [...] 可以自动推导
	// 数组如果只定义了类型长度，没有赋初始值，那么初始值就是该类型的零值
	arr := [...]float64{3.0, 5.0, 1.1, 2.4, 2.0, 50.0}
	total := 0.0
	maxValue := arr[0]
	// for 循环累加每一项，得到总数
	for i := 0; i < len(arr); i++ {
		total += arr[i]
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	avg := fmt.Sprintf("%.2f", total/float64(len(arr)))

	fmt.Printf("总数:%v\n", total)
	fmt.Printf("平均值:%v\n", avg)
	fmt.Printf("最大值:%v\n", maxValue)
}
