package main

import (
	"fmt"
)

func main() {
	poisonedBottle := 425
	fmt.Printf("--- THE CELLAR PROBLEM ---\n")
	fmt.Printf("The secret poisoned bottle is: %d\n\n", poisonedBottle)
	tastersDeadStatus := make([]bool, 10)
	for bottleNum := 1; bottleNum <= 1000; bottleNum++ {
		if bottleNum == poisonedBottle {
			for bitPosition := 0; bitPosition < 10; bitPosition++ {
				if (bottleNum & (1 << bitPosition)) != 0 {
					tastersDeadStatus[bitPosition] = true
				}
			}
		}
	}
	detectedPoisonBottle := 0
	for bitPosition := 0; bitPosition < 10; bitPosition++ {
		if tastersDeadStatus[bitPosition] == true {
			detectedPoisonBottle = detectedPoisonBottle + (1 << bitPosition)
		}
	}
	fmt.Printf("--- THE DETECTION RESULT ---\n")
	fmt.Printf("The King's code found the poison in bottle: %d\n", detectedPoisonBottle)
}
