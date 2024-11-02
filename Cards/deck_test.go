package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}

	if d[0] != "1 of Hearts" {
		t.Errorf("Expected 1 of Hearts but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("./_newDeckTest")

	d := newDeck()

	d.saveToFile("./_newDeckTest")

	loadedDeck := newDeckFromFile("./_newDeckTest")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(loadedDeck))
	}

	os.Remove("./newDeckTest")

}
