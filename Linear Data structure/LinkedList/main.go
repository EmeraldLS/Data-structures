package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) addNode(value int) {
	newNode := node{
		value: value,
		next:  nil,
	}
	if l.head == nil {

	} else {

	}
	fmt.Println(newNode)
}

func main() {

}
