package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected length 16 and got: %v", len(d))

	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 24 {
		t.Errorf("Expected length 16 and got: %v", len(d))

	}
	os.Remove("_decktesting")
}
