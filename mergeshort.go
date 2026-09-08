package main

import "fmt"

func mergeShort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	arr2 := mergeShort(arr[:mid])
	arr3 := mergeShort(arr[mid:])
	result := []int{}
	i, j := 0, 0
	for i < len(arr2) && j < len(arr3) {
		if arr2[i] <= arr3[j] {
			result = append(result, arr2[i])
			i++
		} else {
			result = append(result, arr3[j])
			j++
		}
	}
	result = append(result, arr2[i:]...)
	result = append(result, arr3[j:]...)

	return result

}

func main() {
	arr := []int{2, 4, 7, 3, 6, 9}
	k := mergeShort(arr[:])
	fmt.Println(mergeShort(k))
}
