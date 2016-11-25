package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("settability of v:", v.CanSet())

	v = reflect.ValueOf(&x)

	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	// 使用该函数，间接的使用指针
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}
