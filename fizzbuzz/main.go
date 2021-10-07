package main

import "fmt"

//fizzbuzz
//divisible by 3 -> fizz
//divisible by 5 -> buzz
//divisable by both -> fizzbuzz

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println(i)

		//check for fizz
		if i%3 == 0 {
			fmt.Printf("Fizz")
		}
		//check for buzz
		if i%5 == 0 {
			fmt.Printf("Buzz")
		}
	}

}
