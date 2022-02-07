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

	// pass in a reference to x
	changeValueOfPointer(&x)

	fmt.Println("After func call, X is now", x) // prints out: X is now 25

}

// create a new function that takes a parameter names num that is a pointer of type int and doesn't return anything
func changeValueOfPointer(num *int) {
	// change the value of num to 20
	*num = 25
}

// when you want a reference to a pointer, you prepend  the ampersand (&)
// when you want to get the value of a pointer and change it, you prepend the pointer with an asterisk (*)
// aka when you want to dereference a pointer, you prepend the asterisk (*)
