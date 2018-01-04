package main

import (
	"fmt"
)

func main() {
	// ... 代表变长数组
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}
