package main

import "fmt"

type Stack struct {
	Data int
	next *Stack
}

func Push(Data int, head *Stack) *Stack {
	var temp = Stack{
		Data: Data,
		next: head,
	}
	head = &temp
	return head
}

func Pop(head *Stack) *Stack {
	if head == nil || head.next == nil {
		return nil
	}
	fmt.Printf("Popped: %d\n", head.Data)
	temp := head.next
	head.next = nil
	return temp
}

func IsEmpty(head *Stack) bool {
	if head == nil {
		return true
	}
	return false
}

func Length(head *Stack) int {
	var sum = 0
	for head != nil {
		sum++ // sum = sum + i /sum++ /sum+=1
		head = head.next
	}
	return sum
}

func Display(head *Stack) {
	for head != nil {
		fmt.Printf("| %d |", head.Data)
		fmt.Println("\n____")
		head = head.next
	}
}

func main() {
	var head *Stack
	head = Push(2, head)
	head = Push(6, head)
	head = Push(4, head)
	head = Pop(head)
	fmt.Println("is Stack Empty :", IsEmpty(head))
	Display(head)
	fmt.Println("Length of Stack is :", Length(head))
}
