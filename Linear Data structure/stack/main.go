package main

import "fmt"

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1 //Getting the last element that was added.
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack
	stack.Push("Lawrence Segun")
	stack.Push("Sanni Abdullah")
	stack.Push("EmeraldLS")

	fmt.Println(stack) //Print out all element of the stack
	for !stack.isEmpty() {
		x, y := stack.Pop()
		if y {
			fmt.Println(x)
		}
	}
}
