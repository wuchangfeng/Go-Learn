package main


import (
	"fmt"
	"runtime"
)


func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
// func zeroval(ival int) {
// 	ival = 0
// }

// func zeroptr(iptr *int) {
// 	*iptr = 0
// }

// type geometry interface{
// 	area() float64
// 	perim() float64
// }

// type rect struct{
// 	width,height float64
// }

// type circle struct{
// 	radius float64
// }

// func (r rect) area() float64 {
//     return r.width * r.height
// }
// func (r rect) perim() float64 {
//     return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
//     return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
//     return 2 * math.Pi * c.radius
// }


// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
//     fmt.Println(g.perim())
// }


// func f(from string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from,":",i)
// 	}
// }

// func worker(done chan bool) {
// 	fmt.Println("working")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- false
// }

func main() {

	// done := make(chan bool,1)
	// go worker(done)
	// <-done
	// f("direct")

	// go f("goroutine")

	// var input string
 //    fmt.Scanln(&input)
 //    fmt.Println("done")
    
	// fmt.Println("received", msg2)

	// go func() {
	// 	time.Sleep(time.Second*1)
	// 	fmt.Println("received")
	// }()

	// go func() {
	// 	time.Sleep(time.Second*2)
	// 	fmt.Println("received")
	// }()



	 // for i := 0; i < 2; i++ {
  //       select {
  //       case msg1 := <-c1:
  //           fmt.Println("received", msg1)
  //       case msg2 := <-c2:
  //           fmt.Println("received", msg2)
  //       }
  //   }






	// r := rect{width:3,height:4} 
	// c := circle{radius : 5}
	// measure(r)
	// measure(c)
	// // i := 1
 // //    fmt.Println("initial:", i)
 // //    zeroval(i)
 // //    fmt.Println("zeroval:", i)
	// // zeroptr(&i)
 // //    fmt.Println("zeroptr:", i)
 // //    fmt.Println("pointer:", &i)
	// message := make(chan string)

	// go func() {message <- "ping"}()

	// msg := <-message
	// fmt.Println(msg)

	go say("world")
	say("hello")

}
