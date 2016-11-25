package main
import "fmt"

// TZ 是 int 的别名
type TZ int

func main(){
    var a,b TZ = 3,4
    c := a + b
    fmt.Printf("c has the value: %d",c)

    str := "\nHello," + "go"
    fmt.Printf(str)
}