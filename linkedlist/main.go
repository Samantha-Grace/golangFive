package main

import "fmt"

// first describe a node - create a struct!
// node needs a data - we will use a single number
// it also needs to hold "next": which is address of next node.  the * mean it is a pointer
type node struct {
	data int
	next *node
}

// describe the list
// needs to hold a HEAD and that has to be an address of the pointer of the node (which is the head - you don't need the whole list)
// length = how long linkedList is
type linkedList struct {
	head   *node
	length int
}

// now need to make a method - that puts node inside linkedList
// the reciever is going to be linkedList, and we'll call it addNode.  It is going to take a node to be added at the front
// this is how you define a method. (l *linkedList) is called a method reciever. No * is value reciever, then it is just a copy of it
// inside method make a place called second, put current head in second
// set new node as the head
// let the new head point to the old head (which is second)
// increade the length, because we added a node
func (l *linkedList) addNode(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// short test to add nodes to linkedList
// make a new list called myList
// create nodes
// we need to pass in a pointer &
// then prepend the list
// then print the list
func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 8}
	node3 := &node{data: 6}
	myList.addNode(node1)
	myList.addNode(node2)
	myList.addNode(node3)
	fmt.Println(myList)
}

// output: {0xc00009e240 3}
// this is the address of node 3 because we added that last - so that will be the head
// 3 is the length of list
