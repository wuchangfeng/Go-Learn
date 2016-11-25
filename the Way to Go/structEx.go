package main

import (
	"fmt"
)

type Person struct {
	ages int
	num  float32
	name string
}

func main() {
	/* way 1
	p := new(Person)
	p.ages = 23
	p.num = 15.0
	p.name = "Jack"
	fmt.Println("hello")
	*/

	/*  Way 2
	p := &Person{23, 15.0, "Jack"}
	fmt.Println("the person name is %s:", p.name)
	*/

	var p Person
	p = Person{23, 15.0, "Jack"}
	fmt.Println("the person name is %s:", p.name)
}
