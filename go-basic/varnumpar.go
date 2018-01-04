package main

import (
    "fmt"
)

func main() {
    x := min(1,34,4,2)
    fmt.Println("The min is %d:",x)
}

func min(a...int) int {
    if len(a)==0 {
        return 0
    }

    min := a[0]

    for _,v := range a{
        if v < min{
            min = v
        }
    }

    return min

}
