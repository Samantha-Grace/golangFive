package main

import (
	"container/list"
	"fmt"
)

// define linked list
func two() {
	fmt.Println("This is a linked list")
	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushFront(2)
	// we now have a linked list with '1' at the back of the list
	// and '2' at the front of the list.
}

// import (
// 	"container/list"
// 	"fmt"
// )

// func one() {
// 	l := list.New[int]()

// 	l.PushFront(10)
// 	l.PushFront(20)
// 	l.PushFront(40)

// 	fmt.Println(l)
// }

// func two() {
// 	l := list.New[string]()

// 	l.PushFront("hello")

// 	fmt.Println(l)
// }

// func three() {
// 	l := list.New[int]()

// 	l.PushFront(10)
// 	l.PushFront(20)
// 	l.PushFront(40)
// 	l.PushFront(42)

// 	fmt.Println(l)

// 	s := l.Filter(func(i int) bool {
// 		return i%10 == 0
// 	})

// 	fmt.Println(s)
// }
