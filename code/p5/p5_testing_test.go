package main

import "testing"

func TestSquare(t *testing.T) {
	got := Square(5)
	if got != 25 {
		t.Errorf("Square(5) = %d, but expected 25", got)
	}
}
