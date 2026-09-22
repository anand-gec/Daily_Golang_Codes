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
func PrintTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
}
func main() {
	root := NewNode(4)
	root.left = NewNode(3)
	root.right = NewNode(2)

}
