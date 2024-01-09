package main

import "fmt"

func main() {
	// map 在声明的时候必须 make，用于分配空间
	// 不使用 make 声明的 map，没有分配内存，无法使用
	// 声明语法为 map[keyType]valueType
	mapA := make(map[string]string)
	mapA["name"] = "牛逼"
	// 在声明时直接赋值
	mapB := map[string]int{
		"age": 18,
	}
	mapB["power"] = 9999999
	mapC := map[string]map[string]string{
		"student1": {
			"name": "小丁",
			"age":  "18",
		},
		"student2": {
			"name": "小何",
			"age":  "19",
		},
	}
	// 删除，指定的 key 存在时会删除，不存在时会不进行任何操作
	delete(mapC, "student2")
	fmt.Println(mapA)
	fmt.Println(mapB)
	fmt.Println(mapC)
}
