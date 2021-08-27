package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	dict := map[string]string{}

	// english := [string]{"good", "great", "perfect"}
	// turkish := [string]{"iyi", "harika", "mükemmel"}

	value := dict[query]
	fmt.Printf("%q means %#v\n)", query, value)
	fmt.Printf("# of keys: %d\n", len(dict))
}
