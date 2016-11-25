package main

import (
	"fmt"
)

func main() {
	var min, max int
	min, max = MinMax(7, 5)
	fmt.Println("Minnum is : %d, Maxnum is: % d ", min, max)
}

// 返回值类型和参数名称放在后面
func MinMax(a int, b int) (min int, max int) {
	if a < b {
		max = b
		min = a
	} else {
		min = b
		max = a
	}
	return
}
