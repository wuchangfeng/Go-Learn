package main

import (
	"fmt"
)

// func badCall(){
// 	//panic("bad end")
// }

// func test() {
// 	defer func(){
// 		if e:= recover(); e != nil{
// 			fmt.Printf("Panicing %s \r\n",e)
// 		}
// 	}()
// 	badCall()
// 	fmt.Println("After bad call \r\n")
// }

func vals()(int,int) {
	return 3,7
}


func sum(num...int) {
	fmt.Println(num)
}


func intSeq() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}


func main() {
	

	nextInt := intSeq()
	fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
	//sum(1,2,3,4)
	// a,c := vals()
	// fmt.Println(a)
	// fmt.Println(c)
    // var a [5]int
    // fmt.Println(a)

    // a[4] = 100
    // fmt.Println(a)

    // b := [5]int{1,2,3,4,5}
    // fmt.Println(b)
    // s := make([]string,3)
    // fmt.Println("enp:",s)

    // s = append(s,"d")
    // fmt.Println(s)

    // nums := []int{2,3,4}
    // sum := 0
    // for _,num := range nums{
    // 	sum += num
    // }

    // fmt.Println(sum)
}
