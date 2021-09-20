// DC udem go prog lang crash course
package main

// aggregate types (assar, struct)
// reference types (pointers, slices, maps, functions, channels)
// interface type

//array

// func main() {
// 	var myStrings [3]string

// 	myStrings[0] = "zero"
// 	myStrings[1] = "one"
// 	myStrings[2] = "two"

// 	fmt.Println("1st element is ", myStrings[0])
// }

//struct

type Car struct {
	NumberOfTires int
	Luxury        bool
	Bucketseats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myCar := Car

	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.B
}
