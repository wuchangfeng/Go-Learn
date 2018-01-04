package main

import (
	"fmt"
)

func a() {
	i := 0
	defer fmt.Println(i)
	defer fmt.Println(2)
	defer fmt.Println(3)
	i++
	//fmt.Println(i)
	return
}

func b() int {
	i := 0
	defer fmt.Println(i)
	i++
	return 1
}

func main() {
	//b()
	a()
}
