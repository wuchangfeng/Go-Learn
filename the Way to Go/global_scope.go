package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
   // 这里的 a 是全局变量
   a = "O"
   print(a)
}