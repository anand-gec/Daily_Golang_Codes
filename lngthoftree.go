package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode acts as the constructor for a node
func NewNode(key int) *Node {
	return &Node{data: key}
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lh := maxDepth(root.left)
	rh := maxDepth(root.right)
	return 1 + max(lh, rh)
}

func main() {
	root := NewNode(1)

	root.left = NewNode(2)
	root.right = NewNode(3)

	root.left.left = NewNode(4)
	root.left.right = NewNode(5)

	root.left.right.left = NewNode(6)
	root.left.right.right = NewNode(7)

	fmt.Print("inOrder Traversal: ")
	fmt.Println(maxDepth(root))
}
