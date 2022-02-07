package main

import "fmt"

// pointer points to specific memory location
func main() {
	x := 10

	myFirstPointer := &x // & holdw the address of whrere x's value of 10 is stored in memory

	fmt.Println("x is", x)                             // prints out: x is 10
	fmt.Println("My first pointer is", myFirstPointer) // prints out: My first pointer is 0xc00009e240

	// go to the address in myFirstPointer, and change the contents of that memory location to 20
	*myFirstPointer = 20 // * dereferences the pointer and sets the value of x to 20

	fmt.Println("X is now", x) // prints out: X is now 20

}

// func main() {
// 	var myInt int
// 	myInt = 10

// 	fmt.Println(myInt)
//    }
