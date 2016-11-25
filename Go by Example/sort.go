package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "w", "d"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	s := sort.StringsAreSorted(strs)
	fmt.Println("Sorted:", s)
}
