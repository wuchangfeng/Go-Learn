package main

import (
	"fmt"
)

func main() {

	var i = 5
	// 必须是 printf
	fmt.Printf("An integer:%d,its location in memory is:%p\n", i, &i)

	var intP *int
	intP = &i
	// intP 表示地址，*intP 表示取出数值
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

}
