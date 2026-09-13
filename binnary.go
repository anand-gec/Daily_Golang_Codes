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

func main() {
	arr := [...]int{2, 3, 6, 7, 9}
	target := 6
	fmt.Println("BinarySearch")
	fmt.Println(" if exist than it will give true otherwise false :", binarySearch(arr[:], target))
	fmt.Printf("index of arr if exist than 1-len(arr) otherwise -1 than not exist : %d position", binarySearch2(arr[:], target))

}

func binarySearch2(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid+1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
