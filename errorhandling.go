package main

import "fmt"

func Div(val1, val2 float64) (float64, error) {
	if val1 == 0 || val2 == 0 {
		return 0, fmt.Errorf("You are putting zero in element")
	}
	return val1 / val2, nil
}

func main() {
	fmt.Println("Hare to call functions")
	ans, err := Div(4, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)
}
