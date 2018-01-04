package main

import (
	"errors"
	"fmt"
)

//var errNotFOund error = errors.New("Not found error")

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math-square root of negative number")
	}
	return 1, errors.New("math-square root of negative number")
}

func main() {
	//fmt.Printf("error:%v", errNotFOund)

    err := Sqrt(-1)

    err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
