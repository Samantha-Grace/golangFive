package main

import "fmt"

//maps are keys and values - look up key to get value
// use make keyworkd
func main() {
	// created using make map, he key will be string (how you look things up) and it will stor int
	// map is a reference type
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	//maps are not sorted,  you cannot iterate over your map (below always returns random order)
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//now lets delete
	delete(intMap, "two")

	// test to see if something exists in a map
	el, ok := intMap["two"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}
}
