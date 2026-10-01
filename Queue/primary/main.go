// Queue implemented in LinkedList
package main

import (
	"fmt"
	"os"
)

type Node struct {
	val  int
	next *Node
}
type Queue struct {
	len   int
	Front *Node
	Rear  *Node
}

func (head *Queue) Enqueue(num int) {
	var temp = &Node{val: num, next: nil}
	if head.Front == nil && head.Rear == nil {
		head.Front = temp
		head.Rear = temp
	} else {
		head.Rear.next = temp
		head.Rear = temp
	}
	head.len++
}

func (head *Queue) Dequeue() int {
	var value = 0
	if head.Rear == head.Front {
		value = head.Front.val
		head.Rear = nil
		head.Front = nil
	} else {
		value = head.Front.val
		head.Front = head.Front.next
	}
	head.len--
	return value
}

func (head *Queue) Starting() {
	fmt.Println("Front Node is : ", head.Front.val)
}

func (head *Queue) Ending() {
	fmt.Println("Last Node is : ", head.Rear.val)
}

func (head *Queue) Size() {
	fmt.Println("Size of Queue is : ", head.len)
}

func (head *Queue) Display() {
	var temp *Node = head.Front
	for temp != nil {
		fmt.Print(" -> ", temp.val)
		temp = temp.next
	}
	fmt.Println("")
}

func main() {
	var myQueue = &Queue{}
	var choice string
	for {
		fmt.Println("Welcome to Queue..")
		fmt.Println("Enter Choice")
		fmt.Println("1. Enqueue")
		fmt.Println("2. Dequeue")
		fmt.Println("3. Starting")
		fmt.Println("4. Ending")
		fmt.Println("5. Size")
		fmt.Println("6. Traverse/Display")
		fmt.Println("7. Exit")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			var num int
			fmt.Println("Enter number for Node in Queue.")
			fmt.Scan(&num)
			myQueue.Enqueue(num)
		case "2":
			fmt.Println("Delete Value is :", myQueue.Dequeue())
		case "3":
			myQueue.Starting()
		case "4":
			myQueue.Ending()
		case "5":
			myQueue.Size()
		case "6":
			myQueue.Display()
		default:
			os.Exit(0)
		}
	}
}
