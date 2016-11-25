package main

import "fmt"

func f() (ret int) {
	defer func() {
		fmt.Printf("%d", ret)
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}
