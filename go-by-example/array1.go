package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func fp(a *[3]int) {
	// 打印出不同的值
	fmt.Println(a)
	fmt.Println(*a)
}

func main() {

	per := &Person{"allenwu", 13}

	fmt.Println(per.age)
	a := [...]string{"a", "b", "c"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	for i := 0; i < 3; i++ {
		// 传入 a 的首地址，首地址指向正个数组
		fp(&[3]int{i, i * i, i * i * i})
	}

	// 超过了数组长度，指定数组类型为 Int
	// b := [3]int{1, 2, 3, 4, 5}
}
