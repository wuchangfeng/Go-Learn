package main

func main() {
    // 用 goto 替代循环
    i := 0
    HERE:
        print (i)
        i ++
        if i == 5 {
            return 
        }
        goto HERE
}