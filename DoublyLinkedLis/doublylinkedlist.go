package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

func Append(val int, head *Node) *Node {
	var temp = &Node{
		prev: nil,
		data: val,
		next: head,
	}
	if head == nil {
		head = temp
	}
	//Both are run 
	// if head != nil {
	// 	head.prev = temp
	// }
	return temp
}

func DeleteFirst(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	newHead := head.next
	newHead.prev = nil
	head.next = nil
	return newHead
}

func Display(head *Node) {
	if head == nil {
		fmt.Println("LinkedList is Empty...")
	}
	for head != nil {
		fmt.Printf("%d", head.data)
		if head.next != nil {
			fmt.Print(" <-> ")
		}
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node
	head = Append(23, head)
	head = Append(25, head)
	head = Append(26, head)
	Display(head)
	head = DeleteFirst(head)
	head = DeleteFirst(head)
	// head = DeleteFirst(head)
	Display(head)

}
