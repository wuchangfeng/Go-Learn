package main

import (
    "fmt"
    // 都是声明了没有使用
    //"runtime"
)

func main() {
   //var a string = "abc"
   fmt.Println("hello, world")

   var a int
   var b int32
   a = 15
   // 编译错误，类型不匹配
   b = a + a
   b = b + 5
}