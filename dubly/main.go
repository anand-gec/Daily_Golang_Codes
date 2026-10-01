package main

import (
	"fmt"
	"os"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

type Doubly struct {
	len  int
	head *Node
	tail *Node
}

func (linked *Doubly) InsertAtLast(num int) {
	var temp = &Node{
		val:  num,
		prev: nil,
		next: nil,
	}
	if linked.head == nil && linked.tail == nil {
		linked.head = temp
		linked.tail = temp
	} else {
		temp.prev = linked.tail
		linked.tail.next = temp
		linked.tail = temp
	}
	linked.len++
}

func (linked *Doubly) Display() {
	var p *Node = linked.head
	for p != nil {
		fmt.Printf("| %d |\n", p.val)
		fmt.Println("____")
		p = p.next
	}
}

func (linked *Doubly) DisplayReverse() {
	var p *Node = linked.tail
	for p != nil {
		fmt.Printf("| %d |\n", p.val)
		fmt.Println("____")
		p=p.prev
	}
}

func main() {
	var linked = &Doubly{}
	var choice int
	for {
		fmt.Println("Enter Your Choice...")
		fmt.Println("1. Insert node at End")
		fmt.Println("2. Traverse/Display")
		fmt.Println("3. Traverse Display in Reverse order...")
		fmt.Println("4. Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var data int
			fmt.Println("Enter Number...")
			fmt.Scan(&data)
			linked.InsertAtLast(data)
		case 2:
			linked.Display()
		case 3:
			linked.DisplayReverse()
		default:
			os.Exit(0)

		}
	}
}
