package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Enter year: \n")
	fmt.Scanf("%d", &n)

	if n%4 == 0 {
		if n%100 == 0 {
			if n%400 == 0 {
				fmt.Printf("%d is a leap year", n)
			} else {
				fmt.Printf("%d is not a leap year", n)
			}
		} else {
			fmt.Printf("%d is a leap year", n)
		}
	} else {
		fmt.Printf("%d is not a leap year", n)
	}
}
