package main

import (
	"fmt"
)

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, seasons := range seasons {
		fmt.Println("Season %d is : %s\n", ix, seasons)
	}

	var season string
	// _ 忽略索引
	for _, season = range seasons {
		fmt.Println("%\n", season)
	}
}
