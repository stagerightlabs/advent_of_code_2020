package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	groups := ExtractGroupsFromInput(string(input))

	totalAnyoneCounts := 0
	for _, group := range groups {
		totalAnyoneCounts = totalAnyoneCounts + group.totalAnsweredByAnyone
	}

	fmt.Printf("Answer 1: The count total is %v\n", totalAnyoneCounts)

	totalEveryoneCounts := 0
	for _, group := range groups {
		totalEveryoneCounts = totalEveryoneCounts + group.totalAnsweredByEveryone
	}

	fmt.Printf("Answer 2: The count total is %v\n", totalEveryoneCounts)
}

// ExtractGroupsFromInput splits an input string into a slice of groups
func ExtractGroupsFromInput(input string) []Group {
	var groups []Group

	for _, inp := range strings.Split(input, "\n\n") {
		groups = append(groups, NewGroup(inp))
	}

	return groups
}

// Group represents an answer set from a group of travelers
type Group struct {
	answers                 map[rune]int
	totalAnsweredByAnyone   int
	totalAnsweredByEveryone int
	members                 int
}

// NewGroup creates a Group from an input string
func NewGroup(input string) Group {
	g := Group{}
	g.answers = make(map[rune]int)
	g.members = 0

	// Count the number of unique people in the group
	for _, ans := range strings.Split(input, "\n") {
		if len(ans) > 0 {
			g.members++
		}
	}

	// Count the total number of answers by each letter
	for _, letter := range input {

		if letter < 'a' || letter > 'z' {
			continue
		}

		var count int
		count, exits := g.answers[letter]
		if !exits {
			g.answers[letter] = 0
			count = 0
		}

		g.answers[letter] = count + 1

	}

	// Count all the answers provided by any member
	g.totalAnsweredByAnyone = len(g.answers)

	for _, count := range g.answers {
		if count == g.members {
			g.totalAnsweredByEveryone++
		}
	}

	return g
}
