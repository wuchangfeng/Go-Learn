package main

import (
    "fmt"
)

var {
    fN,lN,s string
    i int
    f float32
    input = "56.12/5212/Go"
    format = "%f /%d /%s"
}

func main() {
    fmt.Println("enter your full name: ")
    fmt.Scanln(&fN,&lN)
}