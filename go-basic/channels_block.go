package  main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(1e9)
	fmt.Println(<-ch1)
}

// 生产者
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// 消费者
func suck(ch chan int) {
    for {
        fmt.Println(<-ch)
    }
}
