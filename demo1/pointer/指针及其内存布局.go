package main
import (
	"fmt"
)

func main() {
	var i int = 10
	fmt.Println("i的值=", i) // 10
	fmt.Println("i的地址值=", &i) // 0xc000016088

	// 获取指针类型所指向的值，需要用 * 号声明
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr) // 0xc000016088
	fmt.Printf("ptr的地址值=%v\n", &ptr) // 0xc00000a030
	fmt.Printf("ptr所指向的值=%v\n", *ptr) // 10
}