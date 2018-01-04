package main

func main(){
    println("before")
    greeting()
    println("after")
}

func greeting(){
    println("In greeting")
}
/* 错误的写法 括号
func greeting()
{
    println("In greeting")
}*/