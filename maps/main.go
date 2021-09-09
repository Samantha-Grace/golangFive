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

	dict := map[string]string{
		"good":    "kotu",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	dict["up"] = "yukari"
	dict["down"] = "asagi"
	dict["good"] = "iyi"
	dict["mistake"] = ""

	// copied := map[string]string{"down":"asagi", "good":"iyi", "great":"harika", "mistake":"", "perfect":"mukemmel", "up":"yukari"}
	// "perfect" means "mukemmel"

	// fmt.Printf("%#v\n", dict)
	// for k, v := range dict {
	// 	fmt.Printf("%q means %#v\n", k, v)
	// }

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.\n", query)

	value := dict[query]
	fmt.Printf("%q means %#v\n)", query, value)
	fmt.Printf("# of keys: %d\n", len(dict))
}
