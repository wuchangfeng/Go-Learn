package main

import "fmt"

type IntVector []int

// 返回值定义了一个 s 值得考虑
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}
