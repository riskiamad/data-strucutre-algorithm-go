package main

import "fmt"

// Node: represents each node in the trie
type Node struct {
	children [26]*Node // represent 26 of alphabets
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

// InitTrie : create new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert: add word to trie
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search: find word in a trie
func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()

	myTrie.Insert("oreo")
	fmt.Println(myTrie)
	fmt.Println(myTrie.Search("oreo"))
}
