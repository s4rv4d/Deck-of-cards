package main

import (
	"os"
	"testing"
)

//create a test for the function newDeck
//"t" is the go test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//test to check if length of deck is 16
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//to check if first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card of Ace of Spades, but got %v", d[0])
	}

	//to check if the last card if Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

//create a test to check whether saving and retrieving from a file is successfull
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	//first we need to clean the environment basically deleting the file with the name _decktesting from previous tests
	os.Remove("_decktesting")

	//create new deck to save to file
	d := newDeck()
	d.saveToFile("_decktesting")

	//load the deck from file
	ld := newDeckFromFile("_decktesting")

	//check if length of ld is 16
	if len(ld) != 16 {
		t.Errorf("Expected length of deck to be 16, but go %v", len(ld))
	}

	//again clean the entire file to be in the safe side
	os.Remove("_decktesting")

}
