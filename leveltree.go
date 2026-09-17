package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func levelOrder(root *Node) [][]int {
	// maps to List<List<Integer>> wrapList
	var wrapList [][]int
	if root == nil {
		return wrapList
	}

	// Use a slice to simulate a Queue
	queue := []*Node{root}

	for len(queue) > 0 {
		levelNum := len(queue)
		var subList []int

		for i := 0; i < levelNum; i++ {
			// queue.peek() and queue.poll() combined:
			curr := queue[0]
			queue = queue[1:]

			if curr.left != nil {
				queue = append(queue, curr.left)
			}
			if curr.right != nil {
				queue = append(queue, curr.right)
			}

			subList = append(subList, curr.data)
		}
		wrapList = append(wrapList, subList)
	}

	return wrapList
}

func main() {
	// Using the 1 to 7 Complete Binary Tree from the previous step
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	fmt.Println("Level Order Traversal:", levelOrder(root))
}
