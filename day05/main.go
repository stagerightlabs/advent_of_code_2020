package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Read the input file
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(f)

	// Read the input line by line and create a slice of Seats
	var seats []Seat
	for scanner.Scan() {
		line := scanner.Text()
		seat := NewSeat(strings.TrimSpace(line))
		seats = append(seats, seat)
	}

	// Sort the seats by ID
	sort.Slice(seats, func(i, j int) bool { return seats[i].ID < seats[j].ID })

	// Find the largest seat ID in our collection of seats.
	max := 0
	for _, seat := range seats {
		if seat.ID > max {
			max = seat.ID
		}
	}

	fmt.Printf("Answer 1: The largets seat ID is %v\n", max)

	// Find our seat by locating a gap between seat IDs somewhere in the middle of the set
	prev := seats[0].ID
	var ourSeatId int
	for _, seat := range seats {
		if seat.ID == prev || seat.ID == prev+1 {
			prev = seat.ID
		} else {
			ourSeatId = seat.ID - 1
			break
		}
	}

	fmt.Printf("Answer 2: Our seat ID is %v\n", ourSeatId)
}
