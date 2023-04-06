package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // 2 packages we use to take input from user: os and bufio
	fmt.Println("Enter the rating for our pizza::: ")

	// comma ok (input,_) || error ok (_,err) syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating: %T\n", input)
}
