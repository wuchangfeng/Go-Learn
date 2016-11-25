package main

import (
	"fmt"
)

// 注意后面两个 int 是指定返回值类型
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 获取返回值的后半部分
	_, c := vals()
	fmt.Println(c)
}
