package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
	//num := <-in
	//fmt.Println(num)
}

func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
}
