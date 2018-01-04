package main

import (
	"fmt"
)

func main() {

	ms := make(chan string, 2)
	ms <- "num1"
	ms <- "num2"

	fmt.Println(<-ms)
	fmt.Println(<-ms)
}

/*

因为这个通道是有缓冲区的，即使没有一个对应的并发接收方，我们仍然可以发送这些值

*/
