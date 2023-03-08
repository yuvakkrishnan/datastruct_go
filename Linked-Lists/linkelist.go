package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// add adds a new node with the given data to the end of the list
func (list *LinkedList) add(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

// remove removes the first node with the given data from the list
func (list *LinkedList) remove(data int) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		return
	}

	currentNode := list.head
	for currentNode.next != nil {
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

// contains returns true if the list contains a node with the given data
func (list *LinkedList) contains(data int) bool {
	currentNode := list.head

	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// print iterates through the list and prints out the data in each node
func (list *LinkedList) print() {
	currentNode := list.head

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

// toSlice returns a new slice containing the data from each node in the list, in order
func (list *LinkedList) toSlice() []int {
	result := []int{}

	currentNode := list.head

	for currentNode != nil {
		result = append(result, currentNode.data)
		currentNode = currentNode.next
	}

	return result
}

func main() {
	list := LinkedList{}

	// add elements to the list
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)

	// remove an element from the list
	list.remove(2)

	// print the list
	list.print()

	// check if the list contains an element
	fmt.Println(list.contains(3)) // true
	fmt.Println(list.contains(5)) // false

	// convert the list to a slice and print it
	slice := list.toSlice()
	fmt.Println(slice) // [1 3 4]
}
