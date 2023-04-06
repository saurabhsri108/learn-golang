package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct 1") // go doesn't have classes

	// No inheritance in golang; No super or parent

	vasu := User{"Vasu", "whocares@gmail.com", false, 28}
	fmt.Println(vasu)
	fmt.Printf("vasu details are: %+v\n", vasu)
	fmt.Printf("Name is %v, age is %v\n", vasu.Name, vasu.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int8
} // capital letters means public in golang.
