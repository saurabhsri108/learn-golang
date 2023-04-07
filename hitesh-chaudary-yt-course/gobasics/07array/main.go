package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[3] = "banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetable List is: ", vegList)
	fmt.Println("Vegetable List is: ", len(vegList))
}
