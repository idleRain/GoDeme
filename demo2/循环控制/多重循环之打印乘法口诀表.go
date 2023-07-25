package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %v * %v = %v ", j, i, j*i)
		}
		fmt.Println("")
	}

for1:
	for i := 0; i < 4; i++ {
	for2:
		for j := 0; j < 5; j++ {
			if j == 2 {
				// break 可以跳出指定标签，否则跳出当前 for 循环
				break for2
			}
			fmt.Println("j=", j)
		}
		if i == 2 {
			// continue 可以跳过指定标签，否则跳过本次 for 循环
			continue for1
		}
	}
}
