// Largest Array return using Functions..
package main

import "fmt"

func largeArr(arr []int) int {
	var large = arr[0]
	for i := 0; i < len(arr); i++ {
		if large < arr[i] {
			large = arr[i]
		}
	}
	return large
}
func main() {
	var arr = [...]int{5, 8, 9, -5, 2, -4, 9}
	
	 fmt.Println("largest array is:", largeArr(arr[:]))
}
