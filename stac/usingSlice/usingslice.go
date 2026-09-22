package main

import (
	"fmt"
	"slices"
)

var arr = []int{}

func Push(Data int) bool {
	arr = append(arr, Data)
	return true
}

func PrintSlice(arr []int) {
	for index := len(arr) - 1; index >= 0; index-- {
		fmt.Printf("| %d |", arr[index])
		fmt.Println("\n___")
	}
}

func IsEmpty(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	return false
}
func PopInArr(arr *[]int) { //return deleted element
	if len(*arr) == 0 {
		return
	}
	*arr = slices.Delete(*arr, len(*arr)-1, len(*arr))
}

func Length(arr []int) {
	fmt.Println("Length of Stack is :", len(arr))
}

func main() {
	Push(2)
	Push(4)
	Push(6)
	PrintSlice(arr)
	fmt.Println("For POP ")
	PopInArr(&arr)
	// PopInArr(&arr)
	// PopInArr(&arr)
	PrintSlice(arr)
	Length(arr)
	IsEmpty(arr)
}
