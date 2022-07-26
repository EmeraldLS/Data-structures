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

func (l *linkedList) AddNodeToBeginningOfThelist(value int) {
	newNode := &node{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = newNode //If there's no node yet, it's will create one.
		l.length++       // Increment the length of the list.
	} else {
		newNode.next = l.head
		l.head = newNode
		l.length++
	}
}

func (l *linkedList) deleteNodeWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.value != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l linkedList) Printlist() {
	if l.head == nil {
		return //If there are no nodes, it should not print anything.
	} else {
		currentNode := l.head    //If there's are nodes, it will set the current node to the first node in the list.
		for currentNode != nil { // Iterate through the list until the current node is pointing to nil.
			listoutput := fmt.Sprintf("%v -> ", currentNode.value)
			currentNode = currentNode.next
			fmt.Print(listoutput)
		}
		fmt.Print(nil, "\n")
	}
}

func main() {
	list := linkedList{}
	list.AddNodeToBeginningOfThelist(5)
	fmt.Println(list.head) //The output for this will be nil because it has no node it's pointing to.
	list.AddNodeToBeginningOfThelist(6)
	list.AddNodeToBeginningOfThelist(7)
	list.AddNodeToBeginningOfThelist(8)
	fmt.Println(list.head) // The output for this will be the address/pointer(memory location) of node with value 7
	list.AddNodeToBeginningOfThelist(9)
	list.AddNodeToBeginningOfThelist(10)

	list.Printlist()
	fmt.Println(list.head) // the output for this will eb address/ pointer(memory location) of 9
}
