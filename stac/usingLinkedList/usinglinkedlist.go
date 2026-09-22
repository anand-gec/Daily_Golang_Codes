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

func LengthOfStack(head *Stack) int {
	var sum = 0
	for head != nil {
		sum++ // sum = sum + i //sum++ /sum+=1
		head = head.next
	}
	return sum
}

func PrintStack(head *Stack) {
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
	IsEmpty(head)
	PrintStack(head)
	fmt.Println("Length of Stack is :", LengthOfStack(head))
}
