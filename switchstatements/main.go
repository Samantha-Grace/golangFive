//LinkedIn Learning "Learning Go" David Gassner
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1

	fmt.Println("day", dow)
}
