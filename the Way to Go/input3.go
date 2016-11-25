package main

import "fmt"

var (
	firstName, lastName, s string
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	//fmt.Printf("Hi %s %s!\n", firstName, lastName)
}
