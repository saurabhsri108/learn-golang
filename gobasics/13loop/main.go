package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Printf("Index: %v <-> Day: %v\n", i, day)
	}

	var rougueValue int

	rougueValue = 1
	for rougueValue < 10 {
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	rougueValue = 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		if rougueValue == 2 {
			goto goToState
		}
		if rougueValue == 8 {
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

goToState:
	fmt.Println("Jumped here")
	// goto goToState
}
