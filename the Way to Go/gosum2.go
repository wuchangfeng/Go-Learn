package main

import (
	"fmt"
)

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	// 取出通道中的数值
	r := <-c
	fmt.Println(r)
}
