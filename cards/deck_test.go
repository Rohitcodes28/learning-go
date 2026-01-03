package main

import "testing"

// In testing our aim is to use the functions defined in the file make temp vars and check all theur expected functionality
func TestNewDeck(t *testing.T) {
	d := newdeck()

	if len(d) != 16 {
		// Know the difference between error and errorf
		// Errorf needs the %v or %d ertc variable position defination

		t.Errorf(" Expected lenghth of 16 got %v", len(d))

	}

}
