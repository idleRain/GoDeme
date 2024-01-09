package main

import "fmt"

func main() {
	// map 切片可以使 map 个数为动态
	mapSlice := []map[string]string{
		{
			"name": "小小丁",
			"age":  "18",
		},
		{
			"name": "小小李",
			"age":  "19",
		},
	}
	fmt.Println(mapSlice)

	// 追加
	mapSlice = append(mapSlice, map[string]string{
		"name": "小小戴",
		"age":  "99",
	})

	for _, v := range mapSlice {
		for key, value := range v {
			fmt.Printf("key:%v, value:%v \n", key, value)
		}
	}
}
