package main

import (
	"os"
	"testing"
)

func TestNewDeckIsCreated(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {

		t.Errorf("Error: Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {

		t.Errorf("Error: Expected first card Ace of Spade, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {

		t.Errorf("Error: Expected last care Four of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {

		t.Errorf("Error: Expected Deck of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
