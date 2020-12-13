package main

import (
	"strconv"
	"strings"
)

// Xmas represents a set of encrypted data
type Xmas struct {
	arr      []int
	consider int
}

// NewXmas creates an array from our preamble input
// with a given length
func NewXmas(input string, length int) Xmas {
	xmas := Xmas{}
	xmas.consider = length
	var lines []string

	// Convert our string input into an array of strings
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	// Set up our list of numbers
	for i := 0; i < len(lines); i++ {
		number, err := strconv.Atoi(lines[i])
		if err != nil {
			panic("invalid input")
		}
		xmas.arr = append(xmas.arr, number)
	}

	return xmas
}
