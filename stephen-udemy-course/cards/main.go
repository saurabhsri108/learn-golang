package main

func main() {
	cards := newDeck()
	cards.shuffle()
	dealDeck, _ := deal(cards, 3)
	dealDeck.print()
}
