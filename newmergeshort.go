package main

import "fmt"

// 1. The main function to call
func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	mergeSort(left)
	mergeSort(right)

	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	result := make([]int, 0, len(arr))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	copy(arr, result)
}

func main() {
	arr := []int{2, 4, 7, 3, 6, 9}
	mergeSort(arr)
	fmt.Println(arr)
}