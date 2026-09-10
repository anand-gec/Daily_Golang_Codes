package main

import "fmt"

func main() {
	name := "vikash"
	names := &name

	fmt.Println(names)         //0x378e5677c070
	fmt.Println(*names)        //vikash
	fmt.Println(&name)         //0x378e5677c070
	arr := []int{2, 4, 6, 9}
	val := &arr
	arr2 := []int{}
	arr2 = arr
	arr2 =append(arr2,10)
	// arr[len(arr)-1]=10
	fmt.Println(*val)        //[2 4 6 9 10]
	fmt.Println(arr2)        //[2 4 6 9 10]
}
