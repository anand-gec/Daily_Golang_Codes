// Queue implemented in LinkedList
package main

import "fmt"

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
	var head = &Queue{}
	head.Enqueue(23)
	head.Enqueue(25)
	head.Enqueue(26)
	head.Dequeue()
	head.Starting()
	head.Ending()
	head.Size()
	head.Display()
}
