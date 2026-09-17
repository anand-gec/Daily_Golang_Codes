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

// PreOrder: Root -> Left -> Right
func printPreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	printPreOrder(root.left)
	printPreOrder(root.right)
}

// InOrder: Left -> Root -> Right
func printInOrder(root *Node) {
	if root == nil {
		return
	}
	printInOrder(root.left)
	fmt.Printf("%d ", root.data)
	printInOrder(root.right)
}

// PostOrder: Left -> Right -> Root
func printPostOrder(root *Node) {
	if root == nil {
		return
	}
	printPostOrder(root.left)
	printPostOrder(root.right)
	fmt.Printf("%d ", root.data)
}

func main() {
	root := NewNode(1)
	
	root.left = NewNode(2)
	root.right = NewNode(3)
	
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	
	root.right.left = NewNode(6)
	root.right.right = NewNode(7)


	fmt.Print("PreOrder Traversal:  ")
	printPreOrder(root)

	fmt.Print("\nIOorder Traversal:   ")
	printInOrder(root)

	fmt.Print("\nPostOrder Traversal: ")
	printPostOrder(root)

}
