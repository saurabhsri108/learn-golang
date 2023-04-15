package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"green": "#0000ff",
	}

	// var fruits map[string]string

	// vegetables := make(map[string]string)
	// vegetables["potato"] = "Potato - Aloo"

	// fmt.Println(colors)
	// fmt.Println(fruits)
	// fmt.Println(vegetables)
	printMap(colors)
	delete(colors, "red")
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	// c["red"] = "No Red" : changing values changes original values
	for color, hex := range c {
		fmt.Printf("%s: %s\n", color, hex)
	}
}

// := is used for declaring and initializing the value in one step
