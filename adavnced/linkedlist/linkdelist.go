package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	h := l.head
	n.next = h
	l.head = n
	l.length++
}

func (l linkedList) printList() {
	h := l.head
	for l.length > 0 {
		fmt.Printf("%d ", h.data)
		h = h.next
		l.length--
	}
	fmt.Println("")
}

func (l *linkedList) deleteData(val int) {
	if l.head == nil {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}
	h := l.head
	for h.next.data != val {
		if h.next.next == nil {
			return
		}
		h = h.next
	}
	h.next = h.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	myList.prepend(&node{data: 1})
	myList.prepend(&node{data: 2})
	myList.prepend(&node{data: 3})
	myList.prepend(&node{data: 4})
	myList.prepend(&node{data: 5})
	myList.prepend(&node{data: 6})
	myList.printList()
	myList.deleteData(9)
	myList.printList()
	emptyList := linkedList{}
	emptyList.deleteData(1)
}
