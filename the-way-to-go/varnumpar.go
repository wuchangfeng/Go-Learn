package main

import (
	"fmt"
)

func main() {
	x := min(1, 3, 2, 0)
	fmt.Println("the min is: %d\n", x)

	arr := []int{7, 8, 4, 1, 2}
	// 数组传入
	x = min(arr...)
	fmt.Println("The min in the array is: %d", x)
}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]
	//..._ 什么作用
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
