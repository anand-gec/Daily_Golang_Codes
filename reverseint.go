package main

import "fmt"

func reverseInt(num int) int {

	final := 0
	for num > 0 {
		final = (final * 10) + (num % 10)
		num /= 10
	}
	return final
}

func main() {
	fmt.Println("Reverse number is :", reverseInt(123))
}
