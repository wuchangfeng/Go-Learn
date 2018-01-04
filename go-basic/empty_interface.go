package main

import (
    "fmt"
)

var i = 5
var str = "ABC"

type Pserson struct{
    name string
    age int
}

type Any interface{}

func main() {
    var val Any
    val = 5
    fmt.Printf("val has the value:%v\n",val)
    val = str
     fmt.Printf("val has the value:%v\n",val)
}




