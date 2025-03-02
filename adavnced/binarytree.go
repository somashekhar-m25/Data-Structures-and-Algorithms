package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) insert(val int) {
	if n == nil {
		n = &node{data: val}
	} else if n.data < val {
		if n.right == nil {
			n.right = &node{data: val}
		} else {
			n.right.insert(val)
		}
	} else {
		if n.left == nil {
			n.left = &node{data: val}
		} else {
			n.left.insert(val)
		}
	}
}

func (n *node) preOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.data)
	n.left.preOrder()
	n.right.preOrder()
}

func (n *node) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Printf("%d ", n.data)
	n.right.inOrder()
}

func (n *node) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d ", n.data)
}

func (n *node) search(val int) bool {
	if n == nil {
		return false
	}
	if n.data < val {
		return n.right.search(val)
	} else if n.data > val {
		return n.left.search(val)
	}
	return true
}

func main() {
	root := &node{data: 50}
	root.insert(30)
	root.insert(70)
	root.insert(20)
	root.insert(40)
	root.insert(60)
	root.insert(80)
	fmt.Println(root)
	root.preOrder()
	fmt.Println("")
	root.inOrder()
	fmt.Println("")
	root.postOrder()
	fmt.Println("")
	fmt.Println(root.search(600))
}
