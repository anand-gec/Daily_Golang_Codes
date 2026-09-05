package main

import "fmt"

type fruitQuantity struct {
	Mango  int
	Banana int
	Apple  int
}

func main() {
	var Quantity = fruitQuantity{
		Apple:  4,
		Banana: 5,
		Mango:  9,
	}
	fmt.Println("Quantity of that is", Quantity)
	store := fruitQuantity{Banana: 4, Apple: 7, Mango: 5}
	fmt.Println("plane", store)
	fmt.Printf("Quantity of :\nBanana is %d\nMango is %d\nApple is %d", store.Banana, store.Mango, store.Apple)
}
