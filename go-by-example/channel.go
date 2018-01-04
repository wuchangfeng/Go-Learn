package main

import (
	"fmt"
)

func main() {
	msg := make(chan string)

	go func() {
		// 向 msg 通道中发送数据
		msg <- "ping"
	}()

	// 取出 msg 通道中的数据
	msgs := <-msg
	fmt.Println(msgs)
}
