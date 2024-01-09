package main

import (
	"fmt"
	"sort"
)

func main() {
	// map 默认是无序的
	m := map[int]int{
		1: 1,
		8: 8,
		4: 4,
		9: 9,
	}
	fmt.Println("没有排序的map", m)
	// 先将 key 放进一个切片排序，然后按照切片排序顺序输出
	var keys []int
	for k := range m{
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)

	// 这时候输出就是有序的
	for _, key := range keys{
		fmt.Printf("map[%v] = %v \n", key, m[key])
	}
}
