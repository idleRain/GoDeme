package main

import "fmt"

func main() {
	arr := []string{"小丁", "小戴", "小李", "小何"}
	fmt.Println("请输入要查找的名字")
	name := ""
	_, _ = fmt.Scanln(&name)
	sequentialSearch(arr, name)
}

// 顺序查找
func sequentialSearch(strArr []string, name string)  {
	index := -1
	for i, v := range strArr {
		if v == name {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v, 下标是%v \n", name, index)
	} else {
		fmt.Println("没有找到", name)
	}
}