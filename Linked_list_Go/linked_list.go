package linked_list

import (
	"fmt"
)

// node contains data and pointer to next data
// this is the mail building block
type Node struct {
	// can be any data type
	data int
	// form a sequential chain
	next *Node
}

// container to manage organization of the data
type LinkedList struct {
	// entry point to the list allowing access
	// first data
	head *Node
	// how many elements we have
	length int
}

// add new data to the list
func (l *LinkedList) Add(n *Node) {
	// if linked list is empty, create a new node and set the data as head
	secondData := l.head
	l.head = n
	l.head.next = secondData
	l.length++
}

func Linked() {
	newList := LinkedList{}
	values := []int{10, 20, 30}
	for _, v := range values {
		data := &Node{data: v, next: nil}
		newList.Add(data)
	}
	fmt.Println(newList)
}

// Package linked_list implements a singly linked list data structure.
//
// The Node struct represents an element in the linked list, containing data and a pointer to the next node.
//
// The LinkedList struct provides methods to manage the linked list, including adding new elements and tracking the length of the list.
//
// Add method inserts a new element with the specified data into the linked list.
//
// Example usage:
//  func main() {
//      list := linked_list.LinkedList{}
//      list.Add(10)
//  }
