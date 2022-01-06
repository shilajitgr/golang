package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	card := newDeck()
	if len(card) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(card))
	}

	if card[0] != "Ace of Hearts" {
		t.Errorf("Expected 1st card to be Ace of Hearts but go %v", card[0])
	}

	if card[len(card)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Ace of Hearts but go %v", card[len(card)-1])
	}
}

func TestSaveToFileandDeckfromFile(t *testing.T) {
	os.Remove("_decktesting")

	card := newDeck()
	card.saveToFile("_decktesting")

	loaded_deck := deckFromFile("_decktesting")

	if len(loaded_deck) != 16 {
		t.Errorf("Expected length of loaded is supposed to be 16 but got %v", len(loaded_deck))
	}

}
