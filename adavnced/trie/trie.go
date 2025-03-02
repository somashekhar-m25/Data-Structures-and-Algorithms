package main

import "fmt"

const alphabetSize = 26

type node struct {
	child [alphabetSize]*node
	isEnd bool
}

type trie struct {
	root *node
}

func NewTrie() trie {
	return trie{
		root: &node{},
	}
}

func (t *trie) insertWord(w string) {
	root := t.root
	for i := 0; i < len(w); i++ {
		index := w[i] - 'a'
		if root.child[index] == nil {
			root.child[index] = &node{}
		}
		root = root.child[index]
	}
	root.isEnd = true
}

func (t *trie) searchWord(w string) bool {
	root := t.root
	for i := 0; i < len(w); i++ {
		index := w[i] - 'a'
		if root.child[index] == nil {
			return false
		}
		root = root.child[index]
	}
	if root.isEnd {
		return true
	}
	return false
}

func main() {
	myTrie := NewTrie()
	myTrie.insertWord("somashekhar")
	fmt.Println(myTrie.searchWord("somashekhar"))
}
