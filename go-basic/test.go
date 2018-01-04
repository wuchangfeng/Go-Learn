package main
import (
    "fmt"
)
func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello go")
    }

    // 取地址符号
    var i1 = 5
    fmt.Println("the memory is %p:",&i1)
    var intP *int = &i1
    fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
    
}
