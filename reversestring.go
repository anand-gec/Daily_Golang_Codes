package main

import (
	"fmt"
	"strings"
)

func stringReverse(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}

	// left := 0
	// right := len(s) - 1

	// for left < right {

	// 	temp := s[left]
	// 	s[left] = s[right]
	// 	s[right] = temp

	// 	left++
	// 	right--
	// }

}

func main() {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	stringReverse(s)
	fmt.Printf("Output: [%s]\n", strings.Join(strings.Split(string(s), ""), ","))
}
