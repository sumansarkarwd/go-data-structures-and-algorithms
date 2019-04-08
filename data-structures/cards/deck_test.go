package main

import (
	"os"
	"testing"
)

func TestNewDect(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))
	}

	if d[0] != "Spades of One" {
		t.Errorf("Expected Spades of One at position 0, but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Three" {
		t.Errorf("Expected Spades of One at position 0, but got %v", d[len(d)-1])
	}
}

func TestSaveNewDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, but got %v", len(loadDeck))
	}

	os.Remove("_decktesting")
}
