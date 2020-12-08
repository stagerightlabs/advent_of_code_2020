package main

import (
	"math"
)

// Seat represents a single seat on our plane
type Seat struct {
	ID     int
	Code   string
	Row    int
	Column int
}

// NewSeat creates a seat from an input code that
// represents a binary space partition
func NewSeat(input string) Seat {
	s := Seat{}

	s.Code = input
	s.Row = calculateRow(input[0:7])
	s.Column = calculateColumn(input[7:])
	s.ID = (s.Row * 8) + s.Column

	return s
}

// Interpret a string as a binary number and return the base 10 value
// "F" represents a 0 and "B" represents a 1
func calculateRow(input string) int {

	row := 0
	position := 0

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 'B' {
			row = row + int(math.Pow(2, float64(position)))
		}
		position++
	}

	return row
}

// Interpret a string as a binary number and return the base 10 value
// "L" represents a 0 and "R" represents a 1
func calculateColumn(input string) int {
	column := 0
	position := 0

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 'R' {
			column = column + int(math.Pow(2, float64(position)))
		}
		position++
	}

	return column
}
