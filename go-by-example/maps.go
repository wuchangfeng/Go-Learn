package main

import (
	"fmt"
)

func main() {

	// 初始化方式一
	m := make(map[string]int)
	m["one"] = 2
	m["two"] = 3
	m["three"] = 4
	fmt.Println(m)

	delete(m, "one")
	fmt.Println(m)

	// 初始化方式二
	n := map[string]int{"x": 2, "y": 4, "u": 8}
	fmt.Println(n)

}
