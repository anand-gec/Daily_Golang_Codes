package main

import (
	"fmt"
)

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

// Print all nodes
func printNode(head *node) {
	for head != nil {
		fmt.Printf("-> %d", head.val)
		head = head.next
	}
}

// insert at first position
func insertFirst(val int, head *node) *node {
	var temp node
	temp.val = val
	temp.next = head
	return &temp
	// var temp = &node{
	// 	val:  val,
	// 	next: head,
	// }
	// if head == nil {
	// 	return temp
	// }
	// temp.next = head
	// head = temp
	// return temp
}

// insert at last position
func insertLast(val int, head *node) *node {
	var temp = &node{
		val:  val,
		next: nil,
	}
	if head == nil {
		return temp
	}
	t := head
	for t.next != nil {
		t = t.next
	}
	t.next = temp
	return head
}

// delete at last
func deleteAtLast(head *node) *node {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return nil
	}
	t := head
	for t.next.next != nil {
		t = t.next
	}
	t.next = nil
	return head
}

// delete at first
func deleteAtFirst(head *node) *node {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return nil
	}
	return head.next
}

// Reverse LinkedList
func reverseList(head *node) *node {
	fmt.Println("Reverse linkedList")
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

func removeNum(val int, head *node) *node {
	// var temp *node
	dummy := &node{next: head}
	temp := dummy

	for temp.next != nil {
		if temp.next.val == val {
			temp.next = temp.next.next
		} else {
			temp = temp.next
		}
	}

	return dummy.next
}

func main() {
	var head *node
	head = addNode(7, head)
	head = addNode(5, head)
	head = addNode(8, head)
	head = addNode(9, head)
	head = addNode(5, head)
	head = insertFirst(12, head)
	head = insertLast(15, head)
	head = deleteAtFirst(head)
	head = deleteAtLast(head)
	head = reverseList(head)
	head = removeNum(5, head)
	printNode(head)
	// printNode(reverseList(head))

}
