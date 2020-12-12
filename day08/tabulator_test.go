package main

import "testing"

func TestTabulation(t *testing.T) {
	instructions := ExtractInstructionsFromInput(getTestInput())
	tabulator := NewTabulator(instructions)

	got := tabulator.Tabulate()
	want := 5

	if got != want {
		t.Errorf("Expected tabulate to get %v, got %v", want, got)
	}

}
