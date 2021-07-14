//LinkedIn Learning "Learning Go" David Gassner
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	//fmt.Println("day", dow)

	result := ""
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	default:
		result = "It's another day"
	}
	fmt.Println(result)
}
