package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaInft Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaInft = sq1

	// 判断 areaIntf 是什么类型
	if t, ok := areaInft.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
