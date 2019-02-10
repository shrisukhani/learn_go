package main

import (
	"os"
	"testing"
)

func TestNewDeckReturnsLength52(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(d))
	}
}

func TestNewDeckFirstCardIsAceOfHearts(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be %v, but was %v", "Ace of Hearts", d[0])
	}
}

func TestNewDeckLastCardIsKingOfClubs(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but was %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	for i := range loadedDeck {
		if d[i] != loadedDeck[i] {
			t.Errorf("The written deck does not match the loaded deck")
		}
	}

	os.Remove("decktesting")
}
