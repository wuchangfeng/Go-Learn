package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
    // 函数内部变量，改变没影响
   a := "O"
   print(a)
}