package main

import "fmt"

func main() {
	fmt.Println("长度为10的斐波那契数列", fibonacci(2))
}

// 返回一个指定长度的斐波那契数列切片
func fibonacci(n int) []int  {
	fbn := make([]int, n)
	fbn[0] = 1
	if n == 1 {
		return fbn
	}
	fbn[1] = 1
	for i := 2; i < n; i++ {
		fbn[i] = fbn[i - 1] + fbn[i - 2]
	}
	return fbn
}