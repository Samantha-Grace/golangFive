package main

import "fmt"

// pointer points to specific memory location
func main() {
	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)                             // prints out: x is 10
	fmt.Println("My first pointer is", myFirstPointer) // prints out: My first pointer is 0xc00009e240
}

// func main() {
// 	var myInt int
// 	myInt = 10

// 	fmt.Println(myInt)
//    }
