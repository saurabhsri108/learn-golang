package main

import "fmt"

func main() {
	var numbers []int

	for number := 0; number < 11; number++ {
		numbers = append(numbers, number)
	}

	fmt.Println("Numbers are:", numbers)

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("Number %v is even\n", number)
		} else {
			fmt.Printf("Number %v is odd\n", number)
		}
	}
}
