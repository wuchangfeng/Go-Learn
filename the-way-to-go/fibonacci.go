package main

import (
	"fmt"
)

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fib(i)
		fmt.Printf("fib(%d) is: %d\n", i, result)
	}
}

func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-2) + fib(n-1)
	}
	return
}
