package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Maths!")
	// var mynumber int = 10
	// var mynumber2 float64 = 10.5

	// fmt.Println(mynumber, mynumber2, mynumber2+float64(mynumber))

	// generate a random number
	// fmt.Println("Random number: ", rand.Intn(5))

	// random from crypto
	myRandomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	fmt.Println("Random number: ", myRandomNumber)
}
