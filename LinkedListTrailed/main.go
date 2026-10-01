package main

import "fmt"

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
	fmt.Println("")
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
	if head == nil || head.val == num && head.next == nil {
		return nil
	}
	t := head
	for t != nil && t.next != nil {
		if t.next.val == num {
			x := t.next.next
			// t.next = nil
			t.next = x
		}
		t = t.next
	}
	// for t.next.val != num {
	// 	t = t.next
	// }
	// t.next = nil
	return head
}
func ReverseList(head *node) *node {
	var reverse *node
	t := head
	for t != nil {
		temp := t.next
		t.next = reverse
		reverse = t
		t = temp
	}
	return reverse
}

func main() {
	var head *node
	head = newNode(16, head)
	head = newNode(12, head)
	head = newNode(8, head)
	head = newNode(6, head)
	head = newNode(4, head)

	// insertAtParticular(4, 3, head)
	// insertAtParticular(6, 5, head)
	head = addLastNode(20, head)
	// head = DeleteParticularNode(4, head)
	PrintNode(head)
	// fmt.Println("\n", IsValuePresent(0, head))
	fmt.Println("reversed")
	head = ReverseList(head)
	PrintNode(head)
}
