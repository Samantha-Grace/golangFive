// // DC udem go prog lang crash course
package main

import "fmt"

// struct is an aggregate type because it holds many pieces of information

//declare a type, we'll call is Car, and it is a struct
type Car struct {
	// describe the kind of info it holds
	NumberofTires int
	Luxury        bool
	Bucketseats   bool
	Make          string
	Model         string
	Year          int
}

// func main() {
// 	// create a variable of that type
// 	var myCar Car

// 	myCar.NumberofTires = 4
// 	myCar.Luxury = false
// 	myCar.Bucketseats = true
// 	myCar.Make = "Ford"
// 	myCar.Model = "Mustang"
// 	myCar.Year = 1969

// }

// shortahnd for declaring and populating a variable
func main() {
	myCar := Car{
		NumberofTires: 4,
		Luxury:        false,
		Bucketseats:   true,
		Make:          "Ford",
		Model:         "Mustang",
		Year:          1969,
	}
	fmt.Printf("My cars is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}

// import "fmt"

// //create an array
// func main() {
// 	// create a new variable that is an array with a length of three and type string
// 	var myStrings [3]string
// 	// we can assign values to the array - arrays count from 0
// 	myStrings[0] = "cat"
// 	myStrings[1] = "dog"
// 	myStrings[2] = "bird"

// 	fmt.Println("First is:", myStrings[0])
// }

//  you will prob use slices more than arrays in go
// //create an array of ints
// func main() {
// 	// create a new variable that is an array with a length of three and type int
// 	var myStrings [3]string
// 	// we can assign values to the array - arrays count from 0
// 	myStrings[0] = 1
// 	myStrings[1] = 2
// 	myStrings[2] = 3

// 	fmt.Println("First is:", myStrings[0])
// }

// // aggregate types (assar, struct)
// // reference types (pointers, slices, maps, functions, channels)
// // interface type

// //array

// func main() {
// 	var myStrings [3]string

// 	myStrings[0] = "zero"
// 	myStrings[1] = "one"
// 	myStrings[2] = "two"

// 	fmt.Println("1st element is ", myStrings[0])
// }

// //struct

// type Car struct {
// 	NumberOfTires int
// 	Luxury        bool
// 	Bucketseats   bool
// 	Make          string
// 	Model         string
// 	Year          int
// }

// func main() {
// 	myCar := Car

// 	myCar.NumberOfTires = 4
// 	myCar.Luxury = false
// 	myCar.B

// }
