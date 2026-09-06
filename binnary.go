package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) []int {
	var arr2 = []int{}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			arr2 = append(arr2, mid)
			break
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return arr2
}

func main() {
	arr := [5]int{2, 3, 6, 7, 9}
	target := 7

	fmt.Println("index of arr is : arr", binarySearch(arr[:], target))
}
