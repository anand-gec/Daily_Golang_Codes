package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return arr[mid]
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	arr := [5]int{2, 3, 6, 7, 9}
	target := 6

	fmt.Println(binarySearch(arr[:], target))
}
