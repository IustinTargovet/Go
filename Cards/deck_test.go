package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expeted deck length of 52, but got %v", len(d))
	}

	if d[0] != "A of Spades" {
		t.Errorf("Expected first value A of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected first value K of Clubs but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("exppected 52 got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
