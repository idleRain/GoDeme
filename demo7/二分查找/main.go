package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 8, 10, 13}
	binarySearch(arr, 6, 0, len(arr) - 1)
	binarySearch(arr, 99, 0, len(arr) - 1)
}

// 二分查找
func binarySearch(arr []int, num int, leftIndex int, rightIndex int)  {
	// 二分查找要求要查找的数组是一个有序数组
	if leftIndex > rightIndex {
		// 如果左边大于右边则说明找不到
		fmt.Println("没有找到", num)
		return
	}
	middleIndex := (leftIndex + rightIndex) / 2
	if arr[middleIndex] > num {
		binarySearch(arr, num, leftIndex, middleIndex - 1)
	} else if arr[middleIndex] < num {
		binarySearch(arr, num, middleIndex + 1, rightIndex)
	} else {
		// 找到了
		fmt.Printf("找到%v, 下标是%v \n", num, middleIndex)
	}
}
