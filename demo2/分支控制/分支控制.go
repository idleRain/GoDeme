package main

import "fmt"

func main()  {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	// go 中没有三元表达式，必须使用分支控制
	if age >= 18 {
		fmt.Println("恭喜你已经成年了，坐牢不打折了！")
	} else {
		fmt.Println("小兔崽子你还没成年！")
	}

	var year int
	fmt.Println("请输入一个年份，判断是不是闰年：")
	fmt.Scanln(&year)
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println(year, "年是闰年")
	} else {
		fmt.Println(year, "年不是闰年")
	}
}
