package main

import (
    "fmt"
)

func MultiReply(a,b int,reply *int) {
    * reply = a * b
}




func main() {
    n := 0
    reply := &n
    MultiReply(10,5,reply)
      fmt.Println("Multiply:", *reply) // Multiply: 50
}

