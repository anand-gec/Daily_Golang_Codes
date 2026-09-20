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
	printPreOrder(root.left)
	fmt.Printf("%d ", root.data)
	printPreOrder(root.right)
}

// inOrder Traversal (Iterative): Left -> Root -> Right
func PrintInOrderTraversal(root *Node) []int {
	inOrder := []int{}
	stack := []*Node{}
	node := root

	for {
		if node != nil {
			stack = append(stack, node)
			node = node.left
		} else {
			if len(stack) == 0 {
				break
			}
			// Pop from the stack
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Process current node
			inOrder = append(inOrder, node.data)

			// Move to right subtree
			node = node.right
		}
	}

	return inOrder
}

func main() {
	root := NewNode(1)

	root.left = NewNode(2)
	root.right = NewNode(3)

	root.left.left = NewNode(4)
	root.left.right = NewNode(5)

	root.right.left = NewNode(6)
	root.right.right = NewNode(7)

	fmt.Print("InOrder Traversal:  ")
	printPreOrder(root)
	fmt.Println(" ")
	fmt.Println("\nInOrder Iterative level ")
	fmt.Print("inOrder Traversal: ")
	fmt.Println(PrintInOrderTraversal(root))

}
