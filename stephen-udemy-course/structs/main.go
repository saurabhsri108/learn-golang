package main

import (
	"encoding/json"
	"fmt"
)

type ContactInfo struct {
	Email   string `json:"email"`
	Zipcode int    `json:"zipcode"`
}

type Person struct {
	Firstname string      `json:"firstname"`
	Lastname  string      `json:"lastname"`
	Contact   ContactInfo `json:"contact"`
}

func main() {
	naruto := Person{
		Firstname: "Naruto",
		Lastname:  "Uzumaki",
		Contact: ContactInfo{
			Email:   "naruto@konoha.com",
			Zipcode: 000000,
		},
	}
	narutoPointer := &naruto // &naruto is pointer to naruto's memory block address. pointer(address) = &value
	narutoPointer.updateName("Kenshin")
	naruto.updateName("Obito") // this also works - edge case in golang - shortcut. Automatic conversion happens here if the receiver is of type pointer *Person in this case
	naruto.print()
}

func (p *Person) updateName(newFirstName string) {
	(*p).Firstname = newFirstName // *p is accessing the value stored inside the memory address p pointer is pointing towards. value = *pointer(address)
}

func (p Person) print() {
	result, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
}

// slice by default behaves like a pointer in golang. Why?

// slice [ptr to underlying array, capacity, length]
// ptr to underlying array -> array[values]
// so, the golang still copies the value but the ptr to underlying array values also gets copied so changing the values changes the original value.
// so, slices are reference type in golang.

// int, float, string, bool, structs -> value types [have to think about pointers for them]
// slices, maps, channels, pointers, functions -> reference types [no need to think about pointers for them]
