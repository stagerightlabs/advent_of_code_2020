package main

import "testing"

func TestFindSpokenNumber(t *testing.T) {
	input := []int{0, 3, 6}

	got := FindSpokenNumber(input, 4)
	want := 0
	if got != want {
		t.Errorf("Expected turn 4 to yield %v, got %v", want, got)
	}

	got = FindSpokenNumber(input, 5)
	want = 3
	if got != want {
		t.Errorf("Expected turn 5 to yield %v, got %v", want, got)
	}

	got = FindSpokenNumber(input, 6)
	want = 3
	if got != want {
		t.Errorf("Expected turn 6 to yield %v, got %v", want, got)
	}

	got = FindSpokenNumber(input, 7)
	want = 1
	if got != want {
		t.Errorf("Expected turn 7 to yield %v, got %v", want, got)
	}

	got = FindSpokenNumber(input, 8)
	want = 0
	if got != want {
		t.Errorf("Expected turn 8 to yield %v, got %v", want, got)
	}

	got = FindSpokenNumber(input, 9)
	want = 4
	if got != want {
		t.Errorf("Expected turn 9 to yield %v, got %v", want, got)
	}

	got = FindSpokenNumber(input, 10)
	want = 0
	if got != want {
		t.Errorf("Expected turn 10 to yield %v, got %v", want, got)
	}

	turn := 2020

	got = FindSpokenNumber([]int{1, 3, 2}, turn)
	want = 1
	if got != want {
		t.Errorf("Expected turn %v for example B to yield %v, got %v", turn, want, got)
	}

	got = FindSpokenNumber([]int{2, 1, 3}, turn)
	want = 10
	if got != want {
		t.Errorf("Expected turn %v for example C to yield %v, got %v", turn, want, got)
	}

	got = FindSpokenNumber([]int{1, 2, 3}, turn)
	want = 27
	if got != want {
		t.Errorf("Expected turn %v for example D to yield %v, got %v", turn, want, got)
	}

	got = FindSpokenNumber([]int{2, 3, 1}, turn)
	want = 78
	if got != want {
		t.Errorf("Expected turn %v for example E to yield %v, got %v", turn, want, got)
	}

	got = FindSpokenNumber([]int{3, 2, 1}, turn)
	want = 438
	if got != want {
		t.Errorf("Expected turn %v for example F to yield %v, got %v", turn, want, got)
	}

	got = FindSpokenNumber([]int{3, 1, 2}, turn)
	want = 1836
	if got != want {
		t.Errorf("Expected turn %v for example G to yield %v, got %v", turn, want, got)
	}

	turn = 30000000
	got = FindSpokenNumber([]int{0, 3, 6}, turn)
	want = 175594
	if got != want {
		t.Errorf("Expected turn %v for example H to yield %v, got %v", turn, want, got)
	}

}
