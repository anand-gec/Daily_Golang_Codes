package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode acts as the constructor for a root
func NewNode(key int) *Node {
	return &Node{data: key}
}

func isValidBST(root *Node) bool {
	var validate func(root *Node, minVal, maxVal int64) bool
	validate = func(root *Node, minVal, maxVal int64) bool {
		if root == nil {
			return true
		}
		if int64(root.data) <= minVal || int64(root.data) >= maxVal {
			return false
		}
		return validate(root.left, minVal, int64(root.data)) &&
			validate(root.right, int64(root.data), maxVal)
	}

	return validate(root, math.MinInt64, math.MaxInt64)
}

func main() {
	root3 := NewNode(8)
	root3.left = NewNode(6)
	root3.left.left = NewNode(4)
	root3.left.right = NewNode(7)
	root3.left.right.left = NewNode(5)

	root3.right = NewNode(12)
	root3.right.left = NewNode(10)
	root3.right.right = NewNode(14)
	root3.right.right.left = NewNode(13)

	root4 := NewNode(13)
	root4.right = NewNode(15)
	root4.right.right = NewNode(17)
	root4.right.right.left = NewNode(16)

	root4.left = NewNode(10)
	root4.left.left = NewNode(7)
	root4.left.right = NewNode(12)
	root4.left.left.right = NewNode(9)
	root4.left.left.right.left = NewNode(8)

	fmt.Println("root3 IS Valid BST or not ", isValidBST(root3))
	fmt.Println("root4 IS Valid BST or not ", isValidBST(root4))
}
