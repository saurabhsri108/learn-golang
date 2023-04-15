package main

import (
	"os"
	"testing"
)

const totalCardsInDeck int = 52

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	length := len(cards)
	lastIndex := len(cards) - 1

	if length != totalCardsInDeck {
		t.Errorf("Expected deck length of 52 but got %v", length)
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %v", cards[0])
	}

	if cards[lastIndex] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %v", cards[lastIndex])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != totalCardsInDeck {
		t.Errorf("Expected 52 cards in deck, got %v", len(deck))
	}

	os.Remove("_decktesting")
}
