package main

import (
	"fmt"
)

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	r := sum(arr[:])
	fmt.Println(r)
}
