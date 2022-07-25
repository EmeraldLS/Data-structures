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

func (l *linkedList) AddNode(value int) {
	newNode := &node{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		l.length++
	} else {
		l.head = newNode
		newNode.next = l.head
		l.length++
	}
}

func (l linkedList) Printlist() {
	if l.head == nil {
		return
	} else {
		currentNode := l.head
		for currentNode != nil {
			fmt.Println(currentNode.value)
			currentNode = currentNode.next
		}
	}

}

func main() {
	list := linkedList{}
	list.AddNode(5)
	list.AddNode(6)
	list.AddNode(8)
	list.Printlist()
}
