package main

import (
	"fmt"
)

func main() {
	// value 为 func() int 类型的
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	// 整形都被映射到函数地址
	fmt.Println(mf)

	// 定义 Map 容量
	map2 := make(map[string]float32, 100)
}
