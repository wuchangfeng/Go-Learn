package main

import (
	"fmt"
)

func main() {
	//var value int
	//var isPresenter bool

	map1 := make(map[string]int, 5)
	map1["One"] = 23
	map1["Two"] = 34
	map1["Three"] = 45

	// 注意必须是 Printf
	for key, value := range map1 {
		fmt.Printf("key is: %s - value is: %d\n", key, value)
	}
}
