package main

import "fmt"

func main() {
	var arr = []int{}
	//push
	arr = append(arr, 4)
	arr = append(arr, 5)
	arr = append(arr, 6)
	for index, val := range arr {
		fmt.Println(index, val)
	}
	ar := arr
	ar[0] = 9
	fmt.Println(arr)
	fmt.Println(ar)
	//pop
}
