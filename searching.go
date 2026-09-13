package main

import "fmt"

func search(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i + 1
		}
	}
	return -1
}

func main() {
	arr := [...]int{2, 3, 5, 7, 9}
	target := 3
	fmt.Printf("if %d is exist than array Position is: %d", target, search(arr[:], target))
}
