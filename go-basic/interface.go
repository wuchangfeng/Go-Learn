package main

import (
    "fmt"
)

type IDuck interface{
    Quack()
    Walk()
}

func DuckDance(duck IDuck) {
    for i := 0; i < 3; i++ {
        duck.Quack()
        duck.Walk()
    }
}

type Bird struct{
    Walk()
}

func (b *Bird) Quack() {
    fmt.Println("I am quacking")
}

func main() {
    b := new(Bird)
    DuckDance(b)
}
