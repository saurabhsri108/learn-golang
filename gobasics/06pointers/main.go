package main

import "fmt"

func main() {
	fmt.Println("Welcomet to pointers")

	a := 1

	ptr := &a

	ptr2 := &ptr

	fmt.Println(a, ptr, ptr2)
	fmt.Println(a, *ptr, *ptr2)
	fmt.Println(a, *ptr, **ptr2)
}
