package main

import "fmt"

func main() {
	students := [...]string{"小丁", "小戴", "小李", "小何"}

	// 常规 for 循环遍历
	for i := 0; i < len(students); i++ {
		fmt.Printf("下标是%v, 值是%v；", i, students[i])
	}
	fmt.Println()

	// for range 遍历
	for index, value := range students {
		fmt.Printf("下表是%v, 值是%v；", index, value)
	}
}
