package main

import (
	"fmt"
)

func main() {

	p2 := Add()
	fmt.Println("sum of 2 + 3 is : %d", p2(2))
}

// 注意最外层也是匿名函数
func Add() func(a int) int {
	return func(b int) int {
		return b + 3
	}
}
