package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, recieved %v ", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, recieved %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, recieved %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, recieved %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
