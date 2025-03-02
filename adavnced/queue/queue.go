package main

import "fmt"

type queue struct {
	data []int
}

func (q *queue) enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *queue) dequeue() int {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *queue) peak() int {
	return q.data[0]
}

func main() {
	myQueue := queue{}
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(3)
	myQueue.enqueue(4)
	myQueue.enqueue(5)
	fmt.Println(myQueue)
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.peak())
	fmt.Println(myQueue.dequeue())
}
