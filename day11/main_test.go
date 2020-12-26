package main

import (
	"testing"
)

func TestNewWaitingRoom(t *testing.T) {
	wr := NewWaitingRoom(getTestInput())

	if len(wr.seats[0]) != 10 {
		t.Errorf("Expected a row to have 10 seats, got %v", len(wr.seats[0]))
	}

	if len(wr.seats) != 10 {
		t.Errorf("Expected to have 10 rows, got %v", len(wr.seats))
	}
}

func TestSeatingSimulation(t *testing.T) {
	wr := NewWaitingRoom(getTestInput())

	complete, _ := SimulateSeating(wr)
	occupiedSeatCount := complete.OccupiedSeatCount()

	if occupiedSeatCount != 37 {
		t.Errorf("Expected %v seats to be occupied, got %v", 37, occupiedSeatCount)
	}
}

func TestAlternateeatingSimulation(t *testing.T) {
	wr := NewWaitingRoom(getTestInput())

	complete, _ := SimulateAlternateSeating(wr)
	occupiedSeatCount := complete.OccupiedSeatCount()

	if occupiedSeatCount != 26 {
		t.Errorf("Expected %v seats to be occupied, got %v", 26, occupiedSeatCount)
	}
}

func TestAlternateTicks(t *testing.T) {
	wrA := NewWaitingRoom(getAlternateFirstTickInput())
	wrB := NewWaitingRoom(getAlternateSecondTickInput())

	attempt := DoAlternateTick(wrA)

	if !IsEqual(wrB, attempt) {
		t.Error("Expected a different result from an alternate tick")
	}
}

func TestFindDirectionalSeats(t *testing.T) {
	wr := NewWaitingRoom(`
.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`,
	)

	count := wr.AlternateAdjacentSeatCount(3, 4)

	if count != 8 {
		t.Errorf("Expected to find %v seats, found %v", 8, count)
	}

}

func getTestInput() string {
	return `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`
}

func getAlternateFirstTickInput() string {
	return `
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`
}

func getAlternateSecondTickInput() string {
	return `
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
`
}
