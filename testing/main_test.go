package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected length to be 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card isn't Ace of Spades, it is: %v", d[0])
	}

	if d[len(d)-1] != "Three of Hearts" {
		t.Errorf("Last card isn't Three of Hearts, it is %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("Expected 9 cards, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
