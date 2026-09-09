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
		fmt.Printf("-> %d", head.val)
		head = head.next
	}
}
func insertFirst(val int, head *node) *node {
	var tempNode node
	tempNode.val = val
	tempNode.next = head
	return &tempNode

}
func insertLast(val int, head *node) *node {
	var temp = &node{
		val:  val,
		next: nil,
	}
	if head == nil {
		head = temp
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = temp
		temp.next = nil
	}

	return head
}

func reverseList(head *node) *node {
	if head == nil {
		return head
	}
	var previous *node
	current := head
	for current != nil {
		temp := current.next
		current.next = previous
		previous = current
		current = temp
	}
	return previous
}

func main() {
	var head *node
	head = addNode(7, head)
	head = addNode(8, head)
	head = addNode(9, head)
	head = insertFirst(12, head)
	// head = insertLast(15, head)
	printNode(head)
	fmt.Println("\nReverse LinkedList")
	printNode(reverseList(head))

}
