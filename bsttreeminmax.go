package main

import (
	"fmt"
)

// Node represents a components of the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	panic("unimplemented")
}

// FindMinBST finds the minimum value in a Binary Search Tree
func FindMinBST(root *Node) (*Node, error) {
	if root == nil {
		return nil, fmt.Errorf("the tree is empty")
	}

	current := root
	// Loop down to find the leftmost leaf
	for current.Left != nil {
		current = current.Left
	}

	return current, nil
}
func FindMaxBST(root *Node) (*Node, error) {
	if root == nil {
		return nil, fmt.Errorf("the tree is empty")
	}
	current := root
	// Loop down to find the leftmost leaf
	for current.Right != nil {
		current = current.Right
	}

	return current, nil
}

func main() {
	// Creating a sample BST
	//        20
	//       /  \
	//      10   30
	//     /  \
	//    5   15
	root := &Node{Value: 20}
	root.Left = &Node{Value: 10}
	root.Right = &Node{Value: 30}
	root.Left.Left = &Node{Value: 5}
	root.Left.Right = &Node{Value: 15}

	minNode, err := FindMinBST(root)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Minimum value in BST: %d\n", minNode.Value)
	}

	maxNode, err := FindMaxBST(root)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Maximum value in BST: %d\n", maxNode.Value)
	}
}
