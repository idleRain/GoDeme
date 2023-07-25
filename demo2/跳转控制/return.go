package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		if i == 5 {
			// 中断当前函数体的运行
			return
		}
		fmt.Println(i)
	}

	fmt.Println("我是函数最后一行") // 因为上面 return 了， 所以这行不会运行
}
