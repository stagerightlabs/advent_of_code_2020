package main

import "fmt"

func main() {
	// Answer 1
	got := FindSpokenNumber([]int{11, 18, 0, 20, 1, 7, 16}, 2020)
	fmt.Printf("Answer 1: The 2020th number spoken will be %v\n", got)

	// Answer 2
	got = FindSpokenNumber([]int{11, 18, 0, 20, 1, 7, 16}, 30000000)
	fmt.Printf("Answer 2: The 30000000th number spoken will be %v\n", got)
}

// Tally keeps track of how often a number is spoken
// and which turns it is spoken in.
type Tally struct {
	turns map[int][]int
}

// Get a turn count slice from the tally
func (t *Tally) get(number int) []int {
	counts, exists := t.turns[number]
	if !exists {
		return []int{}
	}
	return counts
}

// Record a new turn for a given number
func (t *Tally) add(number int, turn int) {
	counts, exists := t.turns[number]
	if !exists {
		counts = []int{}
	}
	counts = append(counts, turn)
	t.turns[number] = counts
}

// NewTally returns a new Tally
func NewTally() Tally {
	tally := Tally{}
	tally.turns = make(map[int][]int)

	return tally
}

// FindSpokenNumber runs through the elf's memory game
// for a specificed number of turns and returns the
// final spoken number.
func FindSpokenNumber(input []int, turns int) int {
	// Ensure that the turn count is larger than
	// the length of the input set
	if turns < len(input) {
		return input[turns-1]
	}

	// Initialize our tally from our input set
	tally := NewTally()
	for turn, number := range input {
		tally.add(number, turn+1)
	}

	// Run through the turns until we reach our target
	var spoken int
	for i := len(input); i < turns; i++ {
		previous := input[i-1]
		counts := tally.get(previous)

		if len(counts) == 1 {
			spoken = 0
		} else {
			length := len(counts)
			spoken = counts[length-1] - counts[length-2]
		}

		input = append(input, spoken)
		tally.add(spoken, i+1)
	}

	return spoken
}
