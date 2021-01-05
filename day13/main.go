package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	// Part One
	time, busses := ParseInputForAnswerOne(string(input))
	bus, departure := FindDepartureOption(time, busses)
	minutes := departure - time
	total := bus * minutes
	fmt.Printf("Answer 1: Bus %v departs at %v, %v minute(s) after the target time. Total: %v\n", bus, departure, minutes, total)

	// Part Two
	busMap := ParseInputForAnswerTwo(string(input))
	timestamp := FindDepartureSequenceTimestamp(busMap)
	fmt.Printf("Answer 2: The departure timestamp is %v\n", timestamp)
}

// ParseInputForAnswerOne extracts our departure time and the bus IDs from the input string
func ParseInputForAnswerOne(input string) (int, []int) {

	input = strings.TrimSpace(input)

	breakout := strings.Split(input, "\n")
	departure, err := strconv.Atoi(breakout[0])
	if err != nil {
		panic("invalid input")
	}
	var busses []int

	for _, id := range strings.Split(breakout[1], ",") {
		number, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		busses = append(busses, number)
	}

	return departure, busses
}

// FindDepartureOption looks at an array of bus IDs to find the
// first bus that leaves after our departure time
func FindDepartureOption(time int, busses []int) (int, int) {
	breakdown := make(map[int]int)

	// Find departure options for each bus
	for _, bus := range busses {
		counter := 1

		for {
			if bus*counter >= time {
				break
			}
			counter++
		}

		breakdown[bus] = bus * counter
	}

	// Find the option that departs as close to our departure time as possible.
	departure := time * 2
	bus := 0
	for b, d := range breakdown {
		if d < departure {
			departure = d
			bus = b
		}
	}

	return bus, departure
}

// ParseInputForAnswerTwo derives a relative departure time map from the input string
func ParseInputForAnswerTwo(input string) map[int]int {
	input = strings.TrimSpace(input)
	busses := make(map[int]int)

	breakout := strings.Split(input, "\n")

	for index, id := range strings.Split(breakout[1], ",") {
		number, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		busses[index] = number
	}

	return busses
}

// FindDepartureSequenceTimestamp determines the first time where
// each bus will depart within the offset defined in the input map.
func FindDepartureSequenceTimestamp(busses map[int]int) uint64 {
	timestamp := uint64(0)
	step := uint64(busses[0])

	for {
		match := true
		var matches []uint64
		for offset, bus := range busses {
			if (timestamp+uint64(offset))%uint64(bus) != 0 {
				match = false
			} else {
				matches = append(matches, uint64(bus))
			}
		}

		step = product(matches)

		if match {
			break
		}

		timestamp += step
	}

	return timestamp
}

// Find the product of a slice of ints
func product(numbers []uint64) uint64 {
	product := numbers[0]
	for i := 1; i < len(numbers); i++ {
		product *= numbers[i]
	}
	return product
}
