package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	// append to the slice just adds to the end of the slice
	animals = append(animals, "cat")
	animals = append(animals, "dog")
	animals = append(animals, "bird")
	animals = append(animals, "fish")

	fmt.Println(animals)

	// another way to iterate over referenced type like slices and maps.  Use the FOR LOOP

	// for _, x := range animals {
	// 	fmt.Println(x)

	// prints out:
	// [cat dog bird fish]
	// cat
	// dog
	// bird
	// fish

	for i, x := range animals {
		fmt.Println(i, x)

		// now starts counting at zero just like arrays
		// [cat dog bird fish]
		// 0 cat
		// 1 dog
		// 2 bird
		// 3 fish
	}

	// print single element from slice
	fmt.Println("element zero is", animals[0])
	// Prints out: element zero is cat

	// print first two elements from that position
	fmt.Println("First two elements are", animals[0:2])
	// Prints out: First two elements are [cat dog]

	fmt.Println("The slice is", len(animals), "elements long")
	// Prints out: The slice is 4 elements long

	fmt.Println("Is the slice empty?", animals == nil)
	// Prints out: Is the slice empty? false

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	// Prints out: Is it sorted? false

	//  Now sort a slice of strings
	sort.Strings(animals)
	fmt.Println("Is it sorted NOW?", sort.StringsAreSorted(animals))
	fmt.Println(animals)
	// Prints out: Is it sorted NOW? true [bird cat dog fish]

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
	// Prints out: [bird cat fish]
}

//How to get rid of a single item in the slice
// add parameters, slice we will call 'a', it is a slice of strong, and index to delete we will call 'i' an int
//
func DeleteFromSlice(a []string, i int) []string {
	// delete an element from a slice, by an index.  Remove element from index i from a
	a[i] = a[len(a)-1]
	// this copies the last element in slice to the index that we are passing.
	// then we erase the last element - using below syntax
	// the length of slice animals, minus one, is equal to nothing
	a[len(a)-1] = ""
	// truncate the slice by deleteing that last element
	a = a[:len(a)-1]
	// give return type in the top part of the func
	return a
}
