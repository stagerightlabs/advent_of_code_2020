package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	wr := NewWaitingRoom(string(input))

	// Part One
	complete, _ := SimulateSeating(wr)
	fmt.Printf("Answer 1: %v occupied seats\n", complete.OccupiedSeatCount())

	// Part Two
	complete, _ = SimulateAlternateSeating(wr)
	fmt.Printf("Answer 2: %v occupied seats, using alternate strategy\n", complete.OccupiedSeatCount())
}

// WaitingRoom represents a collection of seats in a waiting room
type WaitingRoom struct {
	seats [][]rune
}

// NewWaitingRoom creates a WaitingRoom from a string input
func NewWaitingRoom(input string) WaitingRoom {
	var wr = WaitingRoom{}
	for _, row := range strings.Split(input, "\n") {
		if len(row) == 0 {
			continue
		}

		var seats []rune
		for _, seat := range row {
			seats = append(seats, seat)
		}
		wr.seats = append(wr.seats, seats)
	}

	return wr
}

// Print a WaitingRoom to the console
func (wr WaitingRoom) Print() {
	for _, row := range wr.seats {
		for _, seat := range row {
			fmt.Printf("%c", seat)
		}
		fmt.Printf("\n")
	}
}

// OccupiedSeatCount returns the number of occupied seats in the waiting room.
func (wr WaitingRoom) OccupiedSeatCount() int {
	count := 0

	for _, row := range wr.seats {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}

	return count
}

// Clone creates an exact duplicate of a WaitingRoom
func (wr WaitingRoom) Clone() WaitingRoom {
	var cp = WaitingRoom{}

	for _, row := range wr.seats {
		var newRow []rune

		for _, seat := range row {
			newRow = append(newRow, seat)
		}

		cp.seats = append(cp.seats, newRow)
	}

	return cp
}

// AdjacentSeatCount returns the number of occupied seats
// in the eight seat positions next to a given x,y
// coordinate in the WaitingRoom:
// Left, Right, Up, Down, and Diagonal.
func (wr WaitingRoom) AdjacentSeatCount(x, y int) int {
	count := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// Ignore the target seat
			if i == x && j == y {
				continue
			}

			// Check WaitingRoom bounds
			if i < 0 || i >= len(wr.seats) {
				continue
			}
			if j < 0 || j >= len(wr.seats[i]) {
				continue
			}

			if wr.seats[i][j] == '#' {
				count++
			}
		}
	}

	return count
}

// AlternateAdjacentSeatCount returns the number of occupied
// seats, considering the first chair that is visible in
// each of the eight directions around the target seat.
func (wr WaitingRoom) AlternateAdjacentSeatCount(x, y int) int {
	count := 0

	for _, dir := range []string{"N", "NW", "W", "SW", "S", "SE", "E", "NE"} {
		seat, err := wr.FindDirectionalSeat(dir, x, y)
		if err != nil {
			continue
		}
		if seat == '#' {
			count++
		}
	}

	return count
}

var translations = map[string][]int{
	"N":  {0, -1},
	"NW": {1, -1},
	"W":  {1, 0},
	"SW": {1, 1},
	"S":  {0, 1},
	"SE": {-1, 1},
	"E":  {-1, 0},
	"NE": {-1, -1},
}

func translate(dir string, x, y int) (int, int) {
	offset, exists := translations[dir]
	if !exists {
		panic("Invalid direction lookup")
	}

	return x + offset[0], y + offset[1]
}

// FindDirectionalSeat returns the first seat found in a given direction
// from a given set of coordinates.
func (wr WaitingRoom) FindDirectionalSeat(direction string, x, y int) (rune, error) {

	var seat rune

	for {
		xPrime, yPrime := translate(direction, x, y)

		// Check X bounds
		if xPrime < 0 || xPrime >= len(wr.seats[0]) {
			return seat, errors.New("out of bounds")
		}

		// Check Y bounds
		if yPrime < 0 || yPrime >= len(wr.seats) {
			return seat, errors.New("out of bounds")
		}

		// If this new spot contains a seat we
		// will return it. If not, carry on.
		if wr.seats[yPrime][xPrime] == 'L' || wr.seats[yPrime][xPrime] == '#' {
			return wr.seats[yPrime][xPrime], nil
		}

		x = xPrime
		y = yPrime
	}
}

// IsEqual tells us if two WaitingRooms have an identical configuration
func IsEqual(wrA, wrB WaitingRoom) bool {
	for i := 0; i < len(wrA.seats); i++ {
		for j := 0; j < len(wrA.seats[i]); j++ {
			if wrA.seats[i][j] != wrB.seats[i][j] {
				return false
			}
		}
	}

	return true
}

// DoTick simulates a single round of seating in the wating room.
func DoTick(wrA WaitingRoom) WaitingRoom {
	var wrB = wrA.Clone()

	for i := 0; i < len(wrA.seats); i++ {
		for j := 0; j < len(wrA.seats[i]); j++ {

			// Start by copying the current seat value to the new waiting room
			wrB.seats[i][j] = wrA.seats[i][j]

			// Rule:
			// If a seat is empty (L) and there are no occupied seats adjacent to it,
			// the seat becomes occupied.
			if wrA.seats[i][j] == 'L' && wrA.AdjacentSeatCount(i, j) == 0 {
				wrB.seats[i][j] = '#'
			}

			// Rule:
			// If a seat is occupied (#) and four or more seats adjacent to it are
			// also occupied, the seat becomes empty.
			if wrA.seats[i][j] == '#' && wrA.AdjacentSeatCount(i, j) >= 4 {
				wrB.seats[i][j] = 'L'
			}
		}
	}

	return wrB
}

// DoAlternateTick simulates a single round of seating in the wating room,
// using alternate rules.
func DoAlternateTick(wrA WaitingRoom) WaitingRoom {
	var wrB = wrA.Clone()

	for i := 0; i < len(wrA.seats); i++ {
		for j := 0; j < len(wrA.seats[i]); j++ {

			// Start by copying the current seat value to the new waiting room
			// wrB.seats[i][j] = wrA.seats[i][j]

			// Rule:
			// If a seat is empty (L) and there are no occupied seats adjacent to it,
			// the seat becomes occupied.
			if wrA.seats[i][j] == 'L' && wrA.AlternateAdjacentSeatCount(j, i) == 0 {
				wrB.seats[i][j] = '#'
			}

			// Rule:
			// If a seat is occupied (#) and five or more seats adjacent to it are
			// also occupied, the seat becomes empty.
			if wrA.seats[i][j] == '#' && wrA.AlternateAdjacentSeatCount(j, i) >= 5 {
				wrB.seats[i][j] = 'L'
			}
		}
	}

	return wrB
}

// SimulateSeating runs a waiting room through a full seating process
func SimulateSeating(wr WaitingRoom) (WaitingRoom, int) {
	rounds := 0
	var previous WaitingRoom = wr

	for {

		next := DoTick(previous)

		if IsEqual(next, previous) {
			return previous, rounds
		}

		previous = next
		rounds++
	}
}

// SimulateAlternateSeating runs a waiting room through a full seating process
func SimulateAlternateSeating(wr WaitingRoom) (WaitingRoom, int) {
	rounds := 0
	var previous WaitingRoom = wr

	for {

		next := DoAlternateTick(previous)

		if IsEqual(next, previous) {
			return previous, rounds
		}

		previous = next
		rounds++
	}

}
