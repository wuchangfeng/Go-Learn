package main

import (
	"fmt"
)

func intSeq() func() int {

	i := 0
	// 匿名函数里面的不会管外面的 i
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	// output is 1
	fmt.Println(nextInt())
	// output is 2
	fmt.Println(nextInt())
	// output is 3
	fmt.Println(nextInt())

	// 了确认这个状态对于这个特定的函数是唯一的
	newInts := intSeq()
	// output is 1
	fmt.Println(newInts())
}
