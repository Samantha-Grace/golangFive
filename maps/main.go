package main

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
}
