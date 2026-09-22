package main

import (
	"fmt"
	"os"
)

type stack struct {
	val  int
	next *stack
}

func (node *stack) push(num int) *stack {
	temp := &stack{val: num, next: node}
	return temp
}

func (node *stack) pop() *stack {
	if node == nil {
		fmt.Println("Stack is empty")
		return nil
	}
	fmt.Printf("Popped: %d\n", node.val)
	return node.next
}

func (node *stack) peek() {
	if node == nil {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println(node.val)
}

func (node *stack) display() {
	if node == nil {
		fmt.Println("Stack is empty")
		return
	}
	p := node
	for p != nil {
		fmt.Printf("| %d |\n", p.val)
		fmt.Println("____")
		p = p.next
	}
}

func main() {
	var head *stack // Starts empty (nil)
	var choice string
	for {
		fmt.Println("\nEnter Your Choice")
		fmt.Println("1. PUSH value in Stack")
		fmt.Println("2. POP value from Stack")
		fmt.Println("3. PEEK value from Stack")
		fmt.Println("4. Display Stack")
		fmt.Println("5. Exit")
		fmt.Scan(&choice) // Read user choice

		switch choice {
		case "1":
			var data int
			// var data string
			fmt.Println("Enter value to push")
			fmt.Scan(&data)
			// num, _ := strconv.Atoi(data) //contain only int type data if give any type of data than it will contain only 0.
			// head = head.push(num)
			head = head.push(data) //contain only int type data
		case "2":
			head = head.pop()
		case "3":
			head.peek()
		case "4":
			head.display()
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
