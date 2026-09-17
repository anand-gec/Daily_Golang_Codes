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
func printPreOrder(root *Node)[]int {
	if root == nil {
		return nil
	}
	arr:=[]int{}
	arr=append(arr, root.data)
	
	printPreOrder(root.left)
	printPreOrder(root.right)
	return arr
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
	fmt.Printf("%d",printPreOrder(root))

}
