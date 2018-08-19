package main

import (
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
}
