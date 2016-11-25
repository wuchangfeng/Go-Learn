package main

// 显示类型定义
const b string = "adb"
// 隐方式类型定义
const c = "adb"

// 将 a b 都声明为指针类型
var a, b *int

// go 能在编译时候自动确定变量类型，而不是 Python之类在运行时确定
var a = 15
var b = false
var str = "Go says hello to the world!"


// 局部类型变量，在函数体内部声明
a := 1



import (
    "fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init(){}

 func main() {
    var a int 
    Func1()
    fmt.println(a)    
}

func (t T) Method1() {
    //...
}

func Func1() {
    //
}