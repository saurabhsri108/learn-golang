package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 11
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// if err != nil {
	// 	return nil, err
	// }
}
