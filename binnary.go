package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) bool {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// func mergeShort(arr []int) []int {
// 	var arr3 = []int{}
// 	start := 0
// 	end := len(arr) - 1
// 	mid := (start + end) / 2
// 	for

// 	return arr3
// }

func main() {
	arr := [...]int{2, 3, 6, 7, 9}
	target := 7
	fmt.Println("BinarySearch")
	fmt.Println("index of arr is : arr", binarySearch(arr[:], target))

	// arr2 := [...]int{6, 8, 2, 5, 3, 9, 1}
	// fmt.Println("MergeShort")
	// fmt.Println(mergeShort(arr2[:]))
}
