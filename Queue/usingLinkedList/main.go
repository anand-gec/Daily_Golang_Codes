// By ai Queue in LinkedList
package main

import "fmt"

// Node represents a single element in the linked list
type Node struct {
	data int
	next *Node
}

// Global pointers to keep track of the queue boundaries
var head *Node // Points to the FRONT of the queue
var tail *Node // Points to the BACK of the queue

// Enqueue: Adds an element to the back (tail) of the queue
func enqueue(value int) {
	if isFull() {
		fmt.Println("Queue is full!")
		return
	}

	// Create a new node
	newNode := &Node{data: value, next: nil}

	// If the queue is empty, the new node becomes both head and tail
	if head == nil {
		head = newNode
		tail = newNode
	} else {
		// Link the current tail to the new node, then move the tail pointer
		tail.next = newNode
		tail = newNode
	}
	fmt.Printf("Enqueued: %d\n", value)
}

// Dequeue: Removes and returns the element from the front (head)
func dequeue() int {
	if isEmpty() {
		fmt.Println("Error: Queue is empty! Cannot dequeue.")
		return -1
	}

	// Grab the data from the front node
	frontValue := head.data

	// Move the head pointer to the next node in line
	head = head.next

	// If the queue becomes empty after moving head, reset tail to nil too
	if head == nil {
		tail = nil
	}

	return frontValue
}

// Top Element (Peek): Returns the front value without removing it
func topElement() int {
	if isEmpty() {
		fmt.Println("Error: Queue is empty! No top element.")
		return -1
	}
	return head.data
}

// IsEmpty: Returns true if the head pointer is nil
func isEmpty() bool {
	return head == nil
}

// IsFull: Linked lists dynamically grow, so they are practically never full
func isFull() bool {
	return false
}

func main() {
	// 1. Add elements (Enqueue)
	enqueue(10)
	enqueue(20)
	enqueue(30)

	// 2. View the front element
	fmt.Println("Top element is:", topElement()) // Output: 10

	// 3. Remove elements (Dequeue)
	fmt.Println("Dequeued:", dequeue()) // Output: 10
	fmt.Println("Dequeued:", dequeue()) // Output: 20

	// 4. Check if empty
	fmt.Println("Is queue empty?", isEmpty()) // Output: false
}

//Doubly linkedList
/*
package main

import "fmt"

// Node represents a single element with two directions
type Node struct {
	data int
	next *Node // Points to the node ahead of it
	prev *Node // Points to the node behind it
}

// Global pointers for the queue boundaries
var head *Node // Front of the queue (where we Dequeue)
var tail *Node // Back of the queue (where we Enqueue)

// Enqueue: Adds an element to the back (tail) of the queue
func enqueue(value int) {
	if isFull() {
		fmt.Println("Queue is full!")
		return
	}

	// Create a new node
	newNode := &Node{data: value, next: nil, prev: nil}

	// If the queue is empty, the new node is both head and tail
	if head == nil {
		head = newNode
		tail = newNode
	} else {
		// Link the current tail forward to the new node
		tail.next = newNode
		// Link the new node backward to the current tail
		newNode.prev = tail
		// Move the tail pointer to the new node
		tail = newNode
	}
	fmt.Printf("Enqueued: %d\n", value)
}

// Dequeue: Removes and returns the element from the front (head)
func dequeue() int {
	if isEmpty() {
		fmt.Println("Error: Queue is empty! Cannot dequeue.")
		return -1
	}

	// Grab the data from the front node
	frontValue := head.data

	// Move the head pointer forward to the next node
	head = head.next

	// If the queue is now empty, reset tail to nil
	if head == nil {
		tail = nil
	} else {
		// Cut off the link to the old head node
		head.prev = nil
	}

	return frontValue
}

// Top Element (Peek): Returns the front value without removing it
func topElement() int {
	if isEmpty() {
		fmt.Println("Error: Queue is empty! No top element.")
		return -1
	}
	return head.data
}

// IsEmpty: Returns true if the head pointer is nil
func isEmpty() bool {
	return head == nil
}

// IsFull: Doubly linked lists grow dynamically, so they are never full
func isFull() bool {
	return false
}

func main() {
	// 1. Add elements (Enqueue)
	enqueue(10)
	enqueue(20)
	enqueue(30)

	// 2. View the front element
	fmt.Println("Top element is:", topElement()) // Output: 10

	// 3. Remove elements (Dequeue)
	fmt.Println("Dequeued:", dequeue()) // Output: 10
	fmt.Println("Dequeued:", dequeue()) // Output: 20

	// 4. Check if empty
	fmt.Println("Is queue empty?", isEmpty()) // Output: false
}

*/
