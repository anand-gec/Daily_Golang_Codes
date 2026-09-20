package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

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
		fmt.Print(" -> ", head.val)
		head = head.next
	}
}
func deleteAtFirst(head *node) *node {
	if head == nil || head.next == nil {
		return nil
	}
	return head.next
}
func deleteAtLast(head *node) *node {
	if head == nil || head.next == nil {
		return nil
	}
	t := head
	if t.next.next != nil {
		t = t.next
	}
	t.next = nil
	return head
}
func insertAtFirst(val int, head *node) *node {
	var temp node
	temp.val = val
	temp.next = head
	return &temp
}
func insertAtLast(val int, head *node) *node {
	var temp = &node{
		val:  val,
		next: nil,
	}
	if head == nil {
		return temp
	}
	t := head
	if t.next != nil {
		t = t.next
	}
	t.next = temp
	return head

}
func reverseNode(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var privious *node
	current := head
	for current != nil {
		temp := current.next
		current.next = privious
		privious = current
		current = temp
	}
	return privious
}
func main() {
	var head *node
	head = addNode(3, head)
	head = addNode(4, head)
	head = addNode(5, head)
	head = deleteAtFirst(head)
	head = insertAtFirst(6, head)
	// printNode(head)
	// head=reverseNode(head)
	printNode(head)

}
