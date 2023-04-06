package main

import "fmt"

func main() {
	defer fmt.Println("Finally last")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/**
Hello
4
3
2
1
0
World
Finally last
*/
