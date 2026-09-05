// It Reverse the String
package main

import (
	"fmt"
)

func reverse(s string) string {
	// Write code to reverse the string
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
func main() {
	word := "Vikash Kushwaha"
	fmt.Println(reverse(word))

}
