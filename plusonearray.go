package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	arr := []int{1, 3, 9}
	arr2 := []int{2, 5, 6}
	arr3 := []int{9, 9}
	i := plusOne(arr3)
	k := plusOne(arr2)
	m := plusOne(arr)
	fmt.Println(m)
	fmt.Println(k)
	fmt.Println(i)
}
