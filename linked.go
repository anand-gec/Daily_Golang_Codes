//inset after
// Online Go compiler (editor)
// Write and run Go online using this editor.

package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

func newNode(val int, head *node) *node {
	var temp = &node{
		val:  val,
		next: head,
	}
	head = temp
	return head
}
func addLastNode(val int, head *node) *node {
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
func PrintNode(head *node) {
	for head != nil {
		fmt.Print(" -> ", head.val)
		head = head.next
	}
}
func insertAtParticular(num int, val int, head *node) *node {
	var temp = &node{
		val:  val,
		next: head,
	}
	t := head
	for t != nil {
		if t.val == num {
			x := t.next
			t.next = temp
			temp.next = x
		}
		t = t.next
	}
	return head
}
func IsValuePresent(num int, head *node) bool {
	if head == nil {
		return false
	}
	t := head
	for t != nil {
		if t.val == num {
			return true
		}
		t = t.next
	}
	return false
}
func DeleteParticularNode(num int, head *node) *node {
	if head.val == num {
		fmt.Println(head.val)
		return head.next
	}

	// t := head
	// for t.next != nil {
	// 	if t.next.val == num {
	// 		x := t.next.next
	// 		t.next = nil
	// 		t.next = x
	// 		// t.next=t.next.next
	// 		fmt.Println("return the")
	// 	}
	// 	t = t.next
	// }
	return head.next
}

func main() {
	var head *node
	head = newNode(4, head)
	head = newNode(6, head)
	head = newNode(8, head)
	head = newNode(12, head)
	head = newNode(16, head)

	// insertAtParticular(6, 12, head)
	// insertAtParticular(4, 5, head)
	// head = addLastNode(7, head)
	DeleteParticularNode(16, head)
	PrintNode(head)
	// fmt.Println("\n", IsValuePresent(0, head))
}
