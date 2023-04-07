package main

import "fmt"

// functions inside structs are called methods

func main() {
	fmt.Println("Methods in golang")
	vasu := User{"Vasu", "vasu@gmail.com", false, 2}
	vasu.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // u is a copy of an object. Any change here doesn't change any original values
	fmt.Println("Is user active:", u.Status)
}
