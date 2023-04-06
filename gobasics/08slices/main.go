package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("FruitList: ", fruitList)
	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("New FruitList: ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("New FruitList 2: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("New FruitList 3: ", fruitList)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 456
	highscore[3] = 567
	// highscore[4] = 666 // panic: runtime error: index out of range [4] with length 4 (runtime error)
	highscore = append(highscore, 555, 666, 321) // here the whole memory allocation happens again which takes a lot of performance hit depending on computation
	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))
	sort.Ints(highscore)
	fmt.Println("Sorted highscrores: ", highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	// remove elements from slices
	courses := []string{"html", "css", "js"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
