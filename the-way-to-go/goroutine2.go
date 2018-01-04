package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "A"
	ch <- "g"
	ch <- "f"
	ch <- "e"
	ch <- "d"
	ch <- "c"
}

func getData(ch chan string) {
	var input string
	time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
}
