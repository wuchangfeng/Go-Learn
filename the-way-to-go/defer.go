package main

import (
	"fmt"
)

func main() {
	a()
	a2()
	//function1()
}

/*
a 与 a2 的效果类似 在return 后继续执行一个语句
*/
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func a2() {
	i := 0
	//defer fmt.Println(i)
	i++
	return
	fmt.Println(i)
}

// 输出 123
func function1() {
	fmt.Println("1")
	defer function2()
	fmt.Println("2")
}

func function2() {
	fmt.Println("3")
}
