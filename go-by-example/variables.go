package main

import (
	"fmt"
)

func main() {
	var a string = "Hello world"
	fmt.Println(a)

	// 命名且赋值
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 默认打印出 0
	var e int
	fmt.Println(e)

	// 局部变量
	f := "short"
	fmt.Println(f)

	// 常量
	const n = 1000
	fmt.Println(n)
}
