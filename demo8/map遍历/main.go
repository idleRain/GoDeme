package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"no1": "001",
		"no2": "002",
		"no3": "003",
	}

	// 用 for range 遍历
	for key, value := range m {
		fmt.Printf("key:%v, value:%v \n", key, value)
	}
}
