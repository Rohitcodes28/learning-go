package main

import (
	"os"
	"testing"
)

// In testing our aim is to use the functions defined in the file make temp vars and check all theur expected functionality
func TestNewDeck(t *testing.T) {
	d := newdeck()

	if len(d) != 16 {
		// Know the difference between error and errorf
		// Errorf needs the %v or %d ertc variable position defination

		t.Errorf(" Expected lenghth of 16 got %v", len(d))

	}
	if d[0] != "Ace of spades" {
		t.Errorf("the first card should have been Ace of spades but its %v", d[0])
	}
}

func TestSavetodeckandRetriveDeck(t *testing.T) {
	// function name should be same acros all the tests so if we change the main function we can find the test for the same to modify.

	// There is a cleanup function since we will be writing to files and we need to clean up after

	os.Remove("_decktesting")

	deck := newdeck()
	deck.saveToFile("_decktesting")

	loadeddeck := newDeckFromFile("_decktesting")

	if len(loadeddeck) != 16 {
		t.Errorf("the file is not exactly loaded")
	}

	os.Remove("_decktesting")

}
