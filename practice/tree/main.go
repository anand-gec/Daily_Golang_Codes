package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(key int) *Node {
	return &Node{data: key}
}
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(" -> ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}
func main() {
	root := NewNode(4)
	root.left = NewNode(3)
	preOrder(root)

}
