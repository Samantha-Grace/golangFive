package main

import (
	"doctorapp/eliza/doctor"
	"fmt"
)

func main() {
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}
