package main

import "fmt"

func mergeShort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var result []int
	i, j := 0, 0
	mid := len(arr) / 2
	arr2 := mergeShort(arr[:mid])
	arr3 := mergeShort(arr[mid:])
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
	arr := [...]int{7, 2, 6, 3, 9, 1}
	fmt.Println(mergeShort(arr[:]))

}
