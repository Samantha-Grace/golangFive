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
}
