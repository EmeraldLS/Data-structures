package main

import "fmt"

type Queue []string

func (q Queue) isEmpty() bool {
	return len(q) == 0
}

func (q *Queue) Push(ele string) Queue {
	*q = append(*q, ele)
	return *q
}

func (q *Queue) Unshift() (string, bool) {
	if q.isEmpty() {
		return "", false
	} else {
		index := 0
		element := (*q)[index]
		*q = (*q)[index+1:]
		return element, true
	}
}

func (q Queue) Element() {
	for _, item := range q {
		fmt.Println(item)
	}
}

func main() {
	var queue Queue
	queue.Push("Lawrence")
	queue.Push("Segun")
	queue.Push("EmeraldLS")
	fmt.Println(queue)
	queue.Element() // Print out all element of the queue
	// Removes element from the queue begiining from index[0]. First in first out(FIFO)
	queue.Unshift()
	queue.Unshift()
	queue.Unshift()

	fmt.Println("Queue is empty? - ", queue.isEmpty()) //Print a boolean value if the queue is true or false
	fmt.Println(queue)
}
