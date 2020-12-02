package main

import (
	"testing"
)

func TestLocateTwoEntries(t *testing.T) {
	var list = []int32{1721, 979, 366, 299, 675, 1456}
	var target int32 = 2020

	entries, err := LocateTwoEntries(list, target)
	if err != nil {
		t.Fatalf("Could not parse test case")
	}

	var got int32 = entries[0] + entries[1]
	var want int32 = target

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestLocateThreeEntries(t *testing.T) {
	var list = []int32{1721, 979, 366, 299, 675, 1456}
	var target int32 = 2020

	entries, err := LocateThreeEntries(list, 2020)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var got int32 = entries[0] + entries[1] + entries[2]
	var want int32 = target

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
