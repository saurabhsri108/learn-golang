package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 spot")
	case 3:
		fmt.Println("Dice value is 3 spot")
	case 4:
		fmt.Println("Dice value is 4 spot")
	case 5:
		fmt.Println("Dice value is 5 spot")
	case 6:
		fmt.Println("Dice value is 6 spot")
		fallthrough
	default:
		fmt.Println("What was that!")
	}

	// fallthrough is for not breaking the switch cases.

}
