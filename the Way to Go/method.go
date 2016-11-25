package main

import (
	"fmt"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Println("The sum of two is %d:", two1.AddThem())
}

// (tn *TwoInts) 注意这里的参数，吸引结构体 two1 去调用method
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}
