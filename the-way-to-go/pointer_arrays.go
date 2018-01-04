package main

import "fmt"

// f 并不能改变数组的值，只是数组的复制内容
func f(a [3]int) { fmt.Println(a) }

// 传递了数组的指针到这里
func fp(a *[3]int) { fmt.Println(a) }

func main() {
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}
