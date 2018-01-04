package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	j := 1

	for j <= 3 {
		fmt.Println(j)
		j = j + 1
	}

	// 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环
	for {
		fmt.Println("loop")
		break
	}

}
