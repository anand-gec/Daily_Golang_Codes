// Online Go compiler (editor)
// Write and run Go online using this editor.

package main

import (
	"fmt"
	"slices"
)

func main() {
	//([]int,1,5) 1st to starting from 1 2nd Ending is 5 or also capacity is 5 but in slice condition it will increase the capacity of like 5 than 10 than 20 than 40 increase by multiply of 2.
	// make can contain more space rather than slice in golang language.
	// but in slice it can be contain capacity that add value in slice only plus +1

	var nums = make([]int, 1, 2)
	nums = append(nums, 2)

	nums = append(nums, 4)
	nums = append(nums, 4)
	nums = append(nums, 4)

	fmt.Println(nums)
	fmt.Println("Length of nums array", len(nums))
	fmt.Println("Capacity of nums array", cap(nums))
	fmt.Println()

	var sli []int
	sli = append(sli, 4)
	sli = append(sli, 4)
	sli = append(sli, 4)

	fmt.Println(sli)
	fmt.Println("Length of nums array", len(sli))
	fmt.Println("Capacity of nums array", cap(sli))
	fmt.Println()

	var arr2 = []int{1, 2}
	var arr1 = []int{1, 2}
	fmt.Println("Is exactly same", slices.Equal(arr1, arr2))
	
	fn:=process()
	fn(6)
}

// Functions
func process() func(a int) int {
	return func(a int) int {
		return 4
	}
}
