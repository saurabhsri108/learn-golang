package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Functions in golang")
	greeter()

	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')
	greeter2(name)

	result := adder(1, 2)
	fmt.Println("Result:", result)

	result2 := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro Result2:", result2)

	result3, message := proAdderMessage(1, 2, 3, 4, 5)
	fmt.Println("Pro Result3:", result3)
	fmt.Println("Pro Message:", message)
}

func proAdderMessage(values ...int) (int, string) {
	total := proAdder(values...)
	return total, "My Message"
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func adder(value1 int, value2 int) int {
	return value1 + value2
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func greeter2(name string) {
	fmt.Println("Namaste", name)
}
