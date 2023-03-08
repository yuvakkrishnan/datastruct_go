package main

import "fmt"

type Stack struct {
	items []int
}

// Push an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop and return the element at the top of the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek and return the element at the top of the stack without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}
	return s.items[len(s.items)-1]
}

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}
	fmt.Println(stack.IsEmpty()) // true

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.IsEmpty()) // false
	fmt.Println(stack.Peek())    // 3

	fmt.Println(stack.Pop()) // 3
	fmt.Println(stack.Pop()) // 2
	fmt.Println(stack.Pop()) // 1

	fmt.Println(stack.IsEmpty()) // true
}
