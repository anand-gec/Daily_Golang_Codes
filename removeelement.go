package main

import "fmt"

func removeElement(num []int, val int) int {
	j := 0
	for i := 0; i < len(num); i++ {
		if num[i] != val {
			num[j] = num[i]
			j++
		}
	}
	return j
}

func main() {
	arr2 := [4]int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(arr2[:], val))

	fmt.Println("Again Call the functions ")
	//again call
	fmt.Println("")
	arr3 := []int{4, 2, 6, 4}
	val2 := 4

	k := removeElement(arr3, val2)

	// fmt.Println("New length:", k)
	fmt.Println("Modified slice:", arr3[:k])

}
