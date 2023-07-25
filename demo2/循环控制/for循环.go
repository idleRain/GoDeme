package main

import "fmt"

func main()  {
	/** Golang 中没有 while 和 do while 循环，但是可以用 for 循环来实现 **/

	// 第一种写法
	for i := 0; i < 4; i++ {
		fmt.Println("我操牛逼！")
	}

	// 第二种写法
	j := 10
	for j > 0 {
		fmt.Println("真的牛逼啊", j)
		j--
		if j == 6 {
			continue // 跳过本次循环
		}
		if j == 3 {
			break // 跳出循环
		}
	}

	// 第三种写法，通常搭配 break 使用
	var k int
	for {
		k++
		fmt.Println("我在无限循环~~~")
		if k >= 5 {
			fmt.Println("我被中断了！！")
			break
		}
	}

	// for range 遍历字符串，是按字符的方式遍历
	str := "Hello world!"
	for index, val := range str{
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

	// 如果遍历的字符串中有中文
	str2 := "我操牛逼！"
	str3 := []rune(str2) // 需要先将字符串转成切片
	for o := 0; o < len(str3); o++{
		fmt.Printf("%c \n", str3[o])
	}
}
