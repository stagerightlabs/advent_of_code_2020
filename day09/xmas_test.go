package main

import "testing"

func TestNewXmas(t *testing.T) {
	xmas := NewXmas(getTestInput(), 5)

	if len(xmas.arr) != 20 {
		t.Errorf("Expected Xmas to have an arr length of 20, got %v", len(xmas.arr))
	}
	if xmas.consider != 5 {
		t.Errorf("Expected Xmas to have a consideration length of 5")
	}
}
