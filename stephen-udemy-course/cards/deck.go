package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this is a receiver function
// we are attaching this function to the deck type
// we are saying that any variable of type deck gets access to the print method
func (d deck) print() {
	// we are using := because we are creating a new variable called i and card and we are assigning them to the values of the range of cards everytime we loop through the cards
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
