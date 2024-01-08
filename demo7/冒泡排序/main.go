package main

import "fmt"

func main() {
	arr := []int{24, 69, 80, 57, 13}
	fmt.Println("排序前：",arr)
	bubbleSort(arr)
	fmt.Println("排序后：",arr)
}

// 冒泡排序
func bubbleSort(arr []int) {
	temp := 0
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				temp = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}
}