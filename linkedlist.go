package main

import "fmt"

type node struct {
	val  int
	next *node
}

// add node
func addNode(val int, head *node) *node {
	var temp = node{
		val:  val,
		next: head,
	}
	head = &temp
	return head
}

func printNode(head *node) {
	for head != nil {
		fmt.Printf("%d -> ", head.val)
		head = head.next
	}
}
func main() {
	var head *node

	head = addNode(7, head)
	head = addNode(8, head)
	head = addNode(9, head)

	printNode(head)
}
