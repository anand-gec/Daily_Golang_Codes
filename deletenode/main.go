package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode acts as the constructor for a root
func NewNode(key int) *Node {
	return &Node{data: key}
}

func deleteNode(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if root.data == val {
		return helper(root)
	}
	dummy := root
	for root != nil {
		if root.data > val {
			if root.left != nil && root.left.data == val {
				root.left = helper(root.left)
				break
			} else {
				root = root.left
			}
		} else {
			if root.right != nil && root.right.data == val {
				root.right = helper(root.right)
				break
			} else {
				root = root.right
			}
		}
	}
	return dummy
}

func helper(root *Node) *Node {
	if root.left == nil {
		return root.right
	} else if root.right == nil {
		return root.left
	}
	var findLastRight func(root *Node) *Node
	findLastRight = func(root *Node) *Node {
		if root.right == nil {
			return root
		}
		return findLastRight(root.right)
	}
	rightChild := root.right
	lastRight := findLastRight(root.left)
	lastRight.right = rightChild
	return root.left
}

//InOrder
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
}

func main() {
	val := 10
	root3 := NewNode(8)
	root3.left = NewNode(6)
	root3.left.left = NewNode(4)
	root3.left.right = NewNode(7)
	root3.left.right.left = NewNode(5)

	root3.right = NewNode(12)
	root3.right.left = NewNode(10)
	root3.right.right = NewNode(14)
	root3.right.right.left = NewNode(13)
	fmt.Println("Before Deleting")
	InOrder(root3)

	fmt.Println("\nDelete Node ------------")
	deleteNode(root3, val)
	fmt.Println("after deletingInOrder value is ")
	InOrder(root3)
}
