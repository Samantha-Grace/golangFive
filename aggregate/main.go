// DC udem go prog lang crash course
package main

import "fmt"

// aggregate types (assar, struct)
// reference types (pointers, slices, maps, functions, channels)
// interface type

func main() {
	var myStrings [3]string

	myStrings[0] = "zero"
	myStrings[1] = "one"
	myStrings[2] = "two"

	fmt.Println("1st element is ", myStrings[0])
}
