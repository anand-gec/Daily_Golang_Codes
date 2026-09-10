//Valid only 4 digit number

package main

import "fmt"

func insertIntInArray(y int) []int {

	x := 0
	for y > 0 {
		x = (x * 10) + (y % 10)
		y /= 10
	}
	var arr []int
	var z = len(string(x)) + 1
	for i := 0; i < z; i++ {
		arr = append(arr, x%10)
		x /= 10
	}
	return arr
}

func main() {

	fmt.Println(insertIntInArray(6534))
}
