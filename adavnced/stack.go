package main

import "fmt"

type stack struct {
	data []int
}

func (s *stack) push(val int) {
	s.data = append(s.data, val)
}

func (s *stack) pop() int {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *stack) peak() int {
	return s.data[len(s.data)-1]
}

func main() {
	myStack := stack{}
	myStack.push(1)
	myStack.push(2)
	myStack.push(3)
	myStack.push(4)
	myStack.push(5)
	myStack.push(6)
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack.peak())
	fmt.Println(myStack.pop())
}
