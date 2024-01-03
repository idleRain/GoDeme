package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 获取当前时间
	fmt.Printf("值:%v；\n类型:%T\n", now, now)

	// 获取年月日时分秒
	fmt.Printf("年：%v\n", now.Year())
	fmt.Printf("月：%v\n", now.Month()) // 这里得到的是英文月
	fmt.Printf("月：%v\n", int(now.Month())) // 这里得到的是数字月
	fmt.Printf("日：%v\n", now.Day())
	fmt.Printf("时：%v\n", now.Hour())
	fmt.Printf("分：%v\n", now.Minute())
	fmt.Printf("秒：%v\n", now.Second())

	// 格式化时间
	// %d 不补零， %02d 补零
	dateStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("年月日时分秒：", dateStr)

	// 格式化时间第二种方式， format 写法是固定的
	fmt.Printf("%v\n", now.Format("2006-01-02 15:04:05"))
}
