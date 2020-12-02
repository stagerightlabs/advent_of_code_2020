package main

import (
	"testing"
)

func TestParseRule(t *testing.T) {

	rule, err := NewRule("1-3 a")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if rule.min != 1 {
		t.Errorf("got %v want %v", rule.min, 1)
	}

	if rule.max != 3 {
		t.Errorf("got %v want %v", rule.max, 1)
	}

	if rule.letter != 'a' {
		t.Errorf("got '%c' want '%c'", rule.letter, 'a')
	}

	rule, err = NewRule("2-9 c")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if rule.min != 2 {
		t.Errorf("got %v want %v", rule.min, 1)
	}

	if rule.max != 9 {
		t.Errorf("got %v want %v", rule.max, 1)
	}

	if rule.letter != 'c' {
		t.Errorf("got '%c' want '%c'", rule.letter, 'a')
	}
}

func TestValidateSledPassword(t *testing.T) {
	var testcases = make(map[string]bool)
	testcases["1-3 a: abcde"] = true
	testcases["1-3 b: cdefg"] = false
	testcases["2-9 c: ccccccccc"] = true

	for input, want := range testcases {
		rule, password := ParseInput(input)

		got := ValidateSledPassword(password, rule)

		if got != want {
			t.Errorf("for password '%v': got %t, want %t", password, got, want)
		}

	}
}

func TestValidateTobogganPassword(t *testing.T) {
	var testcases = make(map[string]bool)
	testcases["1-3 a: abcde"] = true
	testcases["1-3 b: cdefg"] = false
	testcases["2-9 c: ccccccccc"] = false

	for input, want := range testcases {
		rule, password := ParseInput(input)

		got := ValidateTobogganPassword(password, rule)

		if got != want {
			t.Errorf("for password '%v': got %t, want %t", password, got, want)
		}

	}
}
