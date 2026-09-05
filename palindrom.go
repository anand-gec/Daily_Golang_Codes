package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s int) bool {
	str := strconv.Itoa(s)
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	if reversed == str {
		return true
	}
	return false
}

func main() {
	var num = "43232"
	fmt.Println("the value is Palindrome :", isPalindrome(121))
	fmt.Println("the value 234 is Palindrome:", isPalindrome(234))
	fmt.Println("the value 10 is Palindrome:", isPalindrome(10))
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(newPalindrome(32))
	fmt.Println(Palindrome(123))
	fmt.Println(Palindrome(121))
	fmt.Println("New  Palindrome")
	fmt.Println(palindromeInt(232))
	fmt.Println(palindromeInt(-232))
	fmt.Println(palindromeInt(10))
}

func newPalindrome(x int) bool {
	//only 3 digit number can solve this
	var m = string(x)
	var arr2 = ""
	for i := len(m) - 1; i >= 0; i-- {
		arr2 += string(m[i])
	}
	fmt.Println(arr2)
	if m == arr2 {
		return true
	}
	return false
}

func Palindrome(b int) bool {
	// fully correct
	if b < 0 {
		return false
	}
	num := b
	result := 0
	for num != 0 {
		r := num % 10
		num = num / 10
		result = result*10 + r
	}
	if result == b {
		return true
	}
	return false
}


func palindromeInt(num int)bool{

 var number=num
    final:=0
  for num>0{
    final=(final*10)+(num%10)
    num/=10
  }
  if number==final{
    return true
  }
  return false
}
