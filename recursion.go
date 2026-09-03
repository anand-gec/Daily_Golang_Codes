package main

import "fmt"

func sum(n int) int {
	if n <= 1 {
		return 1
	}
	return n + sum(n-1)
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func inc(n int) {
	if n <= 5 {
		fmt.Print(n)
		inc(n + 1)
	} else {
		fmt.Println("")
	}
}

func dec(n int) {
	if n >= 1 {
		fmt.Print(n)
		dec(n - 1)
	}
}

func main() {
	fmt.Println(sum(5))
	fmt.Println(fact(5))
	inc(1)
	dec(5)
}
