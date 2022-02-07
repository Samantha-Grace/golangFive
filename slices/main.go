package main

import "fmt"

func main() {
	var animals []string
	// append to the slice just adds to the end of the slice
	animals = append(animals, "cat")
	animals = append(animals, "dog")
	animals = append(animals, "bird")
	animals = append(animals, "fish")

	fmt.Println(animals)
}
