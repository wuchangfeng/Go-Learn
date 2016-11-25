package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("write", i, "as")

	// 注意没有 break
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	// case 有多个表达式
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
}
