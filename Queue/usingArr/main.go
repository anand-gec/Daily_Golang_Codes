// By ai Queue in Slice
package main

import "fmt"

// Define the global slice and maximum size
var arr []int
var maxSize int = 5 // Set the maximum capacity of the queue

// Enqueue: Adds an element to the back of the queue
func enqueue(data int) {
	if isFull() {
		fmt.Println("Error: Queue is full! Cannot add", data)
		return
	}
	arr = append(arr, data)
	fmt.Printf("Enqueued: %d\n", data)
}

// Dequeue: Removes and returns the element from the front
func dequeue() int {
	if isEmpty() {
		fmt.Println("Error: Queue is empty! Cannot dequeue.")
		return -1 // Return a dummy error value
	}

	frontElement := arr[0] // Get the first element
	arr = arr[1:]          // Remove the first element by reslicing
	return frontElement
}

// Top (Peek): Returns the front element without removing it
func topElement() int {
	if isEmpty() {
		fmt.Println("Error: Queue is empty! No top element.")
		return -1
	}
	return arr[0]
}

// IsEmpty: Returns true if the queue has no elements
func isEmpty() bool {
	return len(arr) == 0
}

// IsFull: Returns true if the queue reaches its max capacity
func isFull() bool {
	return len(arr) >= maxSize
}

func main() {
	// 1. Try to add elements
	enqueue(10)
	enqueue(20)
	enqueue(30)

	// 2. Check the top element
	fmt.Println("Top element is:", topElement()) // Output: 10

	// 3. Dequeue elements
	fmt.Println("Dequeued:", dequeue()) // Output: 10
	fmt.Println("Dequeued:", dequeue()) // Output: 20

	// 4. Check if empty
	fmt.Println("Is queue empty?", isEmpty()) // Output: false
}
