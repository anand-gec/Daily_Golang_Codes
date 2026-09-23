// array is pass by value arr to duplicate new arrayNAme
// But slice is pass by reference to point slice<-newSlice name
package main

import "fmt"

func main() {
	arr := [...]int{3, 5, 7}
	arr[1] = 6
	arr[2] = 9
	fmt.Println(arr[2:])
	arr2 := arr
	arr[0] = 2
	fmt.Println(arr2)
	fmt.Println(arr)
	array := []int{4, 6, 2, 9}
	array[0] = 1
	array[1] = 2
	array2 := array
	array[0] = 10
	fmt.Println(array)
	fmt.Println(array2)

	var arrays = [5]int{2, 4}
	fmt.Println(arrays)
}
