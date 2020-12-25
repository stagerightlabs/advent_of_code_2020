package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	start := time.Now()
	joltage := ArrangementFromString(string(input))
	distribution := Distribution(joltage)
	product := distribution[1] * distribution[3]
	fmt.Printf("Answer 1: %v * %v = %v\n", distribution[1], distribution[3], product)

	possibilities := uint64(Possibilities(joltage))
	fmt.Printf("Answer 2: There are %d possible arrangements.", possibilities)

	elapsed := time.Since(start)
	fmt.Printf("Calculated in %s\n", elapsed)
}

// Arrangement represents one ordering of adapters
type Arrangement struct {
	adapters []int
	final    int
}

// IsComplete tells us whether or not this arrangement has been completed.
func (a *Arrangement) IsComplete() bool {

	if len(a.adapters) == 0 {
		return false
	}

	return a.adapters[len(a.adapters)-1] == a.final
}

// Add a number to an arrangement's adapter list
func (a *Arrangement) Add(number int) {
	a.adapters = append(a.adapters, number)
}

// ArrangementFromString creates a sorted array of adapter ratings
func ArrangementFromString(input string) Arrangement {
	arrangement := Arrangement{}
	arrangement.adapters = []int{0}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		number, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err.Error())
		}

		arrangement.adapters = append(arrangement.adapters, number)
	}

	// Add our final joltage
	max := max(arrangement.adapters)
	arrangement.adapters = append(arrangement.adapters, max+3)
	arrangement.final = max + 3

	sort.Ints(arrangement.adapters)
	return arrangement
}

func max(joltage []int) int {
	max := joltage[0]
	for _, number := range joltage {
		if number > max {
			max = number
		}
	}
	return max
}

// Distribution creates a map that represents all of the
// differences between values in an array of ints
func Distribution(arrangement Arrangement) map[int]int {
	distribution := make(map[int]int)

	for i := 1; i < len(arrangement.adapters); i++ {
		diff := arrangement.adapters[i] - arrangement.adapters[i-1]

		count, exists := distribution[diff]
		if !exists {
			count = 0
			distribution[diff] = count
		}
		distribution[diff] = count + 1
	}

	return distribution
}

// DiffList creates a list of the subgroups within an adapter arrangement
func DiffList(arrangement Arrangement) [][]int {
	var diffs = make([]int, 0)
	for i := 1; i < len(arrangement.adapters); i++ {
		diff := arrangement.adapters[i] - arrangement.adapters[i-1]
		diffs = append(diffs, diff)
	}

	var subgroups = make([][]int, 0)
	index := 0
	for {

		group := []int{}

		for {
			group = append(group, diffs[index])
			index++
			if index >= len(diffs) {
				break
			}
			if diffs[index] != diffs[index-1] {
				break
			}
		}

		subgroups = append(subgroups, group)

		if index >= len(diffs) {
			break
		}
	}

	return subgroups
}

// Possibilities calculates the total number of possible arrangements.
func Possibilities(arrangement Arrangement) float64 {

	diffs := DiffList(arrangement)

	// Determine how many distinct groups of "1"s we have
	var breakdown = make(map[int]int, 0)
	for _, group := range diffs {
		if group[0] == 3 {
			continue
		}

		length := len(group)
		count, exists := breakdown[length]
		if !exists {
			count = 0
			breakdown[length] = count
		}
		breakdown[length] = count + 1
	}

	// https://schnouki.net/post/2020/advent-of-code-2020-day-10/
	multipliers := map[int]int{
		2: 2,
		3: 4,
		4: 7,
		5: 13,
	}

	var total float64 = 1
	for length, count := range breakdown {
		if length == 1 {
			continue
		}
		total *= math.Pow(float64(multipliers[length]), float64(count))
	}

	return total
}
