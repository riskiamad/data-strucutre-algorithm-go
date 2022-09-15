package main

import "fmt"

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push : add value to stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop : remove value from stack
func (s *Stack) Pop(i int) {
	s.items = s.items[:len(s.items)-1]
}

func main() {
	myStack := &Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(30)
	myStack.Push(60)
	fmt.Println(myStack)

}
