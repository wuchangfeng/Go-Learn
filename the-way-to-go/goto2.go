package main

import (
	"fmt"
)

func main() {
	a := 1
	goto TARGET
	b := 9
	// 错误写法 ,标签的位置要固定
TARGET:
	b += a
	fmt.Println("a is %v *** b is %v", a, b)
}
