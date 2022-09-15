package main

import "fmt"

var count int

// Node: represents a component of binary tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert: add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search : check if tree contain key
func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}

	fmt.Println(tree)
	tree.Insert(54)
	tree.Insert(120)
	tree.Insert(158)
	tree.Insert(123)
	tree.Insert(72)

	fmt.Println(tree.Search(72))
	fmt.Printf("Found at deep level : %d\n", count)
}
