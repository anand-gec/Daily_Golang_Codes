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

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)

}

func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
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
	fmt.Println("PreOrder-----------")
	PreOrder(root)
	fmt.Println("\nInOrder-------------")
	InOrder(root)
	fmt.Println("\nPostOrder--------------")
	PostOrder(root)
}
