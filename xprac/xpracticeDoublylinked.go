package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

func Append(val int, head *Node) *Node {
	temp := &Node{
		prev: nil,
		data: val,
		next: head,
	}
	if head != nil {
		head.prev = temp
	}
	return temp
}

func Display(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("| %d |\n", curr.data)
		fmt.Println("______")
		curr = curr.next
	}
}

func DeleteFirst(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := head.next
	if newHead != nil {
		newHead.prev = nil
	}
	head.next = nil
	return newHead
}

func DeleteLast(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	curr := head
	for curr.next != nil {
		curr = curr.next
	}

	curr.prev.next = nil
	curr.prev = nil
	return head
}

func main() {
	var head *Node

	head = Append(23, head)
	head = Append(25, head)
	head = Append(26, head)

	fmt.Println("Initial List:")
	Display(head)

	head = DeleteFirst(head)
	Display(head)

	fmt.Println("After DeleteLast:")
	head = DeleteLast(head)
	Display(head)
}
  