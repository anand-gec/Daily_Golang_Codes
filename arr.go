package main

import "fmt"

func main() {
	arr := [...]int{3, 5, 7}
	arr[1] = 6
	arr[2] = 9
	fmt.Println(arr[2:])
}
